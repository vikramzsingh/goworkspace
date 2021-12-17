package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	// evaluated but not executed with defer keyword until leaving code body.
	defer func() { // Anonymous function
		if r := recover(); r != nil { // r != nil, means there is a reported error
			fmt.Println("Recovered in f", r)
		}
	}() // self-executing function
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i)) // we passing 4 in panic it calls up in stack for recover with string in panic
		// first it execute defer if any available
		// Then it go to g, g is defered, then it go from g to recover with String in Panic
		//

		//panic("error!!")
	}
	defer fmt.Println("defer in g.", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
