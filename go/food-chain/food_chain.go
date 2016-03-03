// Package foodchain provides a cumulative song functions
package foodchain

import (
	"fmt"
)

const testVersion = 2

var (
	swallowm = map[string]string{
		"cow":    "She swallowed the cow to catch the goat.",
		"goat":   "She swallowed the goat to catch the dog.",
		"dog":    "She swallowed the dog to catch the cat.",
		"cat":    "She swallowed the cat to catch the bird.",
		"bird":   "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
		"spider": "She swallowed the spider to catch the fly.",
		"fly":    "I don't know why she swallowed the fly. Perhaps she'll die.",
	}
	yellm = map[string]string{
		"cow":    "I don't know how she swallowed a cow!",
		"goat":   "Just opened her throat and swallowed a goat!",
		"dog":    "What a hog, to swallow a dog!",
		"cat":    "Imagine that, to swallow a cat!",
		"bird":   "How absurd to swallow a bird!",
		"spider": "It wriggled and jiggled and tickled inside her.",
		"horse":  "She's dead, of course!",
	}
	countm = map[int]string{
		8: "horse",
		7: "cow",
		6: "goat",
		5: "dog",
		4: "cat",
		3: "bird",
		2: "spider",
		1: "fly",
	}
)

// Verse show particular paragraph
func Verse(i int) string {
	ani, ok := countm[i]

	if !ok {
		return ""
	}

	result := fmt.Sprintf("I know an old lady who swallowed a %s.", ani)

	switch {
	case i == 8:
		return result + "\n" + yellm[ani]
	default:
		if s, ok := yellm[ani]; ok {
			result += ("\n" + s)
		}
	}

	for n := i; n > 0; n-- {
		if s, ok := swallowm[countm[n]]; ok {
			result += ("\n" + s)
		}
	}

	return result
}

// Verses show particular paragraphs
func Verses(ints ...int) string {
	s := ""
	for _, v := range ints {
		s += Verse(v) + "\n\n"
	}
	return s[:len(s)-2]
}

// Song show all paragraphs
func Song() string {
	return Verses(1, 2, 3, 4, 5, 6, 7, 8)
}
