package main

// Sources located at page 94

import "fmt"

func main() {
	f := incrementor()
	// fmt.Println(f) // It will point not the value but to the memory address that the value exists.
	fmt.Println(f())
	fmt.Printf("%T\n", f())
	fmt.Printf("%T\n", f)
	fmt.Println("-------")

	f = incrementor() // When assigning again, the 'x' variable will be set to back to 0.
	fmt.Println(f())  // 1
	fmt.Println(f())  // 2
	fmt.Println(f())  // 3
	fmt.Println("-------")

	f = incrementor() // When assigning again, the 'x' variable will be set to back to 0.
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println("-------")

	g := incrementor()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
