package main

// Sources located at page 89

import "fmt"

// Function expression is assigning a function to a variable.

func main() {
	// Assign an anonymous function to a variable.
	x := func() {
		fmt.Println("I belong to an anonymous function.")
	}

	y := func(s string) {
		fmt.Printf("My lastname is %s.", s)
	}

	x()
	y("Vasilopoulos")
}
