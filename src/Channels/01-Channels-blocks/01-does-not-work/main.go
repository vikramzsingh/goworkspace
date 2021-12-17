package main

import "fmt"

func main() {
	c := make(chan int) // Channel onto which i put integers

	// Put something onto that
	c <- 42 // on to c put on number 42
	// channels is blocked in go routine, because no one pulling off this channel, so it is blocked

	// off of c
	fmt.Println(<-c) // take off of it 42 and put it there
}

// Note:- This code does not run
