// Package prime provides function to get prime number
package prime

import (
	"math"
)

const testVersion = 2

// Factors prime factorize number
func Factors(num int64) []int64 {
	r := []int64{}

	for num%2 == 0 && num > 0 {
		r = append(r, 2)
		num = num / 2
	}

	for i := int64(3); float64(i) <= math.Sqrt(float64(num)); i = i + 2 {
		for num%i == 0 {
			r = append(r, i)
			num = num / i
		}
	}
	if num > 2 {
		r = append(r, num)
	}
	return r
}
