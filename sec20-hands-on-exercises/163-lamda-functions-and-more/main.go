package main

import "fmt"

// Sources located at page 105

/*
	An anonymous func, aka a function literal or a lambda function, is a way to define a function
	without giving it a name. Instead, you can directly declare and define a function inline wherever
	it is needed. Anonymous funcs are commonly used when you want to pass a function as an
	argument to another function or when you need to define a short-lived function for a specific
	task.
*/

// An anonymous/lambda function:
// func(p parameter/s)(return/s) {code}()

func main() {

	x := func(name string) string {
		return fmt.Sprint(name)
	}("Panos")

	fmt.Println(x)
	fmt.Println("-----")

	func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d\t", i)
		}
	}()
	fmt.Println()
	fmt.Println("-----")

	// Assigning a function to a variable:
	y()
	fmt.Println()
	fmt.Println("-----")

	// Or like that:
	z := func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d\t", i)
		}
	}

	z()
	fmt.Println()
	fmt.Println("-----")

	// Returned function
	f := outer()
	fmt.Println("The value of the returned function is:", f())
}

// Assigning a function to a variable:
var y = func() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", i)
	}
}

// A function that returns a func
func outer() func() int {
	return func() int {
		return 12 + 12
	}
}
