package main

import (
	"fmt"
	"math"
)

// Exercise 61:
type Person struct {
	firstname string
	age       int
}

// Attaches the function speak() to type person.
func (p Person) speak() {
	fmt.Printf("My name is %s and I am %d years old.\n", p.firstname, p.age)
}

// Exercise 62:
type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	calcArea() float64 // Any type that has this method signature, is also of type 'Shape'.
}

func info(s Shape) float64 {
	return s.calcArea()
}

func (s Square) calcArea() float64 {
	return s.length * s.width
}

func (c Circle) calcArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func displayCircleArea(c Circle) {
	fmt.Printf("%.4f\n", c.calcArea())
}

func displaySquareArea(s Square) {
	fmt.Printf("%.3f\n", s.calcArea())
}

func main() {
	// Exercise 57:
	x := addNumbers([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
	fmt.Println("------")

	// Exercise 58:
	y := foo()
	fmt.Println(y)
	z1, z2 := bar()
	fmt.Println(z1, z2)
	fmt.Println("------")

	// Exercise 59:
	result := takesVariadicParam(1, 2, 3, 4, 5, 60)
	fmt.Println(result)
	//Or
	xi := []int{1, 2, 3, 4, 5, 60}
	fmt.Println(takesVariadicParam(xi...)) // unfurls the slice.
	fmt.Println("------")

	// Exercise 60:
	// Deferred functions run in LIFO order.
	/*
		defer testDefer3() //Printed last
		defer testDefer1() //Printed third
		testDefer2()       // Printed first
		defer testDefer4() //Printed second
		fmt.Println("------")
	*/

	// Exercise 61:
	p1 := Person{
		firstname: "Panos",
		age:       31,
	}

	p2 := Person{
		firstname: "Efraim",
		age:       23,
	}

	p1.speak()
	p2.speak()
	fmt.Println("-------")

	// Exercise 62:
	c1 := Circle{
		radius: 12.24,
	}

	square1 := Square{
		length: 8.20,
		width:  8.20,
	}

	// fmt.Println(c1.calcArea())
	// fmt.Println(square1.calcArea())
	displayCircleArea(c1)
	displaySquareArea(square1)
	fmt.Println()
	//Or using 'info' function (polymorphism):
	fmt.Println(info(c1))
	fmt.Println(info(square1))

}

// Exercise 57: Named returns. Usually we avoid them.
func addNumbers(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return total
}

// Exercise 58:
func foo() int {
	return 12
}

func bar() (int, string) {
	return 10, "Ok"
}

// Exercise 59:
func takesVariadicParam(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

// Exercise 60:
/*
func testDefer1() {
	fmt.Println("testDefer1.")
}

func testDefer2() {
	fmt.Println("testDefer2.")
}

func testDefer3() {
	fmt.Println("testDefer3.")
}

func testDefer4() {
	fmt.Println("testDefer4.")
}
*/
