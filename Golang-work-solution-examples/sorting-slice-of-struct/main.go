package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name         string
	Age          int
	MobileNumber int64
}

func main() {

	// mobile number in descending order
	person := []Person{
		{
			"Vikram",
			24,
			1000,
		},
		{
			"Vijay",
			25,
			999,
		},
		{
			"Jhon",
			26,
			998,
		},
		{
			"Rick",
			27,
			997,
		},
	}
	fmt.Println("before sorting:", person)
	
	sort.SliceStable(person, func(i, j int) bool { return person[i].MobileNumber < person[j].MobileNumber })
	isSorted := sort.SliceIsSorted(person, func(i, j int) bool { return person[i].MobileNumber < person[j].MobileNumber })
	if (isSorted == true) {
		fmt.Println("after sorting:", person)	
	}
}
