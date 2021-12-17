// video 30
package main

import "fmt"

var y string
var z int
var a float64
var b bool

func main() {
	fmt.Println("Printing the value of y", y, "ending") // This will print a space" " at the place of y as space" " is default value of string-type
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"
	fmt.Println(y) // now we  Assigned String-value to y variable
	fmt.Printf("%T\n", y)

	fmt.Println(z) // Default value of z variable
	fmt.Printf("%T\n", z)
	//fmt.Printf("%V\n", z)// can also be used

	fmt.Println(a) // Default value of a variable
	fmt.Printf("%T\n", a)

	fmt.Println(b) // Default value of b variable
	fmt.Printf("%T\n", b)
}
