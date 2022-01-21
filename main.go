package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	resp, err := query()
	if err != nil {
		panic(err)
	}

	if resp.Status != "success" {
		fmt.Printf("Query failure: %s\n", resp.Status)
		os.Exit(1)
	}

	end := resp.Data.Modules["carbon_deadline_1"].Timestamp

	r, c := initialize()
	clear()
	hideCursor()

	top := r/2 - 2
	left := c/2 - 39
	move(top, left)
	fmt.Printf("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	move(top+1, left)
	fmt.Printf("┃ Time to act before we reach irreversible 1.5\u00b0C global temperature rise ┃")
	move(top+2, left)
	fmt.Printf("┃           X years, XXX days, XX hours, XX minutes, XX seconds          ┃")
	move(top+3, left)
	fmt.Printf("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")

	yrpos := left + 12
	daypos := left + 21
	hourpos := left + 31
	minutepos := left + 41
	secondpos := left + 53

	interruptHandling(r-1, 0) // Clean up and end on SIGTERM

	for {
		now := time.Now()
		cdur := getDifference(now, end)

		move(top+2, yrpos)
		fmt.Printf("%d", cdur.Years)
		move(top+2, daypos)
		fmt.Printf("%3d", cdur.Days)
		move(top+2, hourpos)
		fmt.Printf("%2d", cdur.Hours)
		move(top+2, minutepos)
		fmt.Printf("%2d", cdur.Minutes)
		move(top+2, secondpos)
		fmt.Printf("%2d", cdur.Seconds)

		time.Sleep(50 * time.Millisecond)
	}

	// move(r-1, 0)
	// cleanup()
}
