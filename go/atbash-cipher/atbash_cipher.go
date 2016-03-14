// Package atbash provides ...
package atbash

import (
	"strings"
)

var (
	mapString string = "zyxwvutsrqponmlkjihgfedcba"
)

func Atbash(orig string) string {
	r, group := "", 0
	orig = strings.ToLower(cTrim(orig, " ,."))
	for _, c := range orig {
		if x := string(c)[0] - "a"[0]; 0 <= x && x < 26 {
			r += mapString[x : x+1]
		} else {
			r += string(c)
		}
		group++
		if group == 5 {
			r += " "
			group = 0
		}
	}
	return strings.Trim(r, " ")
}

func cTrim(s string, cutset string) string {
	r := s
	for _, c := range cutset {
		r = strings.Replace(r, string(c), "", -1)
	}
	return r
}
