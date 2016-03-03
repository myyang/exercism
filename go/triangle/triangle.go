//+build !example

package triangle

import (
	"math"
	"sort"
)

const testVersion = 2

// Kind defines interface
type Kind interface{}

type nat struct {
	Kind
}

type equ struct {
	Kind
}

type iso struct {
	Kind
}

type sca struct {
	Kind
}

// Types definition
var (
	NaT nat // NaT not a tri
	Equ equ // Equ equalence
	Iso iso // Iso isolence
	Sca sca // Sca scalence
)

// KindFromSides determins triangle types
func KindFromSides(a, b, c float64) Kind {
	v := []float64{a, b, c}
	sort.Float64s(v)
	a, b, c = v[0], v[1], v[2]

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 || (a+b) < c {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c {
		return Iso
	}

	return Sca

}
