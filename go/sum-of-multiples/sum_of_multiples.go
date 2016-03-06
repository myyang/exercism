// Package summultiples provides ...
package summultiples

type sumFn func(int) int

// MultipleSummer supports multiple multiples declartion
func MultipleSummer(muls ...int) sumFn {
	return func(limit int) int {
		sum := 0
		for i := 1; i < limit; i++ {
			for _, v := range muls {
				if i%v == 0 {
					sum += i
					break
				}
			}
		}
		return sum
	}
}
