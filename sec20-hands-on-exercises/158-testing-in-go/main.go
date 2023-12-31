package main

import "fmt"

func main() {
	additionResult := doMath(10, 4, Add)
	fmt.Println(additionResult)
	fmt.Println()

	substractionResult := doMath(10, 4, Substract)
	fmt.Println(substractionResult)
	fmt.Println()

	name := displayMyName("Panos")
	fmt.Println(name)
}

func Add(x int, y int) int {
	return x + y
}

func Substract(x int, y int) int {
	return x - y
}

func doMath(x int, y int, f func(x int, y int) int) int {
	return f(x, y)
}

func displayMyName(s string) string {
	return fmt.Sprint(s)
}
