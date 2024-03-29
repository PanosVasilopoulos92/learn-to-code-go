package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("Processes started...")
}

func main() {
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	fmt.Println("foo function ended.")
	fmt.Println()
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	fmt.Println("bar function ended.")
	fmt.Println()
	wg.Done()
}
