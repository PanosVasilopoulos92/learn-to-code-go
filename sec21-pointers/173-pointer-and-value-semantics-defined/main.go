package main

import "fmt"

// Sources in page 116

/*
	"Value semantics" in Go refers to when the actual data of a variable is passed to a function or
	assigned to another variable. This means that the new variable or function parameter gets a
	completely independent copy of the data - EVERYTHING IN GO IS 'PASS BY VALUE'.
*/

// value semantics(Pass by value)
func addOne(n int) int {
	return n + 1
}

// pointer semantics (Again passes by value, but the value is actually a pointer, so is like passing by reference.)
func addOneP(v *int) {
	*v += 1
}

func main() {
	// value semantics
	a := 2
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)
	fmt.Println()

	// pointer semantics
	b := 4
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
	fmt.Println()
}
