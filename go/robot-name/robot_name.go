// Package robotname provides Robot type
package robotname

import (
	"math/rand"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0987654321"
)

// Robot a string type
type Robot string

// Name a robot
func (r *Robot) Name() string {
	if string(*r) != "" {
		return string(*r)
	}
	r.Reset()
	return string(*r)
}

// Reset robot
func (r *Robot) Reset() {
	*r = Robot(getRnd(letters, len(letters), 2) + getRnd(numbers, len(numbers), 3))
}

// ref: http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func getRnd(src string, mask int, n int) string {
	r, l := make([]byte, n), len(src)
	for i, cache, shiftv := n-1, rand.Int63(), mask; i >= 0; {
		if shiftv == 0 {
			cache, shiftv = rand.Int63(), mask
		}

		if ind := int(cache & int64(shiftv)); ind < l {
			r[i] = src[ind]
			i--
		}

		cache >>= uint(shiftv)
		shiftv--
	}
	return string(r)
}
