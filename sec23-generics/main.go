package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Resources in page 125

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// Generic function
func addT[T int | float64](a, b T) T {
	return a + b
}

// type set interface
type myNumber interface {
	int | float64
}

func addTVersion2[T myNumber](a, b T) T {
	return a + b
}

// The '~' is used in order to be able to pass as argument a type of int or a type that underlying the type int.
type myNumberV2 interface {
	~int | ~float64
}

func addTVersion3[T myNumberV2](a, b T) T {
	return a + b
}

// Create an alias for a type.
type myInt int

// Package constraints
type myNumberV3 interface {
	constraints.Integer | constraints.Float
}

func addTVersion4[T myNumberV3](a, b T) T {
	return a + b
}

func main() {
	var age myInt = 31
	fmt.Println(age)
	fmt.Println(addTVersion3(age, 1))
	fmt.Println()

	fmt.Println(addTVersion4(age, 2))
	fmt.Println()

	fmt.Println(addI(2, 4))
	fmt.Println(addF(2, 4.2))
	fmt.Println()

	fmt.Println(addT(2, 4))
	fmt.Println(addT(2, 4.2))
	fmt.Println()

	fmt.Println(addTVersion2(2, 4))
	fmt.Println(addTVersion2(2, 4.2))
	fmt.Println()
}
