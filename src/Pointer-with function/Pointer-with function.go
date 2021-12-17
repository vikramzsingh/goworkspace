// Pointer_with_function

package main

import "fmt"

func zero(z *int) { // getting address of x, as pointer to int type(*int)
	fmt.Println(z)
	*z = 0 // pointer to int type(*int), cannot assign int type
	// derefrencing pointer to int(*int), to store int type at address of x
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)       // this func require pointer to int type(*int) value(memory address), as para is pointer to int type(*int)
	fmt.Println(x) // x is 0
}
