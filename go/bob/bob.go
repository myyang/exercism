package bob // package name must match the package name in bob_test.go

import (
	"log"
	"regexp"
	"strings"
)

const testVersion = 2 // same as targetTestVersion

func isShout(msg string) bool {
	re, err := regexp.Compile(`[\pL[:alpha:]]*`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	s := strings.Join(re.FindAllString(msg, -1), "")
	if s != "" && strings.ToUpper(s) == s {
		return true
	}
	return false
}

func isQuestion(msg string) bool {
	s := trimBlank(msg)
	if len(s) > 0 && s[len(s)-1:] == "?" {
		return true
	}
	return false
}

func isEmpty(msg string) bool {
	if trimBlank(msg) == "" {
		return true
	}
	return false
}

func trimBlank(msg string) string {
	re, err := regexp.Compile(`[^\pZz[:space:]]*`)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return strings.Join(re.FindAllString(msg, -1), "")
}

// Hey bob and get response
func Hey(msg string) string {
	if isShout(msg) {
		return "Whoa, chill out!"
	} else if isQuestion(msg) {
		return "Sure."
	} else if isEmpty(msg) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
