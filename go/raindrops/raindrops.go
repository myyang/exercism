// Package raindrops provides convertion
package raindrops

import (
	"fmt"
)

const testVersion = 2

// Pair defines primary number and its mapping string
type Pair struct {
	num int
	str string
}

var (
	checkFacs = []Pair{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}
)

// Convert int to string
func Convert(i int) string {
	s := ""

	for _, v := range checkFacs {
		s += facStr(i, v)
	}

	if s == "" {
		return fmt.Sprintf("%d", i)
	}

	return s
}

func facStr(i int, fac Pair) string {
	if i%fac.num == 0 {
		return fac.str
	}
	return ""
}
