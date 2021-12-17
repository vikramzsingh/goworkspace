// Hands-on-exercise-4 video 88

//	Using a COMPOSITE LITERAL:
//		create an ARRAY which holds 5 VALUES of TYPE int
//		assign VALUES to each index position.

package main

import "fmt"

func main() {
	// array created with composite literal
	// TYPE followed by {}
	x := [5]int{1, 2, 3, 4, 5} // this is array as it has size, slice does not have a size
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x) // TYPE is [5]int
}
