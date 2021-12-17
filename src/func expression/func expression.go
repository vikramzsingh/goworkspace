// func_expression

package main

import "fmt"

func main() {
	defer fmt.Println("Hello world")

	// In Golang func is a type just like other type, we can assign it to VARIABLE
	// function is like first class citizen
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	/*
		 func() {
			fmt.Println("my first func expression")
		}()
	*/

	g := func(x int) {
		fmt.Println(x)
	}
	g(42)
}
