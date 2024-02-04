// Package word provides custom function for working with strings and word.
package word

import (
	"strings"
)

// UseCount returns the number of times each word is used in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the numbers of words that exist in a string.
func Count(s string) int {
	// xs := strings.Fields(s)
	// totalWordsCounter := 0
	// for _, v := range xs {
	// 	fmt.Println(v)
	// 	totalWordsCounter++
	// }
	// return totalWordsCounter

	// Or much simpler and faster:
	totalWordsCounter := strings.Fields(s)
	return len(totalWordsCounter)

}
