// Functions_variadic_parameter video 108

package main

import "fmt"

func main() {
	s, s1 := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is:", s)
	fmt.Println("Slic of int by VARIADIC PARAMETER:", s1)
}

func sum(x ...int) (int, []int) { // creating Variadic parameter (x ...int) x then i'm having unlimited number of int
	fmt.Println(x)
	fmt.Printf("%T\n", x) // as variadic parameter is given, so x automatically converter into slice of int

	//var sum int = 0
	sum := 0
	for i, v := range x {
		//sum = sum + v
		sum += v
		fmt.Println("For item in index position, ", i, " we are now adding, ", v, " to the total which is now: ", sum)
	}
	fmt.Println("The total is:", sum)

	return sum, x // as variadic parameter is given, so automatically converter into slice of int
}

// we pass individual value of type int.
// these values stored in VARIABLE x
// Variadic parameter takes Unlimited Number of int/string/etc.
// and then turn that into a slice of int/ slice of string/ etc

/// func (r reciever) identifier(parameter(s)) (return(s)) { code... }
