package main

import "fmt"

func main() {
	// send-only channel
	c := make(chan<- int, 2) // Only send values onto the channel

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------")
	fmt.Printf("%T\n", c)
}
