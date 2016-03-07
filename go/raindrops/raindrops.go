// Package raindrops provides convertion
package raindrops

import (
	"fmt"
)

const testVersion = 2

type pair struct {
	num int
	str string
}

var (
	primePairs = []pair{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}
)

// Convert int to string
func Convert(i int) string {
	s := ""

	for _, v := range primePairs {
		s += facStr(i, v)
	}

	if s == "" {
		return fmt.Sprintf("%d", i)
	}

	return s
}

func facStr(i int, fac pair) string {
	if i%fac.num == 0 {
		return fac.str
	}
	return ""
}
