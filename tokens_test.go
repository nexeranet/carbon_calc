package carbon_calc

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestMintedOCC(t *testing.T) {
	type Test struct {
		current, previous float64
		result            float64 // precision = 3
	}
	tests := []Test{
		{24, 12, 12},
		{6.27, 2.64, 3.63},
	}
	for i, tt := range tests {
		result := MintedOCC(
			decimal.NewFromFloat(tt.current),
			decimal.NewFromFloat(tt.previous))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestOCCBufferPool(t *testing.T) {
	type Test struct {
		minted, percent float64
		result          float64 // precision = 3
	}
	tests := []Test{
		{12, 0, 0.84},
		{12, 0.07, 0.84},
		{12, 0.33, 3.96},
		{3.6, 0.07, 0.252},
	}
	for i, tt := range tests {
		result := OCCBufferPool(decimal.NewFromFloat(tt.minted), tt.percent)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestOCCHolders(t *testing.T) {
	type Test struct {
		minted, percent float64
		result          float64 // precision = 3
	}
	tests := []Test{
		{12, 0, 0.96},
		{12, 0.08, 0.96},
		{12, 0.33, 3.96},
		{3.6, 0.08, 0.288},
	}
	for i, tt := range tests {
		result := OCCHolders(decimal.NewFromFloat(tt.minted), tt.percent)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

func TestOCCMintedPerMonitoringZone(t *testing.T) {
	type Test struct {
		minted, carbonC, zoneC, carbonP, zoneP float64
		result                                 float64 // precision = 3
	}
	tests := []Test{
		{3.6, 6.6, 6.041, 2.77, 2.54, 3.2908},
		{3.6, 6.6, 0.558, 2.77, 0.234, 0.3045},
	}
	for i, tt := range tests {
		result := OCCMintedPerMonitoringZone(
			decimal.NewFromFloat(tt.minted),
			decimal.NewFromFloat(tt.carbonC),
			decimal.NewFromFloat(tt.zoneC),
			decimal.NewFromFloat(tt.carbonP),
			decimal.NewFromFloat(tt.zoneP))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}
