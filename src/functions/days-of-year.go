package functions

import (
	"time"
)

var AllMonth = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func DaysOfYear() int {
	daysOfYear := 0
	year := time.Now().Year()
	for _, vmonth := range AllMonth {
		daysOfYear += DayOfMonth(year, time.Month(vmonth))
	}

	return daysOfYear
}
