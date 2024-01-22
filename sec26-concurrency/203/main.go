package main

import "fmt"

func doSomething(x int) int {
	return x * x
}

func foo() {
	fmt.Println("Printed first.")
}

func main() {
	ch := make(chan int)
	// In order for the goroutine to return a value.
	go func() {
		ch <- doSomething(4)
	}()

	foo()
	fmt.Println(<-ch)
}
