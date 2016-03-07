/*
Package clock provides a simple naive clock that
can be initialized and adjusted, besides, it uses
24-hour time string while printing.
*/
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

// New a value of type Clock
func New(hour, minute int) Clock {
	clk := Clock{hour*60 + minute}
	return clk.refresh()
}

// String format into "HH:MM"
func (clk Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clk.minute/60, clk.minute%60)
}

// Add minutes to affect value
func (clk Clock) Add(minutes int) Clock {
	clk.minute += minutes
	return clk.refresh()
}

func (clk Clock) refresh() Clock {
	for clk.minute < 0 {
		clk.minute += minutesADay
	}

	for clk.minute >= minutesADay {
		clk.minute -= minutesADay
	}
	return clk
}
