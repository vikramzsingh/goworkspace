// defer with panic

package main

import "fmt"

func main() {
	fmt.Println("Befor panic")
	defer random1()
	fmt.Println("After panic")
}

func random1() {
	defer func() { // Anonymous function
		if r := recover(); r != nil {
			fmt.Printf("The function recovered from panic with reason: %v", r)
		}
	}() // No argument in the function, this paranthesis run anonymous function
	// also take arguments if any.
	fmt.Println(1)
	panic("some reason")
	fmt.Println(2)
}
