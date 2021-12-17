package main

import "fmt"

func main() {
	c := make(chan int, 2) // Buffer channel allows certains values to sit in, on that channel
	// that means my buffer channel allows to sit one Value
	// not creating deadlock

	// put onto the Channnel
	c <- 42

	fmt.Println(<-c) // Pull of the Channel
}
