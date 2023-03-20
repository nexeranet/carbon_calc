package carbon_calc

import (
	"math"

	"github.com/shopspring/decimal"
	"gonum.org/v1/gonum/stat/distuv"
)

// Calculate the carbon stored in each tree
// For each tree present in the current stage and the previous affirmed stage,
// calculate the carbon stored in the tree based on the radius and height
// measurements made on the ground as well as specie / forest type specific
// parameters:
// fraction - carbon fraction of tree biomass
// radius - radius of tree
// height - height of tree
// form - form factor of the tree
// density - density (over-bark) of tree depending on its tree species / forest type
// biomass - biomass expansion factor for conversion of tree stem biomass to
// above-ground tree biomass, for tree l depending on tree species / forest type
// ratio - root-shoot ratio for tree l depending on its specie / forest type
func CarbonPerTree(fraction, radius, height, form, density, biomass, ratio decimal.Decimal) decimal.Decimal {
	if fraction.Equal(decimal.Zero) {
		fraction = decimal.NewFromFloat(0.47)
	}
	if form.Equal(decimal.Zero) {
		form = decimal.NewFromFloat(0.3)
	}
	return decimal.NewFromFloat(44.0 / 12.0).
		Mul(fraction).
		Mul(CircleArea(radius)).
		Mul(height).
		Mul(form).
		Mul(decimal.NewFromFloat(1.2)).
		Mul(density).
		Mul(biomass).
		Mul((decimal.New(1, 0).Add(ratio)))
}

// Calculate the carbon stored in each tree and with params validation
// For more comments see CarbonPerTree function
func ValidateCarbonPerTree(fraction, radius, height, form, density, biomass, ratio decimal.Decimal) (decimal.Decimal, error) {
	if height.Cmp(decimal.NewFromFloat(1.3)) == -1 {
		return decimal.Decimal{}, NotEnoughHeight
	}
	return CarbonPerTree(fraction, radius, height, form, density, biomass, ratio), nil
}

// Carbon/ha stored in sample plot p of monitoring zone
// sum - carbon stored in tree of species in sample plot of monitoring zone
// area - area of sample plot of monitoring zone
func CarbonStoredInPlot(sum, area decimal.Decimal) decimal.Decimal {
	return sum.Div(area)
}

// Calculate the carbon stored in each monitoring zone
// sumOfPlots - carbon/ha stored in sample plot of monitoring zone
// area - area of monitoring zone
// numPlots - number of sample plots in monitoring zone
func CarbonStoredInMonitoringZone(sumOfPlots, numPlots, area decimal.Decimal) decimal.Decimal {
	return sumOfPlots.Div(numPlots).Mul(area)
}

// si^2_i
// Variance of tree biomass per hectare across all sample plots in monitoring zone
// carbonStoredPlots - an array of carbon in each plot in the monitoring zone
func VarianceOfTreeBiomass(carbonStoredPlots []decimal.Decimal) decimal.Decimal {
	n := decimal.NewFromInt(int64(len(carbonStoredPlots)))
	sum := decimal.New(0, 0)
	sumSqrt := decimal.New(0, 0)
	for _, value := range carbonStoredPlots {
		sum = sum.Add(value)
		sumSqrt = sumSqrt.Add(value.Pow(decimal.New(2, 0)))
	}
	sum = sum.Pow(decimal.New(2, 0))
	return n.Mul(sumSqrt).
		Sub(sum).
		Div((n.Mul(n.Sub(decimal.New(1, 0)))))
}

// Two-sided Student’s t-value for a confidence level of 90 percent
// freedom - degrees of freedom equal to n – M, where n is total number
// of sample plots within the tree biomass monitoring zones and M is the
// total number of tree biomass monitoring zones
func TDistribution(freedom float64) decimal.Decimal {
	dist1 := distuv.StudentsT{
		Mu:    0,
		Sigma: 1,
		Nu:    freedom,
		Src:   nil,
	}
	return decimal.NewFromFloat(dist1.Quantile(0.95))
}

