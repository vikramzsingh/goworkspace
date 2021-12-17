// video 37
package main

import "fmt"

// create own type with underlying-type int
type hotdog int //type vikram int
var x2 hotdog   //var x2 vikram

func main() {
	fmt.Println(x2)        // print value of x
	fmt.Printf("%T\n", x2) // print Type of x
	// assign 42 to variable x2 using "=" operator
	x2 = 42
	fmt.Println(x2)
}
