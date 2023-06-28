package functions

import "math"

// ยกกำลัง
func Pow(a, b, decimal float64) float64 {
	res := 0.0
	if b == 10.0 {
		res = math.Pow10(int(a))
	} else {
		res = math.Pow(a, b)
	}
	if decimal == 0 {
		return res
	}
	return DecimalPoint(res, decimal)
}
