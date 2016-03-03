// Package slice provides serialize
package slice

// All returns all slice
func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	arr := []string{}
	for i := 0; i < (len(s) - n + 1); i++ {
		arr = append(arr, s[i:i+n])
	}

	return arr
}

// Frist returns first string of array
func Frist(n int, s string) string {
	arr := All(n, s)
	if arr == nil {
		return ""
	}
	return arr[0]
}
