package main

// chan<- <type> : Send-only channel.
// <-chan <type> : Receive-only channel.
// chan <type>   : Bidirectional channel.

import "fmt"

func main() {
	c := gen()

	receive(c)

	fmt.Println("Program exit...")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for n := range c {
		fmt.Println(n)
	}
}
