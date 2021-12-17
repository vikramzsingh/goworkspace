// conversion_not_casting video 33

package main

import "fmt"

var a int

type hotdog int // we created our own type whose underlying-type is int

var b hotdog // now we declared a variable of hotdog type(which is our own type variable)

func main() {
	a = 42 // int type
	b = 43 // hotdog type
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// conversion of hotdog-type(variable b) into int-type(variable b)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
