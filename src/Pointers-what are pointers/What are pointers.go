// Pointers_what_are_pointers

package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // address type is *int (Pointer to int)

	// & gives you the address.
	// * gives you the value stored.

	var b *int = &a // if you want to store address, you have to make same type of that address type. i.e *int
	fmt.Println(b)
	fmt.Println(*b) //when i have address. I am Defrencing an address.
	// or give me the value that is at this address.
	fmt.Println(*&a) // give me the address of a (&a) and defefrence  the address of a that is *&a

	c := &a
	fmt.Println(*c)

	*b = 43 // as address of a (&a) stored in b, Hence by derefrencing b (*b) we are storing value 43 in a VARIABLE.
	// *b means pointing to stored address of a.
	fmt.Println(a)
	fmt.Println(b)
}
