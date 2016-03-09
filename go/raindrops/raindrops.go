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
	s := ""

	for _, v := range primePairs {
		if i%v.num == 0 {
			s += v.str
		}
	}

	if s == "" {
		return fmt.Sprintf("%d", i)
	}

	return s
}
