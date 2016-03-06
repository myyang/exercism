// Package letter provides functions to count word frequency
package letter

// FreqMap is frequency mapping
type FreqMap map[rune]int

// Frequency counts the character frequency
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts frequency concurrently using channel
func ConcurrentFrequency(strArr []string) FreqMap {
	c, m := make(chan FreqMap), FreqMap{}
	for _, v := range strArr {
		go func(s string) { c <- Frequency(s) }(v)
	}

	for range strArr {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
