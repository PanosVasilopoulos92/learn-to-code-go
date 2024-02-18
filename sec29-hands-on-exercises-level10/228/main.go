package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup // WaitGroup to track goroutines

func main() {
	numbers := make(chan int)
	orderedNumbers := []int{}

	// Launch 10 goroutines
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// Every goroutine will send the next 10 numbers to the channel.
		go func(startPoint int) {
			defer wg.Done()

			for j := startPoint; j <= startPoint+10; j++ {
				numbers <- j
			}
		}(i * 10) // Passes the argument to anonymous function.

	}

	// Close the channel (after goroutines finish)
	go func() {
		wg.Wait() // Wait for all goroutines to complete
		close(numbers)
	}()

	for n := range numbers {
		orderedNumbers = append(orderedNumbers, n)
	}

	sort.Ints(orderedNumbers)

	for i := 0; i < len(orderedNumbers); i++ {
		fmt.Println(orderedNumbers[i])
	}

}
