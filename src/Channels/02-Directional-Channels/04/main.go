package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // recieve from channel
	cs := make(chan<- int) // send onto channel

	fmt.Println("-------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
