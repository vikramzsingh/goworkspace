// Slice_Composite_literal video 78

package main

import "fmt"

func main() {
	//A composite literal is created by having the TYPE followed by CURLY BRACES and then putting the appropriate values in the curly brace area.
	//x := type{values} // eg. of COMPOSITE LITERAL

	// Slice of int type declared with Composite Literal syntax
	x := []int{4, 5, 6, 7, 8, 42}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println("99876"[2])   // Prints ASCII value of index 2
	fmt.Println("99876"[0:3]) // it prints value form 0 to 2, i.e 998[0,1,2]

	// String slice
	mySlice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(mySlice)
	fmt.Printf("%T\n", mySlice)
	fmt.Println("ABC"[2])   // it prints the ASCII Number of the given string (ABC),therefor it prints 67 ASCII no. of C in string ABC at index (65)A=0,(66)B=1,(67)C=2
	fmt.Println("ABC"[0:2]) // prints Character from 0 to 1,i.e AB[0,1]
}
