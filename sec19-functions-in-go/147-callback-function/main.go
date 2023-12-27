package main

// Sources located at page 92

import "fmt"

func main() {
	x := doMath(10, 12, add)
	fmt.Println(x)

	y := doMath(100, 50, substract)
	fmt.Println(y)
	fmt.Println()

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", substract)
	fmt.Printf("%T\n", doMath)
}

func add(a int, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

// Callback function (A function that takes as parameter another function).
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
