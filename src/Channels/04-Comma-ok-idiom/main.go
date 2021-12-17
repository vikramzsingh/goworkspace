package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	recieve(even, odd, quit)

	fmt.Println("about to Exit")
}

// send Channel
func send(even, odd chan<- int, quit chan<- bool) { // sending onto Channel
	for i := 0; i < 10 ; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// recieve channel
func recieve(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("The value recieved from even channel", v)
		case v := <-odd:
			fmt.Println("The value recieved from odd channel", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}
