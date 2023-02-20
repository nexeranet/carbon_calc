package carbon_calc

import (
	"fmt"
	"strconv"
	"testing"
)

// TODO: ask @TM about float64 precision
// used for tests 0.0001
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
		// default value
		{fraction, 0.05, 5, form, density, biomass, ratio, 0.0167},
		{0, 0.05, 5, 0, DensityOverBarkOfTrees[ForestTypeTropicalSubtropical][TreeSpeciesMoist], biomass, ratio, 0.0167},
	}
	for i, tt := range tests {
		result := CarbonPerTree(tt.fraction, tt.radius, tt.height, tt.form, tt.density, tt.biomass, tt.ratio)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
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
	}
	for i, tt := range tests {
		result := CarbonStoredInPlot(tt.sum, tt.area)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
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
	}
	for i, tt := range tests {
		result := CarbonStoredInMonitoringZone(tt.sum, tt.num, tt.area)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
		}
	}
}

// TODO: test with real arguments
func TestVarianceOfTreeBiomass(t *testing.T) {
	type Test struct {
		plots  []float64
		result float64
	}
	tests := []Test{
		{[]float64{1, 2, 3, 4, 5}, 2.5},
	}
	for i, tt := range tests {
		result := VarianceOfTreeBiomass(tt.plots)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
		}
	}
}

// TODO: test with real arguments
func TestUncertaintyCarbonStored(t *testing.T) {
	type Test struct {
		tDelta, variance, area, sumCorbonStoredPlots float64
		plotAreas                                    []float64
		result                                       float64 // precision = 3
	}
	tDelta := TDistribution(10)
	tests := []Test{
		{tDelta, 12, 12, 12, []float64{12, 12}, 0.523},
	}
	for i, tt := range tests {
		result := UncertaintyCarbonStored(tt.tDelta, tt.variance, tt.area, tt.sumCorbonStoredPlots, tt.plotAreas)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
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
		result := UncertaintyDiscount(tt.uncertainty)
		if result != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
		}
	}
}

// TODO: test with real arguments
func TestConservativeTotalCarbon(t *testing.T) {
	type Test struct {
		totalCarbon, uncertainty float64
		result                   float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 23},
	}
	for i, tt := range tests {
		result := ConservativeTotalCarbon(tt.totalCarbon, tt.uncertainty)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
		}
	}
}

// TODO: test with real arguments
func TestAboveGroundBiomass(t *testing.T) {
	type Test struct {
		cTotalCarbon, ratio, cfTree, area float64
		result                            float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 12, 12, 0.0017},
		{12, 12, 0, 12, 0.0446},
		{12, 12, 0.47, 12, 0.0446},
	}
	for i, tt := range tests {
		result := AboveGroundBiomass(tt.cTotalCarbon, tt.ratio, tt.cfTree, tt.area)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
		}
	}
}
