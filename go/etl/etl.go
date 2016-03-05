// Package etl provides ETL trainsform function
package etl

import "strings"

// Transform origin map to reverse lookup map
func Transform(given map[int][]string) map[string]int {

	var result = map[string]int{}
	for k, v := range given {
		for _, s := range strings.ToLower(strings.Join(v, "")) {
			result[string(s)] = k
		}
	}
	return result
}
