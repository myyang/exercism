/*
Package scrabble provides function to count score of sentenses or word

Score table as following:

	Letter                           Value
	A, E, I, O, U, L, N, R, S, T       1
	D, G                               2
	B, C, M, P                         3
	F, H, V, W, Y                      4
	K                                  5
	J, X                               8
	Q, Z                               10

*/
package scrabble

import (
	"log"
	"regexp"
	"strings"
)

const targetTestVersion = 3

var scoreMap = map[string]int{
	"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
	"d": 2, "g": 2,
	"b": 3, "c": 3, "m": 3, "p": 3,
	"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
	"k": 5,
	"j": 8, "x": 8,
	"q": 10, "z": 10,
}

// Score the input string, returns int scores
func Score(s string) int {
	re, err := regexp.Compile("[a-z]*")
	if err != nil {
		log.Println("error re:", err)
		return -1
	}

	match := strings.Join(re.FindAllString(strings.ToLower(s), -1), "")

	if match == "" {
		return 0
	}

	score := 0
	for _, v := range match {
		score += scoreMap[string(v)]
	}
	return score
}
