package main

// Resources at page 138

import "fmt"

func main() {
	c := make(chan int, 1)

	// A channel that can only send values and nothing can pull out of it.
	cs := make(chan<- int, 1)

	// A channel that can only receive values, not send.
	cr := make(<-chan int, 1)

	c <- 44
	cs <- 10 // I can send value to this channel.
	// c3 <- 10 // Won't work because I can only retrieve from this channel.

	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Println("------")

	// fmt.Println(<-c2) // I cannot pull a value from this channel.
	fmt.Printf("%T\n", cs)
	fmt.Println("------")

	fmt.Printf("%T\n", cr)
}
