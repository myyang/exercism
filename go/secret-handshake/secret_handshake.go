// Package secret provides handshake
package secret

var smap = map[int]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

// Handshake performs encoding
func Handshake(code int) []string {
	if code < 1 {
		return nil
	}

	b := toBinStr(code)
	if code > 15 {
		b = b[:len(b)-1]
	}
	var result []string
	mul := 1
	for _, v := range b {
		if v != 0 {
			result = append(result, smap[v*mul])
		}
		mul = mul * 2
	}

	if code > 15 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}
	return result
}

func toBinStr(n int) []int {
	if n < 2 {
		return []int{n}
	}
	return append([]int{n % 2}, toBinStr(n/2)...)
}
