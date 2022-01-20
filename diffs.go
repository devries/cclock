package main

import (
	"time"
)

type ClockDifference struct {
	Years   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func getDifference(start, end time.Time) ClockDifference {
	var ret ClockDifference

	delta := end.Sub(start)

	startYear := start.UTC().Year()
	startDoy := start.UTC().YearDay()

	endYear := end.UTC().Year()
	endDoy := end.UTC().YearDay()

	eod := time.Date(startYear, start.UTC().Month(), start.UTC().Day()+1, 0, 0, 0, 0, time.UTC)
	tteod := eod.Sub(start)
	hmsDur := delta - delta.Truncate(time.Hour*24)

	if hmsDur >= tteod {
		startDoy += 1
	}

	if endDoy >= startDoy {
		ret.Years = endYear - startYear
		ret.Days = endDoy - startDoy
	} else {
		eoy := time.Date(startYear, time.December, 31, 12, 0, 0, 0, time.UTC)
		yearDays := eoy.YearDay()
		ret.Years = endYear - startYear - 1
		ret.Days = endDoy + yearDays - startDoy
	}

	msDur := hmsDur - hmsDur.Truncate(time.Hour)
	sDur := msDur - msDur.Truncate(time.Minute)

	ret.Hours = int(hmsDur / time.Hour)
	ret.Minutes = int(msDur / time.Minute)
	ret.Seconds = int(sDur / time.Second)

	return ret
}
