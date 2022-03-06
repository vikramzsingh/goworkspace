package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "Vikram",
		Age: 24,
	}
	fmt.Println(p1)

	// update without pointer not possible
	updateStructWithoutPointer(p1)
	fmt.Println("update without pointer not possible", p1)

	{
		// update within same code block is possible
		p1.Name = "Jerry"
		fmt.Println("update within same code block is possible", p1)
	}
	// ALSO WORKS HERE
	// update within same code block is possible
	// p1.Name = "Jerry"
	// fmt.Println("update within same code block is possible", p1)
	

	// update with pointer is possible
	updateStructWithPointer(&p1)
	fmt.Println("update with pointer is possible", p1)
}

func updateStructWithoutPointer(person Person) {
	person.Name = "Bobby"
	fmt.Println(person.Name)
}

func updateStructWithPointer(person *Person) {
	person.Name = "Bobby"
	fmt.Println(person.Name)
}