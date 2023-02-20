package carbon_calc

import (
	"math"

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
func CarbonPerTree(fraction float64, radius float64, height float64, form float64, density float64, biomass float64, ratio float64) float64 {
	if fraction == 0 {
		fraction = 0.47
	}
	if form == 0 {
		form = 0.25
	}
	return (44 / 12) * fraction * CircleArea(radius) * height * form * 1.2 * density * biomass * (1 + ratio)
}

// Carbon/ha stored in sample plot p of monitoring zone
// sum - carbon stored in tree of species in sample plot of monitoring zone
// area - area of sample plot of monitoring zone
func CarbonStoredInPlot(sum float64, area float64) float64 {
	return sum / area
}

// Calculate the carbon stored in each monitoring zone
// sumOfPlots - carbon/ha stored in sample plot of monitoring zone
// area - area of monitoring zone
// numPlots - number of sample plots in monitoring zone
func CarbonStoredInEachMonitoringZone(sumOfPlots, numPlots, area float64) float64 {
	return (sumOfPlots / numPlots) * area
}

// Variance of tree biomass per hectare across all sample plots in monitoring zone
// carbonStoredPlots - an array of carbon in each plot in the monitoring zone
func VarianceOfTreeBiomass(carbonStoredPlots []float64) float64 {
	n := float64(len(carbonStoredPlots))
	var sum float64 = 0
	var sumSqrt float64 = 0
	for _, value := range carbonStoredPlots {
		sum += value
		sumSqrt += math.Pow(value, 2)
	}
	return (n*sumSqrt - math.Pow(sum, 2)) / (n * (n - 1))
}

// Two-sided Student’s t-value for a confidence level of 90 percent
// freedom - degrees of freedom equal to n – M, where n is total number
// of sample plots within the tree biomass monitoring zones and M is the
// total number of tree biomass monitoring zones
func TDistribution(freedom float64) float64 {
	dist1 := distuv.StudentsT{
		Mu:    0,
		Sigma: 1,
		Nu:    freedom,
		Src:   nil,
	}
	return dist1.Quantile(0.95)
}

// Uncertainty in carbon stock in trees
// tDelta - t-distribution
// variance -  variance of tree biomass
// area - area of all monitoring zones (sum of all areas of monitoring zones)
// sumCorbonStoredPlots - sum of carbon/ha stored in sample plot of monitoring zone
// plotAreas - an array of plot areas in the monitoring zone
func UncertaintyCarbonStored(tDelta, variance, area, sumCorbonStoredPlots float64, plotAreas []float64) float64 {
	var sumA, sumB, b float64
	nPlot := float64(len(plotAreas))
	b = sumCorbonStoredPlots / nPlot
	for _, plot := range plotAreas {
		sumA += math.Pow(plot/area, 2) * variance / nPlot
		sumB += (plot / area) * b
	}
	return (tDelta * math.Sqrt(sumA)) / sumB
}

// If uncertainty > 10%, then carbon stored in monitoring zones are made
// conservative by applying an uncertainty discount
func UncertaintyDiscount(uncertainty float64) float64 {
	if uncertainty <= 0.1 {
		return 0
	} else if 0.1 < uncertainty && uncertainty <= 0.15 {
		return 0.25
	} else if 0.15 < uncertainty && uncertainty <= 0.2 {
		return 0.5
	} else if 0.2 < uncertainty && uncertainty <= 0.3 {
		return 0.75
	} else {
		return 1
	}
}

// Calculate the total carbon stored in all monitoring zones taking into account
// the uncertainty
// totalCarbon - carbon stored in all monitoring zones
// uncertainty - uncertainty in carbon stock in trees
func ConservativeTotalCarbon(totalCarbon, uncertainty float64) float64 {
	return totalCarbon - (1 - uncertainty*UncertaintyDiscount(uncertainty))
}

// Calculate the above ground biomass
// cTotalCarbon -  conservative total carbon
// ratio - root-shoot ratio for tree depending on its specie / forest type
// cfTree - carbon fraction of tree biomass
// area - area of all monitoring zones
func AboveGroundBiomass(cTotalCarbon, ratio, cfTree, area float64) float64 {
	if cfTree == 0 {
		cfTree = 0.47
	}
	return cTotalCarbon * (12 / 44) * (1 / cfTree) * (1/1 + ratio) * (1 / area)
}
