/*
Package pascal provides a function returns pascal triangle array

Pascal triangle may look like this:
        1
       1 1
      1 2 1
     1 3 3 1
    1 4 6 4 1

with represention by array:
	[][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

See also: https://en.wikipedia.org/wiki/Pascal%27s_triangle
*/
package pascal

// Triangle returns a pascal triangle
// represneted with int array of array
func Triangle(n int) [][]int {
	r := [][]int{{1}, {1, 1}}
	switch {
	case n < 1:
		return [][]int{}
	case n == 1:
		return r[:1]
	case n == 2:
		return r
	}

	for i := 3; i < n+1; i++ {
		t := r[len(r)-1]
		s := make([]int, i)
		s[0], s[i-1] = 1, 1
		for j := 1; j < i/2+1; j++ {
			c := t[j] + t[j-1]
			s[j], s[i-1-j] = c, c
		}
		r = append(r, s)
	}
	return r
}
