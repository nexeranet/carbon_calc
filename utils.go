package carbon_calc

import "math"

func Radius(circumference float64) float64 {
	return circumference / (2 * math.Pi)
}

func Circumference(r float64) float64 {
	return 2 * math.Pi * r
}

func CircleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func Sum(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum
}
