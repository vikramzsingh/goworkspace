package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int) // bi-directional channel

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- 1
		}
		close(c)
	}()

	// recieve
	for v := range c {
		fmt.Println(v)
	}

	// time.Sleep(time.Second*2)
	fmt.Println("about to exit")
}
