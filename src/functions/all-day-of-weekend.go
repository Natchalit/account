package functions

import (
	"time"
)

func AllDayOfWeekend() int {
	tn := time.Now().Local()
	return CountWeekEnd(tn)
}

func CountWeekEnd(tn time.Time) int {
	year := tn.Year()
	countWeekEnd := 0

	for _, vMonth := range AllMonth {
		for i := 0; i < DayOfMonth(year, time.Month(vMonth)); i++ {
			day := time.Date(year, time.Month(vMonth), 0, 10, 0, 0, 0, time.Local).Format("Mon")
			if day == "Sat" || day == "Sun" {
				countWeekEnd += 1
			}
		}
	}

	return countWeekEnd
}
