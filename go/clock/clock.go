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
type Clock int

// New a value of type Clock
func New(hour, minute int) Clock {
	return Clock(hour%24*60 + minute).Add(0)
}

// String format into "HH:MM"
func (clk Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clk/60, clk%60)
}

// Add minutes to affect value
func (clk Clock) Add(minutes int) Clock {
	mins := int(clk) + minutes
	for mins < 0 {
		mins += minutesADay
	}

	for mins >= minutesADay {
		mins -= minutesADay
	}
	return Clock(mins)
}
