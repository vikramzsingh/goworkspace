// Hands on exercise 3 video 105
//●	Create a new type: vehicle.
//○	The underlying type is a struct.
//○	The fields:
//■	doors
//■	color
//●	Create two new types: truck & sedan.
//○	The underlying type of each of these new types is a struct.
//○	Embed the “vehicle” type in both truck & sedan.
//○	Give truck the field “fourWheel” which will be set to bool.
//○	Give sedan the field “luxury” which will be set to bool. solution
//●	Using the vehicle, truck, and sedan structs:
//○	using a composite literal, create a value of type truck and assign values to the fields;
//○	using a composite literal, create a value of type sedan and assign values to the fields.
//●	Print out each of these values.
//●	Print out a single field from each of these values.

package main

import "fmt"

type vehical struct {
	doors int
	color string
}

type truck struct {
	//vehical vehical // eda sirf provided type ho jande hai
	vehical   // par eda pura struct wali type ho jande hai, entire struct-type access ho jande hai
	fourWheel bool
}

type sedan struct {
	vehical
	luxury bool
}

func main() {
	t := truck{
		// vehical thats going to be a type vehical
		vehical: vehical{
			doors: 2,
			color: "White",
		},
		fourWheel: true,
	}

	/* also possible
	t := truck{
		vehical{
			4,
			"White",
		},
		true,
	}
	*/

	s := sedan{
		vehical{
			4,
			"red",
		},
		true,
	}

	//fmt.Println(t)
	//fmt.Println(s)

	fmt.Println(t.doors) //inner type promotion
	fmt.Println(t.color)
	fmt.Println(t.fourWheel)

	fmt.Println(s.doors)
	fmt.Println(s.color)
	fmt.Println(s.luxury)
}
