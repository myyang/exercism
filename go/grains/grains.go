// Package grains provides functions that relates to a story
// See more on wiki: https://en.wikipedia.org/wiki/Wheat_and_chessboard_problem
package grains

// NError is custom invalid n error
type NError struct{}

// Error return error msg
func (e NError) Error() string {
	return "Invalid N"
}

// Square returns the grains number in the square by given n
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, NError{}
	}
	uresult, ushift := uint64(1), uint(n-1)
	return uresult << ushift, nil
}

// Total returns the total gratins number on the board
func Total() uint64 {
	uresult := uint64(1)
	return uresult<<65 - 1
}
