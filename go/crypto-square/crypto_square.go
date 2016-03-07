// Package cryptosquare provides square encryption function
package cryptosquare

import (
	"log"
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

// Encode string to square code
func Encode(s string) string {
	re, err := regexp.Compile("[[:alnum:]]*")
	if err != nil {
		log.Println("error re:", err)
		return ""
	}

	match := strings.Join(re.FindAllString(strings.ToLower(s), -1), "")
	if match == "" {
		return ""
	}

	slength := float64(len(match))
	slicen := int(math.Ceil(math.Sqrt(slength)))
	result, itern := "", int(slength)/slicen
	for i := 0; i < slicen; i++ {
		for j := 0; j < itern; j++ {
			result += string(match[i+j*slicen])
		}
		if i+itern*slicen < int(slength) {
			result += string(match[i+itern*slicen])
		}
		result += " "
	}

	return result[0 : len(result)-1]
}
