package main

import (
	"fmt"

	"github.com/PanosVasilopoulos92/learn-to-code-go/sec35-hands-on-exercises-level13/255/mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 10}
	b := []int{0, 8, 10, 100}
	c := []int{1000, 4, 6, 8, 10}
	d := []int{111, 4, 6, 8, 999}
	e := [][]int{a, b, c, d}
	return e
}
