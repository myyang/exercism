// Package strand provides ...
package strand

import (
	"regexp"
	"strings"
)

const testVersion = 3

var (
	rnaMap = map[string]string{
		"G": "c",
		"C": "g",
		"T": "a",
		"A": "u",
	}
)

// ToRNA string from dna string
func ToRNA(dna string) string {

	for k, v := range rnaMap {
		re, err := regexp.Compile(k)
		if err != nil {
			return ""
		}
		dna = re.ReplaceAllString(dna, v)
	}
	return strings.ToUpper(dna)
}
