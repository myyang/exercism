// Package letter provides functions to count word frequency
package letter

import (
	"sync"
)

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

func channelFreq(s string, c chan FreqMap, gl *sync.WaitGroup) {
	defer gl.Done()
	c <- Frequency(s)
}

func channelManage(c chan FreqMap, gl *sync.WaitGroup) {
	gl.Wait()
	close(c)
}

// ConcurrentFrequency counts frequency concurrently using channel
func ConcurrentFrequency(strArr []string) FreqMap {
	gl := &sync.WaitGroup{}
	c, m := make(chan FreqMap), FreqMap{}
	for _, v := range strArr {
		gl.Add(1)
		go channelFreq(v, c, gl)
	}

	go channelManage(c, gl)

	for v := range c {
		for k, x := range v {
			if _, ok := m[k]; ok {
				m[k] += x
			} else {
				m[k] = x
			}
		}
	}
	return m
}
