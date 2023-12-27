package main

// Sources located at page 88

import "fmt"

// A named funtion:
// func (r receiver) identifier(p parameter/s) (return/s) {code}

// An anonymous function:
// func(p parameter/s)(return/s) {code}()

func main() {
	foo()
	fmt.Println()

	func() {
		fmt.Println("I belong to an anonymous function.")
	}() // The '()' at the end will enable the anonymous faction to be executed.

	func(s string) {
		fmt.Printf("My lastname is %s.", s)
	}("Vasilopoulos") // will pass that as an argument to the anonymous faction.
}

func foo() {
	fmt.Println("I belong to a named function.")
}
