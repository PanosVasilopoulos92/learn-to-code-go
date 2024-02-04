package mymath

import (
	"slices"
)

// CenteredAvg computes the average of a list of numbers after it removes the smallest and the biggest values.
func CenteredAvg(xi []int) float64 {
	slices.Sort(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
