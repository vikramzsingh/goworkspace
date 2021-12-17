package main

import (
	"fmt"
	"runtime"
	"sync"

)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("go routines", runtime.NumGoroutine())

	counter := 0

	const gs = 100 // constant declared, we going to launch 100 go routines
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter      // every go func accessing different func, not accessing same function
			// time.Sleep(time.Second)      // sleep for spcified duration
			runtime.Gosched() // allow go to run something else if you had time.
			v++
			counter = v
			wg.Done()
		}() // self executing function, anonymous function
		fmt.Println("go routines", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("go routines", runtime.NumGoroutine())
	fmt.Println("Count", counter)
}
