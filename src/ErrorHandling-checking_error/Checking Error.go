// ErrorHandling_checking_error

package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello") // it returns (n int, err errro)
	if err != nil {                // err != nil, means their is a reported error
		fmt.Println(err)
	}
	fmt.Println(n)
}
