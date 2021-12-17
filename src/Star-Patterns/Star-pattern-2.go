// error in program
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("")
			for k := 1; k <= j+i; k++ {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
}
