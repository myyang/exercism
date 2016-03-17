// Package dna provides count dna factor
package dna

import (
	"errors"
)

// Series string
type Series string

// DNA return new series
func DNA(s string) Series {
	return Series(s)
}

// Count given factor
func (ser Series) Count(s byte) (int, error) {
	if s != 'A' && s != 'C' && s != 'T' && s != 'G' {
		return 0, errors.New("Not valid charater")
	}
	var count int
	for _, c := range ser {
		if byte(c) == s {
			count++
		}
	}
	return count, nil
}

// Histogram map
type Histogram map[rune]int

// Counts all factor
func (ser Series) Counts() Histogram {
	h := Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0}
	for k, _ := range h {
		c, err := ser.Count(byte(k))
		if err != nil {
			return Histogram{}
		}
		h[k] = c
	}
	return h
}
