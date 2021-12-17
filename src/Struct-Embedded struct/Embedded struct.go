// Embedded_structs video 99
// video 100 is theory about Reading Documentation

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Embedded struct means struct inside a struct (structure)
type SecretAgent struct {
	person // in Embedded Struct no need to declare type, its type is of that struct(person)
	// Also called anonymous field or Embedded field
	first string
	ltk   bool
}

func main() {
	// automatically becomes type of SecrateAgent.
	sa1 := SecretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "something collision",
		ltk:   true,
	}

	p2 := person{
		first: "vikram",
		last:  "Singh",
		age:   23,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	// collision case by first
	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.ltk)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk) // these fields are part of person embedded struct, but they got promoted to SecretAgent
	fmt.Println(p2.first, p2.last, p2.age)

	// checking type
	fmt.Printf("%T, %T, %T, %T, %T ", sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Printf("\n%T, %T, %T, ", p2.first, p2.last, p2.age)
	fmt.Printf("\n%T", sa1)
	fmt.Printf("\n%T", sa1.person)
	fmt.Printf("\n%T", p2)
}
