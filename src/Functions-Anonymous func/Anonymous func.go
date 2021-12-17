// Functions_Anonymous_func video 113
// or Anonymous self-executing func

package main

import "fmt"

func main() {
	foo("Normal foo", "function") // normal function

	// Anonymous foo function
	// Anonymous function does not have any name, but it has parameters
	// At the end of Anonymous function paranthesis is used to run Anonymous function
	// These paranthesis also takes arguments(if any)
	func(s string, s1 string) {
		fmt.Println(s, s1)
	}("Anonymous foo", "function") // calling function and passing arguments

	func() {
		fmt.Println("Hello world")
	}() //In anonymous function paranthesies required to run anonymous function

	func(x int) {
		fmt.Println(x)
	}(42) //In anonymous function paranthesies required to run anonymous function
}

func foo(s string, s1 string) {
	fmt.Println(s, s1)
}
