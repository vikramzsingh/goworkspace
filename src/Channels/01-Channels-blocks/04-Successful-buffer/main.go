package main

import "fmt"

func main() {
	c := make(chan int, 1)

	// put onto The Channel
	c <- 42
	c <- 43

	// Pull off the Channel
	fmt.Println(<-c)
}
