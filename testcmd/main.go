package main

import (
	"fmt"
	"time"
)

func main() {
	val, err := time.Parse(time.RFC3339, "2029-07-23T00:46:03+00:00")
	// val, err := time.Parse(time.RFC3339, "2022-01-20T21:45:00+00:00")
	if err != nil {
		panic(err)
	}

	now := time.Now()
	fmt.Printf("NOW: %s\n", now.UTC())
	fmt.Printf("END: %s\n", val.UTC())
	delta := val.Sub(now)

	dayDuration := delta.Truncate(time.Hour * 24)
	partialDuration := delta - dayDuration
	hourDuration := partialDuration.Truncate(time.Hour)
	partialDuration -= hourDuration
	minDuration := partialDuration.Truncate(time.Minute)
	partialDuration -= minDuration

	fmt.Printf("%d days %d hours %d minutes %.0f seconds\n", dayDuration/time.Hour/24, hourDuration/time.Hour, minDuration/time.Minute, partialDuration.Seconds())

	nowYear := now.UTC().Year()
	nowDoy := now.UTC().YearDay()

	endYear := val.UTC().Year()
	endDoy := val.UTC().YearDay()

	eod := time.Date(nowYear, now.UTC().Month(), now.UTC().Day()+1, 0, 0, 0, 0, time.UTC)
	tteod := eod.Sub(now)
	hmsDur := delta - delta.Truncate(time.Hour*24)
	fmt.Printf("To EOD: %s, deltaTrunc: %s\n", tteod, delta-delta.Truncate(time.Hour*24))

	if hmsDur > tteod {
		nowDoy += 1
	}

	var years, days int

	if endDoy >= nowDoy {
		years = endYear - nowYear
		days = endDoy - nowDoy
	} else {
		eoy := time.Date(nowYear, time.December, 31, 12, 0, 0, 0, time.UTC)
		yearDays := eoy.YearDay()
		days = endDoy + yearDays - nowDoy
		years = endYear - nowYear - 1
	}
	fmt.Printf("Today: Year: %d, Day of Year: %d\n", now.UTC().Year(), now.UTC().YearDay())
	fmt.Printf("End  : Year: %d, Day of Year: %d\n", val.UTC().Year(), val.UTC().YearDay())
	// fmt.Printf("Diff : %d years %d days\n", years, days)

	fmt.Printf("Diff : %d years %d days ", years, days)

	msDur := hmsDur - hmsDur.Truncate(time.Hour)
	sDur := msDur - msDur.Truncate(time.Minute)
	fmt.Printf("%d hours %d minutes %d seconds\n", hmsDur/time.Hour, msDur/time.Minute, sDur/time.Second)
}
