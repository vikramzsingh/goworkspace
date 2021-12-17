// Functions_Interface_amd_polymorphism

package main

import "fmt"

type person struct {
	first string
	last  string
}

type SecretAgent struct {
	person
	ltk bool
}

// attaching method to SecretAgent
// then method also becomes type of SecretAgent
func (s SecretAgent) speak() {
	fmt.Println(s)
	fmt.Println(s.first, s.last)
	fmt.Println(s.ltk)
	fmt.Printf("%T\n", s.speak)
}

// attaching method to person
// then method also becomes type of person
func (p person) speak() {
	fmt.Println(p)
	fmt.Println(p.first)
	fmt.Printf("%T\n", p.speak)
}

// INTERFACE
type human interface {
	speak() // Anybody of any type who has method speak is also of type human
}

func boo(h human) { // human INTERFACE as a parameter
	fmt.Println(h) // human INTERFACE reciving multiple type of values.
	// these multiple types are SecretAgent and person
	// Both SecretAgent and person has value, is also of human type

	// this assestion only work with interface as human type is INTERFACE
	switch h.(type) { // this as am assert; asserting, "x is of type this
	case person:
		fmt.Println(h.(person).first)
		fmt.Println(h.(person).last)
	case SecretAgent:
		fmt.Println(h.(SecretAgent).first)
		fmt.Println(h.(SecretAgent).last)
		fmt.Println(h.(SecretAgent).ltk)
	}
}

func main() {
	sa1 := SecretAgent{
		person: person{
			first: "Vikram",
			last:  "singh",
		},
		ltk: false,
	}

	sa2 := SecretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "money",
		last:  "penny",
	}

	sa1.speak() // method speak of type SecretAgent
	sa2.speak()
	p1.speak()
	fmt.Println("After attaching to SecretAgent and person:")
	fmt.Printf("%T\n", sa1.speak) // after attaching to SecretAgent
	fmt.Printf("%T\n", p1.speak)  // after attaching to person
	fmt.Println()

	//this is also polymorphism as sa1 is of two types SecretAgent and human
	boo(sa1) // as sa1 is of SecretAgent type and method speak is attached to SecretAgent, and Also INTERFACE human has method speak
	// Anybody of any type who has method speak is also of type human.
	boo(sa2)
	boo(p1)
}
