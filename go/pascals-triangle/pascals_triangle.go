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
	result := [][]int{{1}, {1, 1}}
	switch {
	case n < 1:
		return [][]int{}
	case n == 1:
		return result[:1]
	case n == 2:
		return result
	}

	for i := 3; i < n+1; i++ {
		lastRow := result[len(result)-1]
		// new row
		newRow := make([]int, i)
		// fill head and tail
		newRow[0], newRow[i-1] = 1, 1
		for j := 1; j < i/2+1; j++ {
			// count row value
			value := lastRow[j] + lastRow[j-1]
			newRow[j], newRow[i-1-j] = value, value
		}
		result = append(result, newRow)
	}
	return result
}
