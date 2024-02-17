package main

import "fmt"

func main() {
	c := make(chan int)
	// c <- 1 // This would create a deadlock.
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
