// Package raindrops provides convertion
package raindrops

import (
	"fmt"
)

const testVersion = 2

var (
	smap = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	checkFacs = []int{3, 5, 7}
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

func facStr(i, fac int) string {
	if i%fac == 0 {
		return smap[fac]
	}
	return ""
}
