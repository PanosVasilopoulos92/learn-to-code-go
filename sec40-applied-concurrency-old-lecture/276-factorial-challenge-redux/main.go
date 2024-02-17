package main

import (
	"fmt"
	"log"
)

func main() {
	f := factorial(4)
	for n := range f {
		fmt.Println(n)
	}

	// Calculate the factorial using concurrency techniques
	in := gen()
	f2 := factorialV2(in)
	for n := range f2 {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		result := 1
		for i := 0; i > 0; i-- {
			result *= i
		}
		out <- result
		close(out)
	}()
	return out
}

// Calculate the factorial using concurrency techniques

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorialV2(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	if n <= 0 {
		log.Fatalf("Cannot accept negative or zero number values.")
	}

	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}
