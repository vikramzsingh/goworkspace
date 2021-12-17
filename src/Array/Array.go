// Array in data structure video 77

package main

import "fmt"

// Use slice instead of array, Golang says
func main() {
	//syntax:
	// array_name[length]Type
	var x [5]int // elements of same type

	//or
	// var x[5]int

	//var y [6]int // length is a part of its type, and its different from above array

	fmt.Println(x)

	x[3] = 42 // array index start from 0 to 4
	fmt.Println(x[3])
	fmt.Println(x)

	//or also a SLICE of int type
	array1 := []int{3, 4, 5, 6, 7}
	fmt.Println(array1)

	z := [5]int{1, 2, 3, 4, 5} //This is array, as Slice ob byte does not include size[5]
	fmt.Println(z)
}
