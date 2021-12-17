// Hands on exercise 2 video 104

package main

import (
	"fmt"
)

// structure of type person with underlying type is struct
type person1 struct {
	firstname string
	lastname  string
	favice    []string
}

func main() {

	p1 := person1{
		firstname: "Vikram",
		lastname:  "Singh",
		favice:    []string{"strawberry", "vanella", "chocolate"},
	}

	p2 := person1{
		firstname: "James",
		lastname:  "Bond",
		favice:    []string{"flavour1", "flavour2", "flavour3"},
	}

	// map[Key_Type]Value_Type{}
	m := map[string]person1{
		// the technique fetch the lastname of p1, p2 and use them as a Key_Type, and then we stored the struct p1 and p2 in map
		p1.lastname: p1,
		p2.lastname: p2,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("\nfor Key:", k)
		fmt.Println(v)
		fmt.Println(v.firstname) // because type is common in map and struct i.e person1 (it is a struct data)
		fmt.Println(v.lastname)
		for i, val := range v.favice {
			fmt.Println(i, val)
		}
	}
}
