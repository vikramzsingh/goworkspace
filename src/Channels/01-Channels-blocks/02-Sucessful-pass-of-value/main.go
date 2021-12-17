package main

import "fmt"

func main() {
	c := make(chan int)

	go func() { // the cuncurrent pattren design
		c <- 42 // Put onto The Channel
	}() // This code Blocked in this Go routine

	// This one block untill it takes the value Off
	fmt.Println(<-c) // Pull/take Off the Channel
}
