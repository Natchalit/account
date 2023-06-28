package functions

// ผลรวมทั้งหมด
func SumTotal(nums []float64, float float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return DecimalPoint(total, float)
}
