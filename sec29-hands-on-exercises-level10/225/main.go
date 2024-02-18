package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("Program exit...")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			fmt.Println("exiting goroutine")
			return
		case <-time.After(1 * time.Millisecond):
			fmt.Println("Waiting more than 1ms.")
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 30; i++ {
			time.Sleep(10 * time.Millisecond)
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
