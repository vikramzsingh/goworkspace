// Slice__append_to_a_slice video 81

package main

import "fmt"

func main() {
	// slice with composite literal
	x := []int{1, 2, 3, 4, 42}
	fmt.Println(x)

	// func append(slice []T, elements...T )
	// append function return SLICE ( []T ) of same type.
	// Note:- []T is []Type, T is type of elements

	x = append(x, 66, 88, 99, 1041)
	fmt.Println(x)

	// if we have two Slice
	y := []int{233, 455, 655, 987}
	x = append(x, y...) // y... add the entire slice y to x
	fmt.Println(x)

}
