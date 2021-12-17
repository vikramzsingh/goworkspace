package main

import (
	"fmt"
	"runtime"
	"sync"

)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CUPs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // at this point we already have main go routine

	// control flow send this out to execute
	// and continue to run rest of the program ( beyond bar() )
	// Signature of Add function:
	// 					func (wg *WaitGroup) Add(delta int)

	wg.Add(1) // counter // added count to wait for one thing to done
	go foo() // now we have launched another go routine of foo function
	bar()

	fmt.Println("CUPs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // at this point there are total 2 go routines
	wg.Wait() // after done it will run por woit untill the counter is 0
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // remove that one thing done
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}