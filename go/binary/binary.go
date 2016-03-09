// Package binary provides binary string paser function
package binary

import "strconv"

// ParseBinary return result interger
func ParseBinary(bin string) (int, error) {
	mul, result := 1, 0
	for i := len(bin) - 1; i > -1; i-- {
		x, _ := strconv.Atoi(bin[i : i+1])
		result += x * mul
		mul = mul << 1
	}
	return result, nil
}
