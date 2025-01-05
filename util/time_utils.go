package util

import (
	"fmt"
	"time"
)

//see LocalDateTimes.java
func AfterBusinessDays(date time.Time, days int) (res time.Time, err error) {
	if days < 0 {
		err = fmt.Errorf("days parameter is less than zero")
		return
	}
	// equivalent to do { ... } while (days >= 0)
	for days > 0 {
		// start with tomorrow
		date = date.AddDate(0, 0, 1)
		// check if weekend
		// see: https://golang.org/pkg/time/#Weekday
		// Sunday = 0, Saturday = 6
		weekDay := date.Weekday()
		// TODO: check also that the date isn't one of finnish holidays (see HolidayUtil.java)
		if weekDay != time.Saturday && weekDay != time.Sunday {
			days--
		}
	}
	res = date
	return
}

// see TimeUtil.java
func FormatDMY(date time.Time) (res string) {
	day := date.Day()
	month := date.Month()
	year := date.Year()
	res = fmt.Sprintf("%02d.%02d.%d", day, month, year)
	return
}

// see TimeUtil.java
func FormatHM(date time.Time) (res string) {
	hour := date.Hour()
	minute := date.Minute()
	res = fmt.Sprintf("%02d:%02d", hour, minute)
	return
}
