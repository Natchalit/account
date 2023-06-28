package functions

import "time"

func DayOfMonth(year int, month time.Month) int {

	if month == 1 {
		year = year + 1
	}
	return time.Date(year, month, 0, 10, 0, 0, 0, time.Local).Day()

}
