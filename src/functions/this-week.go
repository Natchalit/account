package functions

import (
	"fmt"
	"time"
)

func ThisWeek() int {
	tn := time.Now().Local()
	test := tn.YearDay()
	fmt.Printf(`%d`, test)
	_, week := tn.ISOWeek()
	return week
}

func ThisWeekbyTime(tn time.Time) int {
	_, week := tn.Local().ISOWeek()
	return week
}
