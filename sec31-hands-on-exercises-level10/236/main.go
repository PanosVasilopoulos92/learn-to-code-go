package main

// If we want a struct to implement an interface, it must implement EVERY METHOD this interface has.

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Error info: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "error in customerID",
	}
	foo(c1)
}

func foo(err error) {
	fmt.Println("foo ran -", err)
	fmt.Println("foo ran -", err.(customErr).info) // Assertion
}
