// Package clock provides a simple naive clock
package clock

import (
	"fmt"
)

const (
	testVersion = 3
	minutesADay = 24 * 60
)

// Clock is a time-only naive clock
type Clock struct {
	minute int
}

// New a clock instance
func New(hour, minute int) Clock {
	clk := Clock{hour*60 + minute}
	clk.refresh()
	return clk
}

// String format into "HH:MM"
func (clk Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clk.minute/60%24, clk.minute%60)
}

// Add minutes to affect value
func (clk Clock) Add(minutes int) Clock {
	clk.minute += minutes
	clk.refresh()
	return clk
}

func (clk *Clock) refresh() {
	for clk.minute < 0 {
		clk.minute += minutesADay
	}

	for clk.minute > minutesADay {
		clk.minute -= minutesADay
	}
}
