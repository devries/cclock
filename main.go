package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	initialize()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGWINCH)

	runClock(end, sigChan)
}

func runClock(end time.Time, sigChan <-chan os.Signal) {
	r, c := resize()
	clear()
	hideCursor()

	top := r/2 - 2
	left := c/2 - 37
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

		select {
		case s := <-sigChan:
			if s == syscall.SIGWINCH {
				runClock(end, sigChan)
			} else {
				move(r-1, 0)
				cleanup()
				os.Exit(0)
			}
		case <-time.After(50 * time.Millisecond):
		}
	}
}
