// Package say provides convert function
package say

import (
	"strings"
)

var (
	numMap = map[uint64]string{
		0:    "zero",
		1:    "one",
		2:    "two",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "ninteen",
		20:   "twenty",
		30:   "thirty",
		40:   "forty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		1e2:  "hundred",
		1e3:  "thousand",
		1e6:  "million",
		1e9:  "billion",
		1e12: "trillion",
		1e15: "quadrillion",
		1e18: "quintillion",
	}
)

// Say string of number
func Say(num uint64) string {
	if num == 0 {
		return convert(num)
	}

	var r string
	k := uint64(1)
	for num != 0 {
		x, tmp := num%uint64(1e3), ""
		if x != 0 {
			tmp = convert(x)
			if k > 1 {
				tmp = tmp + " " + numMap[k]
			}
			r = tmp + " " + r
		}
		num, k = num/1e3, k*1e3
	}
	return strings.TrimRight(r, " ")
}

func convert(num uint64) string {
	var r string
	div, rem := num/100, num%100
	if rem < 20 {
		r += numMap[rem]
	} else {
		l := rem % 10
		r += numMap[rem-l]
		if l > 0 {
			r += "-" + numMap[l]
		}
	}

	if div > 0 && rem != 0 {
		r = numMap[div] + " " + numMap[100] + " " + r
	} else if div > 0 && rem == 0 {
		r = numMap[div] + " " + numMap[100]
	}
	return r
}
