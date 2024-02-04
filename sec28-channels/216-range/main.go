package main

import "fmt"

func main() {
	c := make(chan int)
	c2 := make(chan int)

	// send
	go func() {
		counter := 0
		for i := 0; i <= 20; {
			c <- i
			i++
			counter++
		}
		close(c)

		c2 <- counter
		close(c2)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Total numbers displayed:", <-c2)
	fmt.Println()

	fmt.Println("Program exits...")

}
