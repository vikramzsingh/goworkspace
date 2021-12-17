package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// recieve	
	recieve(eve, odd, quit) 

	fmt.Println("about to exit")
}

func recieve(e, o, q <-chan int) {
	for {
		select {
		case v := <- e:
			fmt.Println("From even Channel:", v)
		case v := <- o:
			fmt.Println("From odd Channel:", v)
		case v := <- q:
			fmt.Println("From quit Channel:", v)
			return	
		} 
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}