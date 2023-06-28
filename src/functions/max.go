package functions

// หาค่า Max
func Max(nums []float64) float64 {
	var max *float64
	for _, num := range nums {
		if max == nil {
			max = &num
		}
		if num > *max {
			max = &num
		}
	}

	return *max
}
