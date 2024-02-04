// Package acdc contains the function 'Sum'.
package acdc

// Sum adds an unlimited number of integers
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
