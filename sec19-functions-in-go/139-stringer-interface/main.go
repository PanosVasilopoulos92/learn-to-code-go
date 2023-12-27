package main

// Sources located at page 76

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

// !! The following will attach this function as a method to the type inside the receiver parameter. !!
// So now the 'book' type has the method 'String() string', which belongs to the interface 'Stringer'.
func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type price float64

// !! The following will attach this function as a method to the type inside the receiver parameter. !!
// The return value will be deferent than the previous, which means that we have polymorphism in our code.
func (p price) String() string {
	return fmt.Sprint("It's price is ", strconv.FormatFloat(float64(p), 'f', 2, 64)) // Converts a float64 value to a string
}

func main() {
	b := book{
		title: "Once upon a time",
	}
	var priceofBook price = 42.60

	fmt.Println(b)
	fmt.Println(priceofBook)
}
