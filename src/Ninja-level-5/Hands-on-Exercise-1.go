// hands on exercise-5 video 103
/*Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
●	first name
●	last name
●	favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors
*/

package main

import "fmt"

// create structure your own type person, whose underlying type is struct
type person struct {
	firstname string
	lastname  string
	favice    []string
}

func main() {

	p1 := person{
		firstname: "Vikram",
		lastname:  "Singh",
		favice:    []string{"strawberry", "vanella", "chocolate"},
	}

	p2 := person{
		firstname: "James",
		lastname:  "Bond",
		favice:    []string{"favour1", "flavour2", "flavour3"},
	}

	// for p1
	fmt.Println(p1)
	fmt.Println(p1.firstname, p1.lastname, p1.favice)

	for i, v := range p1.favice {
		fmt.Println(i, v)
	}
	fmt.Println()

	// for p2
	fmt.Println(p2)
	fmt.Println(p2.firstname, p2.lastname, p2.favice)
	//fmt.Println(p2.firstname, p2.lastname, p2.favice[2])

	for _, v := range p2.favice {
		fmt.Println(v)
	}

	fmt.Printf("%T", p1.favice)
}
