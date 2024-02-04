package main

// Resources at page 135

import "fmt"

func main() {
	c := make(chan int)

	// Buffered channel
	c2 := make(chan int, 2) // Second argument uses as a buffer for the number of values that this channel can contain.

	// In order for this to work the corresponding channel must have a buffer.
	c2 <- 33
	c2 <- 12

	go func() {
		c <- 42
	}()

	fmt.Println(<-c2)
	fmt.Println(<-c2) // For the second value that this channel contains.
	fmt.Println(<-c)
}
