package main

import (
	"fmt"
	"math"
	"time"
)

// Resources at page 106

/*
Callbacks are essentially functions that are passed as arguments to other functions and are
intended to be called (or "executed") later in your program.
*/

func main() {
	//Callbacks
	result := printSquare(square, 8)
	fmt.Println(result)
	fmt.Println()
	fmt.Println("-------")

	// Closures
	x := powinator(2)
	fmt.Println(x()) // 2 ** 1
	fmt.Println(x()) // 2 ** 2
	fmt.Println(x()) // 2 ** 3
	fmt.Println(x()) // 2 ** 4
	fmt.Println()
	fmt.Println("-------")

	// !! A wrapper function wraps or modifies another function's behavior, !!
	// while a callback function is a function passed as an argument to be executed at a specific point or event.
	timeFunc(iterateBetweenNumbers)
}

// Callbacks
func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("The number %v squared is %v", a, x)
}

func square(n int) int {
	return n * n
}

// Closures
func powinator(a float64) func() float64 {
	c := 0.0
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

// Wrapper function
func timeFunc(f func()) {
	start := time.Now()
	// Wrapped function
	f()
	end := time.Since(start)
	fmt.Printf("Function 'iterateBetweenNumbers' executed in %v", end)
}

// Wrapper function
func iterateBetweenNumbers() {
	i := 0
	for i < 10 {
		// Wrapped function
		fmt.Println(i)
		i++
	}

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// }
}
