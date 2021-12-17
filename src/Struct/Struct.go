// Struct Video 98

package main

import (
	"fmt"
)

// it is goes type then we give it a name(person) identifire and the underlying type is struct
// this is user-defined type, the type is person and the underlying type is struct
type person struct {
	first string // first is going to be a string
	last  string // last is going to be a string
	age   int    // last is going to be a int
}

func main() {
	p1 := person{ // in this p1 variable becomes person type automatically.
		first: "james",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Vikram",
		last:  "Singh",
		age:   23,
	}

	/*
		or
		var p1,p2 person

		p1 = person{
			first: "james",
			last: "Bond",
		}

		p2 = person{
			first: "Vikram",
			last: "Singh",
		}
	*/

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age) // its similarly works like object in java
	fmt.Println(p2.first, p2.last, p2.age)

	// slice of struct person, p1 and p2 struct added to slice of person
	//ps := []person{p1,p2}
	var people []person
	people = []person{p1, p2}
	fmt.Println(people)
}
