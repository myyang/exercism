// Package beer provides sss
package beer

import (
	"errors"
	"fmt"
	"strings"
)

var (
	checkBeer = "%v of beer on the wall, %v of beer.\n"
	takeBeer  = "Take one down and pass it around, %v of beer on the wall.\n"
	takeLast  = "Take it down and pass it around, no more bottles of beer on the wall.\n"
	buyBeer   = "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
)

// Verse n paragraph
func Verse(n int) (string, error) {
	if n > 99 {
		return "", errors.New("Out of range")
	}
	s := ""
	if n > 1 {
		beerN, beerN1 := fmt.Sprintf("%v bottles", n), fmt.Sprintf("%v bottles", n-1)
		if n-1 == 1 {
			beerN1 = "1 bottle"
		}
		s = fmt.Sprintf(checkBeer, beerN, beerN) + fmt.Sprintf(takeBeer, beerN1)
	} else if n == 1 {
		beerN := "1 bottle"
		s = fmt.Sprintf(checkBeer, beerN, beerN) + takeLast
	} else {
		beerN := "No more bottles"
		s = fmt.Sprintf(checkBeer, beerN, strings.ToLower(beerN)) + buyBeer
	}
	return s, nil
}

// Verses upper to lower paragraphs
func Verses(upper, lower int) (string, error) {
	if upper > 100 || lower < 0 || lower > upper {
		return "", errors.New("Out of range")
	}
	s := ""
	for i := upper; i >= lower; i-- {
		tmp, err := Verse(i)
		if err != nil {
			return "", errors.New("Error verse")
		}
		s += (tmp + "\n")
	}
	return s, nil
}

// Song whole song
func Song() string {
	song, err := Verses(99, 0)
	if err != nil {
		return ""
	}
	return song
}
