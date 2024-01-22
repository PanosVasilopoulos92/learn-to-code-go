package main

// Resources at page 132

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	counter := 0

	const gs int = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(1000)
			runtime.Gosched() // With this line a race condition will be born.
			v++
			counter = v
			wg.Done()
		}()
		// fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println(counter)
}
