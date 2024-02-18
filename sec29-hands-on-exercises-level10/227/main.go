package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 30; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for n := range c1 {
			c2 <- n
		}
		close(c2)
	}()

	for n := range c2 {
		fmt.Println(n)
	}
}
