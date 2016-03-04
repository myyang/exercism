// Package gigasecond provides a function to add giga seconds to given time
package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds 1000,000,000 seconds to given time
func AddGigasecond(originTime time.Time) time.Time {
	return originTime.Add(time.Second * 1e9)
}
