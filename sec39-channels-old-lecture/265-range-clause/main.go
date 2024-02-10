package main

import "fmt"

func main() {
	c := make(chan int, 10) // buffered channel. Will load 10 values, channel closes and after the values will be printed one by one in the range loop.

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
