// Package gigasecond providees giga function
package gigasecond

import "time"

const testVersion = 4 // find the value in gigasecond_test.go

// AddGigasecond adds 1000,000,000
func AddGigasecond(in time.Time) time.Time {
	return in.Add(time.Second * 1e9)
}
