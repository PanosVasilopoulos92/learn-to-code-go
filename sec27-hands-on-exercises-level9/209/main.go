package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	incrementor := 0
	gs := 38
	wg.Add(gs)

	/*
		go func() {
			runtime.Gosched()
			fmt.Println(incrementor, "first func start.")
			tempIncrementor1 := incrementor
			tempIncrementor1++
			incrementor = tempIncrementor1
			fmt.Println(incrementor, "first func end.")
			fmt.Println()
			wg.Done()
		}()

		for i := 0; i <= 120; i++ {
			go func() {
				fmt.Println(incrementor, "second func start.")
				fmt.Println(runtime.NumGoroutine())
				fmt.Println(runtime.NumCPU(), "CPUs")
				tempIncrementor2 := incrementor
				runtime.Gosched()
				tempIncrementor2++
				incrementor = tempIncrementor2
				fmt.Println(incrementor, "second func end.")
				fmt.Println()
				wg.Done()
			}()
		}

		go func() {
			runtime.Gosched()
			fmt.Println(incrementor, "third func start.")
			tempIncrementor3 := incrementor
			tempIncrementor3++
			incrementor = tempIncrementor3
			fmt.Println(incrementor, "third func end.")
			fmt.Println()
			wg.Done()
		}()

	*/

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementor
			runtime.Gosched() // will create a race condition
			v++
			incrementor = v
			fmt.Println(incrementor)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementor)
}
