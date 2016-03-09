// Package anagram provides function to check anagram
package anagram

import (
	"reflect"
	"strings"
)

// FreqMap and Frequency are copied from ex:parallel letter frequency
// Do not reinvent the wheel
type FreqMap map[rune]int

// Frequency counts the character frequency
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Detect use letter frequency to detect anagram
func Detect(subject string, candidates []string) []string {
	freq, results := Frequency(strings.ToLower(subject)), []string{}
	for _, c := range candidates {
		c = strings.ToLower(c)
		tmpFreq := Frequency(c)
		if reflect.DeepEqual(freq, tmpFreq) && c != subject {
			results = append(results, c)
		}
	}
	return results
}
