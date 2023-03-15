package carbon_calc

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestCarbonPerTree(t *testing.T) {
	type Test struct {
		fraction, radius, height, form, density, biomass, ratio float64
		result                                                  float64 // precision = 4
	}
	// Default values
	fraction := 0.47
	ratio := 0.3
	form := 0.25
	density := 0.55
	biomass := 1.15
	tests := []Test{
		{0.47, 0.05, 5, form, density, biomass, ratio, 0.0167},
		{fraction, 0.03, 1.53, form, density, biomass, ratio, 0.0018},
		{fraction, 0.025, 1.68, form, density, biomass, ratio, 0.0014},
		{fraction, 0.015, 1.54, form, density, biomass, ratio, 0.0005},
		{fraction, 0.02, 1.45, form, density, biomass, ratio, 0.0008},
		{fraction, 0.025, 1.44, form, density, biomass, ratio, 0.0012},
		{fraction, 0.036, 1.6, form, density, biomass, ratio, 0.0028},
		{fraction, 0.036, 1.6, form, density, biomass, ratio, 0.0028},
		{fraction, 0.036, 1.9, form, density, biomass, ratio, 0.0033},
		{fraction, 0.031, 2.4, form, density, biomass, ratio, 0.0031},
		// default value
		{fraction, 0.05, 5, form, density, biomass, ratio, 0.0167},
		{0, 0.05, 5, 0, DensityOverBarkOfTrees[ForestTypeTropicalSubtropical][TreeSpeciesMoist], biomass, ratio, 0.02},
	}
	for i, tt := range tests {
		result := CarbonPerTree(
			decimal.NewFromFloat(tt.fraction),
			decimal.NewFromFloat(tt.radius),
			decimal.NewFromFloat(tt.height),
			decimal.NewFromFloat(tt.form),
			decimal.NewFromFloat(tt.density),
			decimal.NewFromFloat(tt.biomass),
			decimal.NewFromFloat(tt.ratio))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestCarbonStoredInPlot(t *testing.T) {
	type Test struct {
		sum, area float64
		result    float64 // precision = 4
	}
	tests := []Test{
		{Sum([]float64{0.0167, 0.0167, 0.0167}), 0.0201, 2.4925},
		{Sum([]float64{0.0167, 0.0167}), 0.0201, 1.6617},
		{Sum([]float64{0.0167, 0.0018}), 0.03, 0.6167},
		{Sum([]float64{0.0014, 0.0005, 0.0008, 0.0012}), 0.03, 0.13},
		{Sum([]float64{0.0027, 0.0033, 0.0031, 0.006}), 0.02, 0.755},
		{Sum([]float64{0.0064, 0.0004, 0.0051, 0.0065}), 0.02, 0.92},
		{Sum([]float64{0.0085, 0.0017, 0.0106}), 0.02, 1.04},
	}
	for i, tt := range tests {
		result := CarbonStoredInPlot(
			decimal.NewFromFloat(tt.sum),
			decimal.NewFromFloat(tt.area))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestCarbonStoredInEachMonitoringZone(t *testing.T) {
	type Test struct {
		sum, num, area float64
		result         float64 // precision = 3
	}
	tests := []Test{
		{Sum([]float64{0.59, 0.12}), 2, 1, 0.355},
		{Sum([]float64{2.49, 2.49, 1.66}), 3, 8, 17.707},
		{Sum([]float64{0.75, 0.91, 1.03}), 3, 8, 7.173},
		{Sum([]float64{0.38, 0.95}), 2, 1, 0.665},
	}
	for i, tt := range tests {
		result := CarbonStoredInMonitoringZone(
			decimal.NewFromFloat(tt.sum),
			decimal.NewFromFloat(tt.num),
			decimal.NewFromFloat(tt.area))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestVarianceOfTreeBiomass(t *testing.T) {
	type Test struct {
		plots  []float64
		result float64
	}
	tests := []Test{
		{[]float64{1, 2, 3, 4, 5}, 2.5},
		{[]float64{0.359, 0.736, 0.889}, 0.074},
		{[]float64{0.27, 0.7}, 0.092},
		{[]float64{0.75, 0.91, 1.05}, 0.023},
		{[]float64{0.38, 0.95}, 0.162},
	}
	for i, tt := range tests {
		plots := []decimal.Decimal{}
		for _, plot := range tt.plots {
			plots = append(plots, decimal.NewFromFloat(plot))
		}
		result := VarianceOfTreeBiomass(plots)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestUncertaintyCarbonStored(t *testing.T) {
	type CarbonedZoneTest struct {
		area  float64
		plots []float64
	}
	type Test struct {
		tDelta decimal.Decimal
		zones  []CarbonedZoneTest
		result float64 // precision = 3
	}
	tests := []Test{
		{TDistribution(3), []CarbonedZoneTest{
			{
				area:  8,
				plots: []float64{0.359, 0.736, 0.889},
			},
			{
				area:  1,
				plots: []float64{0.27, 0.7},
			},
		}, 0.521},
		{TDistribution(3), []CarbonedZoneTest{
			{
				area:  8,
				plots: []float64{0.75, 0.91, 1.03},
			},
			{
				area:  1,
				plots: []float64{0.38, 0.95},
			},
		}, 0.213},
	}
	for i, tt := range tests {
		zones := []CarbonedZone{}
		var tArea, nPlots float64
		for _, zone := range tt.zones {
			var plots []decimal.Decimal
			for _, item := range zone.plots {
				plots = append(plots, decimal.NewFromFloat(item))
				nPlots++
			}
			zones = append(zones, CarbonedZone{
				plots: plots,
				area:  decimal.NewFromFloat(zone.area),
			})
			tArea += zone.area
		}
		result := UncertaintyCarbonStored(tt.tDelta,
			decimal.NewFromFloat(tArea),
			decimal.NewFromFloat(nPlots),
			zones)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestUncertaintyDiscount(t *testing.T) {
	type Test struct {
		uncertainty float64
		result      float64
	}
	tests := []Test{
		{0.01, 0},
		{0.09, 0},
		{0.1, 0},
		{0.11, 0.25},
		{0.15, 0.25},
		{0.16, 0.5},
		{0.2, 0.5},
		{0.21, 0.75},
		{0.3, 0.75},
		{0.4, 1},
		{0.5, 1},
		{0.9, 1},
	}
	for i, tt := range tests {
		result := UncertaintyDiscount(decimal.NewFromFloat(tt.uncertainty))
		if result.InexactFloat64() != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestConservativeTotalCarbon(t *testing.T) {
	type Test struct {
		totalCarbon, uncertainty float64
		result                   float64 // precision = 3
	}
	tests := []Test{
		{5.78, 0.519, 2.78},
		{7.86, 0.214, 6.598},
	}
	for i, tt := range tests {
		result := ConservativeTotalCarbon(
			decimal.NewFromFloat(tt.totalCarbon),
			decimal.NewFromFloat(tt.uncertainty))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestAreaConservativeCarbon(t *testing.T) {
	type Test struct {
		conservativeCarbon, carbonArea, totalAreasCarbon float64
		result                                           float64 // precision = 3
	}
	tests := []Test{
		{2.77, 5.29, 5.78, 2.535},
		{2.77, 0.49, 5.78, 0.235},
		{6.6, 7.2, 7.86, 6.046},
		{6.6, 0.66, 7.86, 0.554},
	}
	for i, tt := range tests {
		result := AreaConservativeCarbon(
			decimal.NewFromFloat(tt.conservativeCarbon),
			decimal.NewFromFloat(tt.carbonArea),
			decimal.NewFromFloat(tt.totalAreasCarbon))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestAboveGroundBiomass(t *testing.T) {
	type Test struct {
		cTotalCarbon, ratio, cfTree, area float64
		result                            float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 0, 12, 0.0446},
		{12, 12, 0.47, 12, 0.0446},
		{2.54, 0.3, 0.47, 8, 0.1417},
		{0.234, 0.3, 0.47, 1, 0.1044},
		{6.041, 0.3, 0.47, 8, 0.3371},
		{0.588, 0.3, 0.47, 1, 0.2625},
	}
	for i, tt := range tests {
		result := AboveGroundBiomass(decimal.NewFromFloat(tt.cTotalCarbon),
			decimal.NewFromFloat(tt.ratio),
			decimal.NewFromFloat(tt.cfTree),
			decimal.NewFromFloat(tt.area))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}
