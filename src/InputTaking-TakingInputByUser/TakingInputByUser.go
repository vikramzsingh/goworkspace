// InputTaking_TakingInputByUser

package main

import "fmt"

func main() {
	var a, b, c int
	var d, e, f int

	//fmt.Println("Enter First Number:",a,"Possible")
	fmt.Println("Enter First Number:")
	fmt.Scanln(&a) // used to take input from the user in the next line
	fmt.Println("Enter Second Number:")
	fmt.Scan(&b) //used to take input from the user in the same
	// line
	c = a + b
	fmt.Println("The Addition of Two numbers is:", c)

	// Another way to take input
	// The above Technique is better, and suggestiable
	fmt.Println("\n")
	fmt.Printf("Enter first Number for subtraction:")
	fmt.Scanf("%v", &d)
	fmt.Printf("Enter second Number for subtraction:")
	fmt.Scanf("%v", &e)
	f = d - e
	fmt.Printf("The Subtraction of Two numbers:%v", f)

}
