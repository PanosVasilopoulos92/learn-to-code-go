package main

import (
	"fmt"
	"math/rand"
	"time"
)

// FAN IN pattern is when multiple channels are writing to the same channel.

func main() {
	c := fanIn(boring("Joe"), boring("Mary"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // Here the main will block as the channel is waiting for a value.
	}
	fmt.Println("You're both boring. I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // It takes the value from value 'input1' channel and pass it to the channel 'c'.
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
