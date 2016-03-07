/*
Package luhn provides Luhn algorithm validation and encrytion functions

for more Luhn detail, please check http://en.wikipedia.org/wiki/Luhn_algorithm
*/
package luhn

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Valid the given string that is a legal/proper luhn string
func Valid(s string) bool {
	if len(s) <= 0 {
		return false
	}
	return s == AddCheck(s[:len(s)-1])
}

// AddCheck return appending chechsum to given string
func AddCheck(raw string) string {
	re, err := regexp.Compile("[[:digit:]]*")
	if err != nil {
		log.Fatal("Regex error:", err)
		return ""
	}

	s := strings.Join(re.FindAllString(strings.ToLower(raw), -1), "")
	if s == "" {
		return ""
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Error atoi: ", err)
		return ""
	}

	checks, mul := 0, true
	for i := v; i > 0; i /= 10 {
		var x int
		if mul {
			x = (i % 10) * 2
			for x >= 10 {
				x = x/10 + x%10
			}
		} else {
			x = i % 10
		}
		mul = !mul
		checks += x
	}
	checks = checks * 9 % 10
	return fmt.Sprint(raw, strconv.Itoa(checks))
}
