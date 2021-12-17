// star pattren 2
package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
