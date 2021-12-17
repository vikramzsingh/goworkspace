// Map_add_element_and_range video 86

package main

import "fmt"

func main() {
	// create a map with composite literal
	// syntax of map:-
	// map[Key_Type]Value_Type{
	// 		key1:val1,
	//		key2:val2,
	// }
	m := map[string]int{
		"vikram": 23,
		"jean":   20,
		"karan":  42,
	}

	fmt.Println(m)
	fmt.Println(m["vikram"])
	fmt.Println(m["jean"])
	fmt.Println(m["karan"])
	fmt.Println("Key not present in map m, so value is:", m["Jin"]) // missing key zero value (as value type is int,its default value is zero)

	// way to add key and value in map
	m["vijay"] = 22
	for k, v := range m { // for key and value over range m(return key and value of the m(map). )
		fmt.Println(k, v)
	}

}
