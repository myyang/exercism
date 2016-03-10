// Package romannumerals provides ...
package romannumerals

import (
	"errors"
	"strings"
)

const testVersion = 2

// RomanLevels defines roman value and string
type RomanLevel struct {
	Unit, Five, Ten string
	Value           int
}

var (
	UNIT = RomanLevel{"I", "V", "X", 1}
	TENS = RomanLevel{"X", "L", "C", 10}
	HUND = RomanLevel{"C", "D", "M", 100}
	THOU = "M"
)

// ToRomanNumeral converts integer to string
func ToRomanNumeral(in int) (string, error) {

	if 1 > in || in > 3999 {
		return "", errors.New("Value should be 1~3999")
	}

	r, enStr, val, rem := "", "", in/1000, in%1000
	r += strings.Repeat(THOU, val)
	for _, roi := range []RomanLevel{HUND, TENS, UNIT} {
		enStr, rem = conv(rem, roi)
		r += enStr
	}
	return r, nil
}

func conv(v int, roi RomanLevel) (r string, rem int) {

	b, x := v/roi.Value, v%roi.Value
	switch b {
	case 0, 1, 2, 3:
		return strings.Repeat(roi.Unit, b), x
	case 4:
		return roi.Unit + roi.Five, x
	case 5, 6, 7, 8:
		return roi.Five + strings.Repeat(roi.Unit, (b-5)), x
	case 9:
		return roi.Unit + roi.Ten, x
	}
	return r, rem
}
