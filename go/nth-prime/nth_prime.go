// Package prime provides function to get n-th prime
package prime

func generate(out chan int) {
	for i := 2; ; i++ {
		out <- i
	}
}

func filter(in chan int, out chan int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}

// Optimized ref: http://play.golang.org/p/9U22NfrXeq
func sieve(limit int) int {
	values, r := make(chan int), 0
	go generate(values)

	for i := 0; i < limit; i++ {
		prime := <-values
		r = prime
		primes := make(chan int)
		go filter(values, primes, prime)
		values = primes
	}
	return r

}

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	return sieve(n), true
}
