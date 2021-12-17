// Map_delete video 87

package main

import "fmt"

func main() {
	// create a map with COMPOSITE LITERAL:
	// map[Key_Type]Value_Type{
	// 			key1: value1,
	//			key2: value2,
	// }
	m := map[string]int{
		"vikram": 23,
		"vijay":  22,
		"james":  34,
	}

	fmt.Println(m)

	// syntax to delete:-
	// delete(<map_name>, "Key")
	delete(m, "james")
	fmt.Println(m)

	delete(m, "shiva ji") //No error is thrown by delete(built-in method), if you use a key which does not exist.
	fmt.Println(m)
	fmt.Println(m["vikram"])
	fmt.Println("key dose not exist, so value is zero:", m["shiva ji"])

	//To confirm you delete a key/value, verify that the key/value exists with the comma ok idiom.
	// if x := 42; x == 42{
	//		fmt.Println(x)
	// 	}
	if v, ok := m["vikram"]; ok {
		fmt.Println(v, ok)
		fmt.Println("value:", v)
		delete(m, "vikram")
	}
	fmt.Println(m)

	if v, ok := m["shiva ji"]; ok { // comma ok idiom become false as "shiva ji"(key) does not exist in map, so this will not run
		fmt.Println(v, ok)
	}
}
