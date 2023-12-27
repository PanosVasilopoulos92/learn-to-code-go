package main

// Sources located at page 91

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", foo())
	fmt.Printf("%T\n", foo)
	fmt.Println()

	// Now that I return an anonymous function as a value, the assigned variable will actually be a function.
	y := bar()
	fmt.Println(y())
	fmt.Printf("%T \n", y())
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", bar())
	fmt.Printf("%T\n", bar)
}

func foo() int {
	return 32
}

func bar() func() int {
	return func() int {
		return 42
	}
}
