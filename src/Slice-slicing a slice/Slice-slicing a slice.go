// Slice_slicing_a_slice video 80

package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	fmt.Println("With for Loop:")
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println("With for Loop and len():")
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println("With range clause:")
	for i, v := range x { // for index and value over ramge in x(x is a SLICE)
		fmt.Println(i, v)
	}
}
