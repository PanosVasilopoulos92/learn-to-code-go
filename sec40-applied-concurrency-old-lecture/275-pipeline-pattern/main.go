package main

import "fmt"

func main() {
	// Set up the pipeline
	c := gen(2, 3, 4, 5)
	out := sq(c)

	// Consume the output. For every value I pass in the gen() I need to create a "<-out" in order to be displayed.
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
