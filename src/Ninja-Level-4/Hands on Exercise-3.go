//Hands on Exercise-3 video 90

//use SLICING to create the following new slices which are then printed:
//●	[42 43 44 45 46]
//●	[47 48 49 50 51]
//●	[44 45 46 47 48]
//●	[43 44 45 46 47]

package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//s = s[:5]
	//fmt.Println(s) or

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
}
