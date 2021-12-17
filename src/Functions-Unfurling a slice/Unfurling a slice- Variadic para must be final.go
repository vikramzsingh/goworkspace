// Functions_Unfurling_a_slice
// In variadic parameter function, if multiple parameter is used, then the variadic parameter must be final

package main

import "fmt"

func main() {
	xi := make([]int, 5, 5)
	xi = []int{2, 3, 4, 5, 6, 7, 8, 9}

	s := sum1("DJ", xi...) // Unfurnling the above slice of type []int, in order to change into int type with "..." operator
	// type slice xi is Unfurled
	fmt.Println("Total is:", s)
	xi = append(xi, 454, 45, 56) // original length is 8, we provided 5 capacity, which can go up to DOUBLE in make function
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}

// variadic parameter must be final parameter in variadic function parameter
func sum1(s string, x ...int) int { // Varidaic means zeor 0 or more of type int
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for index,", i, " Value", v, " is added to sum:", sum)
	}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(s)
	return sum
}

//func sum(x ...int, s string) // not possible if variadic parameter is use it must be final parameter
//func sum(s string, x ...int) // only this is possible, variadic parameter must be final parameter in variadic function parameter
