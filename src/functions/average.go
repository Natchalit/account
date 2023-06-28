package functions

// หาค่าเฉลี่ย
func Average(numbers []float64, decimal float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	avg := sum / float64(len(numbers))
	return DecimalPoint(avg, decimal)
}
