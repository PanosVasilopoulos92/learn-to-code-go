package main

// Now we gonna fix the race condition that was created in '209' folder, with the use of Mutex.

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var m sync.Mutex

func main() {
	incrementor := 0
	gs := 42
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementor
			// runtime.Gosched() // will create a race condition
			v++
			incrementor = v
			fmt.Println(incrementor)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	fmt.Println("end value:", incrementor)
}
