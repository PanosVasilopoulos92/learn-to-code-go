package main

// Now we gonna fix the race condition that was created in '209' folder, with the use of Mutex.

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementor int64 // !! Package Atomic does not guaranty that the order displayed, just that every value will be accessed once!!
	gs := 42
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			/*
				m.Lock()
				v := incrementor
				// runtime.Gosched() // will create a race condition
				v++
				incrementor = v
				fmt.Println(incrementor)
				m.Unlock()
				wg.Done()
			*/
			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	fmt.Println("end value:", incrementor)
}
