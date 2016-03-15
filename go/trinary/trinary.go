package trinary

import (
	"errors"
	"strconv"
)

func ParseTrinary(s string) (int64, error) {
	if len(s) < 1 {
		return 0, errors.New("String length error.")
	}
	mul, sum := int64(1), int64(0)
	for i := int64(len(s) - 1); i >= 0; i-- {
		num, err := strconv.Atoi(s[i:])
		if err != nil {
			return 0, errors.New("Atoi error")
		}
		sum += mul * int64(num)
		s = s[:i]
		mul = mul * 3
	}
	if sum < 0 {
		return 0, errors.New("overflow")
	}

	return sum, nil
}
