package carbon_calc

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestBaselineInMonitoringZone(t *testing.T) {
	type Test struct {
		manual, area, deltaTime float64
		result                  float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 12, 1728},
	}
	for i, tt := range tests {
		result := BaselineInMonitoringZone(
			decimal.NewFromFloat(tt.manual),
			decimal.NewFromFloat(tt.area),
			decimal.NewFromFloat(tt.deltaTime))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestNetGHGEmissions(t *testing.T) {
	type Test struct {
		fertilizes, cO2eNdirectt, cO2eNindirectt float64
		result                                   float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 12, 288},
	}
	for i, tt := range tests {
		result := NetGHGEmissions(
			decimal.NewFromFloat(tt.fertilizes),
			decimal.NewFromFloat(tt.cO2eNdirectt),
			decimal.NewFromFloat(tt.cO2eNindirectt))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestCO2eNdirectt(t *testing.T) {
	type Test struct {
		massSynthFertz, nContSynthFertz, massOrgFertz, nContOrgFertz, nitrOxdEmissSOC, gWarmingPotentl float64
		result                                                                                         float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 12, 12, 12, 12, 65170.286},
	}
	for i, tt := range tests {
		result := CO2eNdirectt(
			decimal.NewFromFloat(tt.massSynthFertz),
			decimal.NewFromFloat(tt.nContSynthFertz),
			decimal.NewFromFloat(tt.massOrgFertz),
			decimal.NewFromFloat(tt.nContOrgFertz),
			decimal.NewFromFloat(tt.nitrOxdEmissSOC),
			decimal.NewFromFloat(tt.gWarmingPotentl))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestCO2eNdirecttDefault(t *testing.T) {
	type Test struct {
		massSynthFertz, massOrgFertz float64
		result                       float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 10.994},
	}
	for i, tt := range tests {
		result := CO2eNdirecttDefault(
			decimal.NewFromFloat(tt.massSynthFertz),
			decimal.NewFromFloat(tt.massOrgFertz))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestCO2eNindirectt(t *testing.T) {
	type Test struct {
		nfertVolatIT, nfertLeachIT float64
		result                     float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 24},
	}
	for i, tt := range tests {
		result := CO2eNindirectt(
			decimal.NewFromFloat(tt.nfertVolatIT),
			decimal.NewFromFloat(tt.nfertLeachIT))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestNfertVolatIT(t *testing.T) {
	type Test struct {
		massSynthFertz, nContSynthFertz, massOrgFertz, nContOrgFertz, allFractSynth, allFractOrg, nitrOxdEmissWS, gWarmingPotentl float64
		result                                                                                                                    float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 12, 12, 12, 12, 12, 12, 782043.429},
	}
	for i, tt := range tests {
		result := NfertVolatIT(
			decimal.NewFromFloat(tt.massSynthFertz),
			decimal.NewFromFloat(tt.nContSynthFertz),
			decimal.NewFromFloat(tt.massOrgFertz),
			decimal.NewFromFloat(tt.nContOrgFertz),
			decimal.NewFromFloat(tt.allFractSynth),
			decimal.NewFromFloat(tt.allFractOrg),
			decimal.NewFromFloat(tt.nitrOxdEmissWS),
			decimal.NewFromFloat(tt.gWarmingPotentl))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestNfertVolatITDefault(t *testing.T) {
	type Test struct {
		massSynthFertz, massOrgFertz float64
		result                       float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 21.987},
	}
	for i, tt := range tests {
		result := NfertVolatITDefault(
			decimal.NewFromFloat(tt.massSynthFertz),
			decimal.NewFromFloat(tt.massOrgFertz))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestNfertLeachIT(t *testing.T) {
	type Test struct {
		massSynthFertz, nContSynthFertz, massOrgFertz, nContOrgFertz, nFractSoil, nitrOxdEmissLR, gWarmingPotentl float64
		result                                                                                                    float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 12, 12, 12, 12, 12, 782043.429},
	}
	for i, tt := range tests {
		result := NfertLeachIT(
			decimal.NewFromFloat(tt.massSynthFertz),
			decimal.NewFromFloat(tt.nContSynthFertz),
			decimal.NewFromFloat(tt.massOrgFertz),
			decimal.NewFromFloat(tt.nContOrgFertz),
			decimal.NewFromFloat(tt.nFractSoil),
			decimal.NewFromFloat(tt.nitrOxdEmissLR),
			decimal.NewFromFloat(tt.gWarmingPotentl))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestNfertLeachITDefault(t *testing.T) {
	type Test struct {
		massSynthFertz, massOrgFertz float64
		result                       float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 24.736},
	}
	for i, tt := range tests {
		result := NfertLeachITDefault(
			decimal.NewFromFloat(tt.massSynthFertz),
			decimal.NewFromFloat(tt.massOrgFertz))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestNetEmissionsRemoval(t *testing.T) {
	type Test struct {
		cTotalCarbon, baseline, leakeage float64
		emissions                        []float64
		result                           float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 22, []float64{}, -264},
	}
	for i, tt := range tests {
		emms := []decimal.Decimal{}
		for _, val := range tt.emissions {
			emms = append(emms, decimal.NewFromFloat(val))
		}
		result := NetEmissionsRemoval(
			decimal.NewFromFloat(tt.cTotalCarbon),
			decimal.NewFromFloat(tt.baseline),
			decimal.NewFromFloat(tt.leakeage),
			emms...)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}
