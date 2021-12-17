// Slice__make video 83

package main

import "fmt"

func main() {
	// make(slice, length, capacity)
	// capacity is underlying size of the array/slice
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999 // at 10 th position

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3221)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3222)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3223) // at this, it gives the double size of the underlying array
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	/*
		x = append(x, x...)
		fmt.Println(x)
		fmt.Println(len(x))
		fmt.Println(cap(x))
	*/
}
