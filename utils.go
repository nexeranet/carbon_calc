package carbon_calc

import "math"


func Circumference(r float64) float64{
    return 2 * math.Pi * r
}

func CircleArea(r float64) float64 {
    return math.Pi * math.Pow(r ,2)
}
