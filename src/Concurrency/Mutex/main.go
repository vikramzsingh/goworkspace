package main

import (
	"fmt"
	"runtime"
	"sync"

)

func main() {
	counter := 0

	const gs = 100 // for launching 100 go routines
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex // to solve race condition

	for i := 0; i < gs; i++ {
		go func() {
			// locks for reading and writing
			// when some body using it, then someone can't able to uses it
			mu.Lock() // to pervent race condition
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // unlock when you are done
			wg.Done() // remove count for each add every time
		}()
		fmt.Println("Go routine", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Go routine", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
