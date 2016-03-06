// Package hamming provides ...
package hamming

const testVersion = 3

// HammError defines custom error
type HammError struct{}

// Error return naive error msg
func (e HammError) Error() string {
	return "Hamming Error"
}

// Distance shows hamming distance count between two string
func Distance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return -1, HammError{}
	}

	val := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			val++
		}
	}

	return val, nil
}
