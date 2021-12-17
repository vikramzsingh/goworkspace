// short-declaration-operator(variable-declaration) video 27
package main

import "fmt"

func main() {
	x := 42 //variable declaration
	// This short-declaration-operator(:=) can declare variable and assign value at the same time
	fmt.Println(x)

	x = 99 // This time (=)equal to is used because x is already declared, now we are assigning value(99) to x variable
	fmt.Println(x)

	y := 100 + 24 // this is a statement, expression as 100+24 is 124
	fmt.Println(y)

	z := "Hello Vikram"
	fmt.Println("Get up,", z)
}
