// creating_your_own_type video 32

package main

import "fmt"

var a int

type hotdog int // we created our own type whose underlying-type is int

var b hotdog // now we declared a variable of hotdog type(which is our own type variable)

func main() {
	a = 42 // int type
	b = 43 // hotdog type
	fmt.Println(a)
	fmt.Printf("%T\n", a) // printing type of a

	fmt.Println(b)
	fmt.Printf("%T\n", b) // printing type of b
	// a = b // Now we assigning hotdog-type to a int-type variable which is not possible
	// Go is about type we cannot be able to assign other than int-type in variable a
}
