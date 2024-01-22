package main

// Resources at page 136

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)

	go second()

	go one()

	go func() {
		fmt.Println("Goroutine one, from the main function.")
		wg.Done()
	}()

	go func() {
		fmt.Println("Goroutine two, from the main function.")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Exit program...")
}

func one() {
	fmt.Println("Goroutine number one.")
	wg.Done()
}

func second() {
	fmt.Println("Goroutine number two.")
	wg.Done()
}
