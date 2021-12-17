package main

import "fmt"

func main() {
	c := make(<-chan int) // recieve-only from channel

	go func() {
		c <- 42
	}()

	fmt.Printf(<-c)
}
