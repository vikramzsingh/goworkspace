// Slice_for_range video 79

package main

import "fmt"

func main() {
	// slice of int
	x := []int{1, 2, 3, 5, 42}
	fmt.Println("length of slice of int:", len(x))
	//fmt.Println(cap(x))--> same as len() function
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println("range clause:")

	// for index(i)[0,1,2,3,4] value(v)[1,2,3,5,42], range over x(range x) which is my SLICE
	for i, v := range x {
		fmt.Println(i, v) // prints index i and value v
	}
	fmt.Println("This range clause throws index with underscore(_) and only prints value in v")
	for _, v := range x {
		fmt.Println(v) // prints index i and value v
	}

	// slice of string
	fmt.Println("String Slice:")
	s := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(s))
	//fmt.Println(cap(s))--> same as len() function
}
