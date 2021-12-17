package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{2, 4, 6, 3, 8, 9, 7}
	xs := []string{"James", "Q", "M", "Monneypeny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("-------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
