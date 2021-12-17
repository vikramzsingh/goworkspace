// Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	name := "James Bond"
	switch {
	case name == "James Bond":
		fmt.Println("James Bond, should print")
	case name == "Vikram":
		fmt.Println("Vikarm singh, should not print")
	default:
		fmt.Println("this is default case")

	}
}
