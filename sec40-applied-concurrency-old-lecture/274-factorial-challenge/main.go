package main

import "fmt"

func main() {
	f := factorial(4)
	c := factorialV2(20)
	fmt.Println("Result of factorial =", f)
	fmt.Println("Result of factorialV2=", <-c)
}

func factorial(n int) int {
	if n <= 0 {
		return -1
	}
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialV2(n int) chan int {
	out := make(chan int)
	go func() {
		result := 1
		for i := n; i > 0; i-- {
			result *= i
		}
		out <- result
		close(out)
	}()
	return out
}
