package main

// Sources in page 111

import "fmt"

/*
	- Address operator (&): The & operator is used to get the memory address of a variable. If
	`x` is an integer variable then `&x` will give you a pointer to `x` — that is, a memory address
	where the integer `x` is stored.
	- Dereferencing operator (*): The '*' operator is used to get the value stored at a memory
	address. If `p` is a pointer to an integer then `*p` gives you the integer that `p` points to
*/

func main() {
	// Seeing type and value for pointers
	x := 43
	pointerX := &x
	fmt.Println(pointerX)
	fmt.Printf("'pointerX' is of type: %T\n", pointerX)
	fmt.Printf("'pointerX' points to this memory address: %v\n", pointerX)
	fmt.Println()

	lastname := "Vasilopoulos"
	fmt.Println(lastname)
	fmt.Printf("type: %T\n", &lastname)
	fmt.Printf("points to this memory address: %v\n", &lastname)
	fmt.Println()
	fmt.Println("-----")

	// accessing the value of a pointer:
	valueOfPointerX := *pointerX
	fmt.Println(valueOfPointerX)
	fmt.Printf("%v is of type: %T\n", valueOfPointerX, valueOfPointerX)
	fmt.Printf("'pointerX' points to this memory address: %v\n", pointerX)
	fmt.Printf("'valueOfPointerX' points to this memory address: %v\n", &valueOfPointerX)
	fmt.Println()

	// Change the value of 'x' by modifying the *y:
	fmt.Println(x)
	*pointerX = 12
	fmt.Println("Value of 'x' after modifying a pointer that points in the same address as 'x': ", x)

}
