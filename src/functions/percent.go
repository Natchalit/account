package functions

// Percent
func Percent(num, float float64) float64 {
	return DecimalPoint((num / 100.0), float)
}
