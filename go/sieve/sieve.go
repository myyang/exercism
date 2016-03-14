/*
Package sieve provides sieve algorithm to check prime

Optimized ref: http://play.golang.org/p/9U22NfrXeq

*/
package sieve

// Generate all values
func Generate(out chan<- int, limit int) {
	for i := 2; i < limit+1; i++ {
		out <- i
	}
	close(out)
}

// Filter primes
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
	close(out)
}

// Sieve returns array
func Sieve(limit int) []int {
	values := make(chan int)
	go Generate(values, limit)

	result := []int{}

	for {
		tmp, ok := <-values
		if !ok {
			break
		}
		result = append(result, tmp)
		primes := make(chan int)
		go Filter(values, primes, tmp)
		values = primes
	}

	return result
}
