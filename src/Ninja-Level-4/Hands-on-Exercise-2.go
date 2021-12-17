//Hands on Exercise-2 video 89

//●	Using a COMPOSITE LITERAL:
//●	create a SLICE of TYPE int
//●	assign 10 VALUES
//●	Range over the slice and print the values out.
//●	Using format printing
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", s)

}
