// Slice_deleting_from_a_slice video 82

package main

import "fmt"

func main() {
	// create slice of int with composite  literal
	x := []int{2, 3, 4, 5, 42, 66, 77, 99, 1041}
	fmt.Println(x)
	fmt.Println("Deleting 4,5 form slice x:")
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
