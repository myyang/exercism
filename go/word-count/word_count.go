package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Frequency defines return type.
type Frequency map[string]int

func mapper(out chan Frequency, phrase string) {
	re, err := regexp.Compile("[[:alnum:]]*")
	if err != nil {
		close(out)
	}
	for _, s := range re.FindAllString(strings.ToLower(phrase), -1) {
		if s != "" {
			out <- Frequency{s: 1}
		}
	}
	close(out)
}

func reducer(in chan Frequency, out chan Frequency) {
	fm := Frequency{}
	for f := range in {
		for k, v := range f {
			fm[k] += v
		}
	}
	out <- fm
}

// WordCount run with concurrent mapper reducer
func WordCount(phrase string) Frequency {
	pipe, result := make(chan Frequency), make(chan Frequency)
	go mapper(pipe, phrase)
	go reducer(pipe, result)
	return <-result
}
