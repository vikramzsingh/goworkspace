package main

import "fmt"

func main() {
	x := 0
	fmt.Println(&x)
	fmt.Println(x)
	foo1(&x) // address where int is stored
	fmt.Println(&x)
	fmt.Println(x)
}

func foo1(y *int) { // pointer to int (*int)
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43 // the value at this address set it at 43
	fmt.Println(y)
	fmt.Println(*y)
}
