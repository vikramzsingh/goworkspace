package main

import "fmt"

func main() {
	c := make(chan<- int) // send-only onto channel

	go func() {
		c <- 42
	}()

	fmt.Printf(<-c)
}
