// Exploring-Type(Checking of type of variable{int,string,float etc}) video 29
package main

import "fmt"

var y = 42 // globle variable declared of integer-Type

// Declared the variable with identifier "z"
// is of type string
// and assigned value "Power ranger, not spiderman"
var z string = "Power ranger, not spiderman" // variable declaration with its type(string) and assigned a value "Power ranger, not spiderman"
// go language is STATIC programming language not a DYANAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // Printf function used to identify or check type of variable
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	/* "z" variable is declared of STRING type
	//so we cannot assing a INTEGER value to the "z variable"
	// (error raised)cannot use 43 (type int) as type string in assignment
	z = 43
	fmt.Println(z)
	fmt.Printf("%T", z)
	*/
}
