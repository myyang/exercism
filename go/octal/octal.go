// Package octal provides function to parse octal
package octal

import (
	"errors"
	"strconv"
)

// ParseOctal return decimal value
func ParseOctal(input string) (int64, error) {
	mul, sum := int64(1), int64(0)
	for i := len(input) - 1; i >= 0; i-- {
		digit := input[i:]
		x, err := strconv.Atoi(digit)
		if err != nil || !isValid(digit) {
			return 0, errors.New("Convert or overflow")
		}
		sum += int64(x) * mul
		mul, input = mul*8, input[:i]
	}
	return sum, nil
}

func isValid(c string) bool {
	x := c[0] - "0"[0]
	return x >= 0 && x < 8
}
