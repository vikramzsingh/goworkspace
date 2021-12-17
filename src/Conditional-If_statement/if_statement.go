//Conditional_If_statement video 61

package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	if 2 != 2 {
		fmt.Println("006")
	}
	//2 == 2 is becomes ture, by applying !(not) it becomes false
	if !(2 == 2) {
		fmt.Println("007")
	}

	if !(2 != 2) {
		fmt.Println("008")
	}
}
