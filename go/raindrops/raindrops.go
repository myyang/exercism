// Package raindrops provides convertion
package raindrops

import (
	"fmt"
)

const testVersion = 2

var (
	primePairs = []struct {
		num int
		str string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}
)

// Convert int to string
func Convert(i int) string {
	var s string

	for _, v := range primePairs {
		if i%v.num == 0 {
			s += v.str
		}
	}

	if s == "" {
		return fmt.Sprint(i)
	}

	return s
}
