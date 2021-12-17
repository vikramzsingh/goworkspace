// Pointers_Pointer1

package main

import "fmt"

func main() {
	a := 5
	b := &a // assigning b to a pointer of a

	fmt.Println(a, b)  // value in a and address of a
	fmt.Println(a, *b) // value in a and value at this adderss( b is dereffrenced )

	fmt.Println(a, &a) // value in a and address of a
	fmt.Println(b, &b) // value in be and address of b
	fmt.Println(*&b)   // b stores address of a

	*b = 10
	fmt.Println(a)
	fmt.Println("Address at this pointer:", &*b)

	var c *int = &a // creating pointer c
	fmt.Println(c)
}
