// Building on the previous hands-on exercise, create a program that uses “else if” and “else”

package main

import (
	"fmt"
)

func main() {
	x := 100
	if x == 10 {
		fmt.Println("If Statement in action")
	} else if x == 100 {
		fmt.Println("else If Statement in action")
	} else {
		fmt.Println("else Statement in action")
	}
}
