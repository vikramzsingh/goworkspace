// Struct_Anonymous_structs video 101
// video 102 is Therotical about oops

package main

import "fmt"

// create a structure(struct)
// type struct-Type/Nane struct-keyword(underlying type){
// field Type
// Embedded struct
//}
/*
type person struct{
	first string
	last string
	age int
}*/

func main() {
	/*
		p1 := person{
			first: "James",
			last: "Bone",
			age: 32,
		}*/

	// Anonymous structure
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
