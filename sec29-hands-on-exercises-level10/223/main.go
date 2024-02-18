package main

// chan<- <type> : Send only channel.
// <-chan <type> : Receive only channel.
// chan <type>   : Bidirectional channel.

import "fmt"

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42 // Passes a new value to the channel
	}()
	fmt.Println(<-cs) // Receives a value from this channel.

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
