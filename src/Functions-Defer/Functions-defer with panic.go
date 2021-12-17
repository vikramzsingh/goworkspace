// defer with panic

package main

import "fmt"

func main() {
	fmt.Println("Befor panic")
	defer random()
	fmt.Println("After panic")
}

func random() {
	defer fmt.Println(0)
	defer fmt.Println(1) // even function is in panic, the statement or function with defer KEYWORD will run
	// as the function get panic, but the 0,1 will still prints out
	fmt.Println(2)
	defer panic("Here we have panic")
	defer fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
}
