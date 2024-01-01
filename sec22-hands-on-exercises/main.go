package main

import "fmt"

func main() {
	x := 44
	y := &x
	fmt.Println(&x) // wil print the memory address in which the value of 'x' exist.
	fmt.Println(*y) // wil print the value that exist in the memory address that keeps the value of 'x'.
}
