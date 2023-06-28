package functions

// หาค่า Min
func Min(nums []float64) float64 {

	var min *float64
	for _, num := range nums {
		if min == nil {
			min = &num
		}
		if num < *min {
			min = &num
		}
	}

	return *min
}
