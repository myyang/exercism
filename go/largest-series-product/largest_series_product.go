// Package lsproduct provides function to compute largest product
package lsproduct

import "strconv"

const testVersion = 3

type lspError struct{}

// Error returns naive string
func (err lspError) Error() string {
	return "Larges series product error"
}

// LargestSeriesProduct return largest product value
func LargestSeriesProduct(digits string, span int) (int64, error) {
	var (
		maxPro int64
	)
	if span < 0 || span > len(digits) {
		return -1, lspError{}
	}
	for k := 0; k < len(digits)-span+1; k++ {
		n, pass := int64(1), false
		for _, y := range digits[k : k+span] {
			i, err := strconv.Atoi(string(y))
			if err != nil {
				return int64(-1), lspError{}
			}
			if i == 0 {
				pass = true
				break
			}
			n *= int64(i)
		}
		if !pass && n > maxPro {
			maxPro = n
		}
	}
	return maxPro, nil
}