// plots - array contains calculated carbon in each plot
// area - area of zone (ha)
type CarbonedZone struct {
	Plots []decimal.Decimal
	Area  decimal.Decimal
}

// Uncertainty in carbon stock in trees
// tDelta - t-distribution
// area - area of all monitoring zones (sum of all areas of monitoring zones)
// zones - array of zones with area and array of cabon in each plot
func UncertaintyCarbonStored(tDelta, tArea decimal.Decimal, zones []CarbonedZone) decimal.Decimal {
	sumAi := decimal.New(0, 0)
	sumAiPow := decimal.New(0, 0)
	for _, zone := range zones {
		nI := decimal.NewFromInt(int64(len(zone.Plots)))
		aiDiv := zone.Area.Div(tArea)
		aiDivPow := aiDiv.Pow(decimal.New(2, 0))
		sumAi = sumAi.Add(aiDiv.Mul(SumDecimal(zone.Plots).Div(nI)))
		sumAiPow = sumAiPow.Add(aiDivPow.Mul(VarianceOfTreeBiomass(zone.Plots).Div(nI)))
	}
	sumSqrt := decimal.NewFromFloat(math.Sqrt(sumAiPow.InexactFloat64()))
	return tDelta.Mul(sumSqrt).Div(sumAi.Abs())
}

// If uncertainty > 10%, then carbon stored in monitoring zones are made
// conservative by applying an uncertainty discount
func UncertaintyDiscount(uncertainty decimal.Decimal) decimal.Decimal {
	uncrt := uncertainty.InexactFloat64()
	if uncrt <= 0.1 {
		return decimal.Zero
	} else if 0.1 < uncrt && uncrt <= 0.15 {
		return decimal.NewFromFloat(0.25)
	} else if 0.15 < uncrt && uncrt <= 0.2 {
		return decimal.NewFromFloat(0.5)
	} else if 0.2 < uncrt && uncrt <= 0.3 {
		return decimal.NewFromFloat(0.75)
	} else {
		return decimal.NewFromFloat(1)
	}
}

// Calculate the total carbon stored in all monitoring zones taking into account
// the uncertainty
// totalCarbon - carbon stored in all monitoring zones
// uncertainty - uncertainty in carbon stock in trees
func ConservativeTotalCarbon(totalCarbon, uncertainty decimal.Decimal) decimal.Decimal {
	return totalCarbon.Mul(decimal.New(1, 0).Sub(uncertainty.Mul(UncertaintyDiscount(uncertainty))))
}

// TODO: change names of arguments
// Carbon stock in trees in monitoring zone i taking into account the
// uncertainty
// conservativeCarbon - Carbon stock in trees in all monitoring zones taking
// into account the uncertainty
// carbonArea - Carbon stock in trees in monitoring zone
// totalAreasCarbon - Carbon stock in trees in all monitoring zones
func AreaConservativeCarbon(conservativeCarbon, carbonArea, totalAreasCarbon decimal.Decimal) decimal.Decimal {
	return conservativeCarbon.Mul(carbonArea.Div(totalAreasCarbon))
}

// Calculate the above ground biomass
// areaConsCarbon -  conservative total carbon in monitoring zone
// ratio - root-shoot ratio for tree depending on its specie / forest type
// cfTree - carbon fraction of tree biomass
// area - area of all monitoring zones
func AboveGroundBiomass(areaConsCarbon, ratio, cfTree, area decimal.Decimal) decimal.Decimal {
	if cfTree.Equal(decimal.Zero) {
		cfTree = decimal.NewFromFloat(0.47)
	}
	return areaConsCarbon.Mul(decimal.New(12, 0)).
		Div(decimal.New(44, 0).
			Mul(cfTree).
			Mul(decimal.New(1, 0).Add(ratio)).
			Mul(area))
}
