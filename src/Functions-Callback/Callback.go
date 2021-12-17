// Functions_Callback

package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	x := sum(xi...) // unfurling slice
	// or
	// x := sum(2,3,4,5,6,7,8,9)

	fmt.Println("All Numbers:", x)

	s2 := even(sum, xi...)
	fmt.Println("Even Numbers:", s2)

	s3 := odd(sum, xi...)
	fmt.Println("Odd Numbers:", s3)

	d := foo() // d stores the func expression of returned anonymous function
	d()        // with this paranthesis () we are able to run the anonymous function

}

func sum(xi ...int) int {
	//fmt.Printf("%T\n", xi)

	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// this function takes function as a argument with signature of above function [ (xi ...int) int ]
// assigned function to f
// function as a argument also reqired signature of that function
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...) // before returning it runs sum() function to add up Even Numbers
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...) // before returning it runs sum() function to add up Odd Numbers
}

func foo() func() {
	return func() {
		fmt.Println("helloo return func()")
	}
}
