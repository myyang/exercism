// Package igpay provides provide naive translation
package igpay

import (
	"strings"
)

var (
	vowelList      = []string{"a", "e", "i", "o", "u"}
	vowelSoundList = []string{"y", "x"}
)

// PigLatin is naive traslation function
func PigLatin(in string) string {
	words := strings.Split(in, " ")
	r := ""
	for _, w := range words {
		if startsWithVowel(w) || startsWithVowelSound(w) {
			r += (w + "ay")
		} else {
			i := indexFirstVowelSound(w)
			r += (w[i:] + w[:i] + "ay")
		}
		r += " "
	}
	return r[:len(r)-1]
}

func startsWithVowel(w string) bool {
	for _, c := range vowelList {
		if w[:1] == c {
			return true
		}
	}
	return false
}

func startsWithVowelSound(w string) bool {
	for _, c := range vowelSoundList {
		if w[:1] == c && !startsWithVowel(w[1:]) {
			return true
		}
	}
	return false
}

func indexFirstVowelSound(w string) int {
	exceptList := []string{"qu"}

	for i, c := range w {
		if startsWithVowel(string(c)) {
			skip := false
			for _, exw := range exceptList {
				if x := strings.Index(w, exw); i < x+len(exw) {
					skip = true
				}
			}
			if !skip {
				return i
			}
		}
	}
	return -1
}
