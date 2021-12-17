// Functions_Methods video 111
// Anonymous type becomes the field name
package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type SecretAgent struct {
	person // embedded structure
	ltk    bool
}

//func speak(s SecretAgent){ code ...}
// when we do this
//func (s SecretAgent) speak(){ code ...}
// (s SecretAgent), means attaching function speak() to the value of that type (SecretAgent)
// which means we have access to the values of type SecretAgent structure
func (s SecretAgent) speak() { // attaching the function speak() to the value of type SecretAgent ( like sa1, sa2 ), so we can have access to the all values of SecretAgent
	// by attaching speak() to SecretAgent, speak() becones of type SecretAgent
	fmt.Println("The Name is", s.first, s.last)
	fmt.Println(s.ltk)
}

func main() {
	sa1 := SecretAgent{
		person: person{ // Anonymous type becomes the field name
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := SecretAgent{
		person: person{ // Anonymous type becomes the field name
			first: "Miss",
			last:  "MoneyPenny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Printf("%T", sa1) // sa1 is a value of type SecretAgent
	fmt.Println(sa2)

	//speak(sa1)
	sa1.speak()
	sa2.speak()
}
