package functions

import "math"

// หารปัดเศษทิ้ง
func Quotient(set, div, decimal float64) float64 {
	return math.Floor(Divine(set, div, decimal))
}
