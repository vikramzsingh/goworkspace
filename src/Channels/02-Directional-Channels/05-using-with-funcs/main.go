package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int) // bi-directional channel

	// send
	go foo(c)

	// recieve
	wg.Add(1)
	go bar(c)

	wg.Wait()

	// time.Sleep(time.Second*2)
	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) { // channel of int
	c <- 42
}

// recieve
func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
