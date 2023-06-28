package functions

func RoundStang(value float64) int {
	if value >= 1 && value < 13 {
		return 0
	} else if value < 38 {
		return 25
	} else if value < 62 {
		return 50
	} else if value < 88 {
		return 75
	} else {
		return 100
	}
}
