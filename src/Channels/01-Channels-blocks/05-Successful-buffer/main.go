package main

import "fmt"

func main() {
	c := make(chan int, 2) // buffer Channel allows to values sit in

	// put Onto The Channel
	c <- 42 // send onto Channel
	c <- 43

	fmt.Println(<-c) // Pull/Take off the Channel
	fmt.Println(<-c) // Recieve from Channel
}
