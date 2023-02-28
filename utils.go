package carbon_calc

import (
	"math"

	"github.com/shopspring/decimal"
)

func Radius(circumference float64) float64 {
	return circumference / (2 * math.Pi)
}

func Circumference(r float64) float64 {
	return 2 * math.Pi * r
}

func CircleArea(r decimal.Decimal) decimal.Decimal {
	return decimal.NewFromFloat(math.Pi).Mul(r.Pow(decimal.New(2, 0)))
}

func Sum(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum
}
