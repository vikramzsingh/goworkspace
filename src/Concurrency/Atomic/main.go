package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go Routines", runtime.NumGoroutine())

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) // write to counter 
			runtime.Gosched() // if you have time do other things
			fmt.Println("counter\t", atomic.LoadInt64(&counter)) // read from counter
			wg.Done()
		}()
		fmt.Println("Go Routines", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Go Routines", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
