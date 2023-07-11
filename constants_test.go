package carbon_calc

import (
	"testing"
)

func TestGetRainfallType(t *testing.T) {
	type Test struct {
		result RainfallType
		amount int64
	}
	tests := []Test{
		{RainfallTypeDry, 999},
		{RainfallTypeDry, 1000},
		{RainfallTypeMoist, 1400},
		{RainfallTypeMoist, 1001},
		{RainfallTypeMoist, 1999},
		{RainfallTypeWet, 2400},
		{RainfallTypeWet, 2000},
	}

	for i, tt := range tests {
		result := GetRainfallType(tt.amount)
		if result != tt.result {
			t.Fatalf("Test number %d, expect: %d, have: %d", i, tt.result, result)
		}
	}
}
