package carbon_calc

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestOCCBufferPool(t *testing.T) {
	type Test struct {
		minted, percent float64
		result          float64 // precision = 3
	}
	tests := []Test{
		{12, 0, 0.84},
		{12, 0.07, 0.84},
		{12, 0.33, 3.96},
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

// TODO: test with real arguments
func TestOCCMintedPerMonitoringZone(t *testing.T) {
	type Test struct {
		minted, carbon, zone float64
		result               float64 // precision = 3
	}
	tests := []Test{
		{12, 12, 12, 12},
	}
	for i, tt := range tests {
		result := OCCMintedPerMonitoringZone(
			decimal.NewFromFloat(tt.minted),
			decimal.NewFromFloat(tt.carbon),
			decimal.NewFromFloat(tt.zone))
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result.InexactFloat64()), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result.InexactFloat64())
		}
	}
}

// TODO: test with real arguments
func TestMintedOCC(t *testing.T) {
	type Test struct {
		current, previous float64
		result            float64 // precision = 3
	}
	tests := []Test{
		{24, 12, 12},
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
