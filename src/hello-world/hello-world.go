// Its video 25 hello world
package main

import "fmt"

func main() {
	fmt.Println("Hello World,")

	foo() // its a function creation or Function call

	fmt.Println("After running foo function")

	// for loop syntax
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

// Function declaration of foo
func foo() {
	fmt.Println("I'm Vikram singh")
}

func bar() {
	fmt.Println("and then we exited")
}

//control flow:
//(1) sequence
//(2) loop: iterative
//(3) conditional
