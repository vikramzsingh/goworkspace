// Print every No. from 1 to 10,000

package main

import "fmt"

/*
func main(){
	x := 1
	// this type of for loop work as while loop, And in Go there is no do-while Loop
	for x <= 10000{
		fmt.Println(x)
		x++
	}
}
*/
// or

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}
