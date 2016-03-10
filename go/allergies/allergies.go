// Package allergies provides ...
package allergies

import (
	"strconv"
)

var scoremap = map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// Allergies accept scores and return allergies list
func Allergies(input int) []string {
	binStr, mul, arr := strconv.FormatInt(int64(input), 2), 1, []string{}
	for i := len(binStr) - 1; i > -1; i-- {
		ind, _ := strconv.Atoi(binStr[i : i+1])
		str, ok := scoremap[ind*mul]
		if ok {
			arr = append(arr, str)
		}
		mul = mul << 1
	}
	return arr
}

// AllergicTo test wheather given allergen in target score
func AllergicTo(input int, allergen string) bool {
	arr := Allergies(input)
	for _, v := range arr {
		if v == allergen {
			return true
		}
	}
	return false
}
