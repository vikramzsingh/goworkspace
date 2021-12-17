// Functions_Unfurling_a_slice video 109

package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(xi...) // Unfurnling the above slice of type []int, in order to change into int type with "..." operator
	// tyhe slice xi is Unfurled
	fmt.Println("Total is:", s)
}

func sum(x ...int) int { // Varidaic means zeor 0 or more of type int
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for index,", i, " Value", v, " is added to sum:", sum)
	}

	return sum
}

//func sum(x ...int, s string) // not possible if variadic parameter is use it must be final parameter
//func sum(s string, x ...int) // only this is possible, variadic parameter must be final parameter in variadic function parameter
