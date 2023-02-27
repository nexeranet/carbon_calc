package carbon_calc

import (
	"fmt"
	"strconv"
	"testing"
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
		result := OCCBufferPool(tt.minted, tt.percent)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
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
		result := OCCHolders(tt.minted, tt.percent)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
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
		result := OCCMintedPerMonitoringZone(tt.minted, tt.carbon, tt.zone)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
		}
	}
}

// TODO: test with real arguments
func TestMintedOCC(t *testing.T) {
	type Test struct {
		current, previous float64
		result               float64 // precision = 3
	}
	tests := []Test{
		{24, 12, 12},
	}
	for i, tt := range tests {
		result := MintedOCC(tt.current, tt.previous)
		rounded, err := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		if err != nil {
			t.Fatal(err)
		}
		if rounded != tt.result {
			t.Fatalf("Test number %d, expect: %f, have: %f", i, tt.result, result)
		}
	}
}
