package main

// Sources located at page 94

import "fmt"

func main() {
	// Factorial of 4 is: 4 * 3 * 2 * 1

	result1 := factorial(8)
	fmt.Println(result1)
	fmt.Println("------")

	result2 := factoLoop(8)
	fmt.Println(result2)
}

func factorial(n int) int {
	// When it reach Base Case
	if n == 0 {
		return 1
	}

	result := n * factorial(n-1)
	return result
}

func factoLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
