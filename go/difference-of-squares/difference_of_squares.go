// Package diffsquares provides diff sqaures
package diffsquares

// SquareOfSums performs
func SquareOfSums(n int) int {
	s := 0
	for i := 1; i < n+1; i++ {
		s += i
	}
	return s * s
}

// SumOfSquares performs
func SumOfSquares(n int) int {
	s := 0
	for i := 1; i < n+1; i++ {
		s += i * i
	}
	return s
}

// Difference performs
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
