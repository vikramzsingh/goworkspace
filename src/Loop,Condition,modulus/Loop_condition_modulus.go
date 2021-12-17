// Loop_Condition_modulus(% operator) video 63

package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 { // modulus(%) operator
			fmt.Println(i)
		}
	}
}
