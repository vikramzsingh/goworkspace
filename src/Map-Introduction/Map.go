// Map_Introduction video 85

package main

import "fmt"

func main() {
	// Composite literal:- TYPE followed by{}, i.e TYPE{}
	// TYPE is-> map[string]int
	// followed by {}

	// Syntax of map:-
	// map[Key_Type]Value_Type{
	// 			key1: value1,
	//          key2: value2,
	// }

	m := map[string]int{
		"Vikram":    23,
		"Vijay":     22,
		"James":     36,
		"moneypeny": 53,
	}

	fmt.Println(m)
	fmt.Println(m["Vikram"])
	fmt.Println(m["Vijay"])
	fmt.Println(m["James"])
	fmt.Println(m["moneypeny"])
	fmt.Println(m["Jin"]) // missing entry from a zero value, i.e key not present in map, SOLUTION is check comma ok idiom(true/false)

	// way to fetch single value from single key
	val := m["James"]
	fmt.Println("James age:", val)

	/*
		//Way to check for the given key is there any value present in map
		v, ok :=m["Vikram"]

		if ok{
			fmt.Println("Value in v variable is the value of key(Vikram) is:",v)
			fmt.Println("comma ok idiom/variable gives us true/false(boolean):",ok)
		}
	*/

	// NOTE:- if statement; condition{} --> e.g
	// if x := 42; x == 42{ // it also validate scope of x variable
	// 		fmt.Println(Some variable)
	// } // here a used example below

	//Way to check for the given key is there any value present in map
	if v, ok := m["Vikram"]; ok {
		fmt.Println("Value in v variable is the value of key(Vikram) is:", v)
		fmt.Println("comma ok idiom/ok variable gives us true/false(boolean):", ok)
	}

	if v, ok := m["Jin"]; ok { // Jin is not in map, so ok becomes false, hence if statement does not run
		fmt.Println("Value in v variable key(Vikram) is:", v)
		fmt.Println("ok variable gives us true/false(boolean):", ok)
	}
}
