// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var condition = true // assume as NextToken of AWS.
	count := 0

	// First way to perform do-while loop
	// 	for ok := true; ok; ok = condition {
	//         work()
	// }
	for ok := true; ok; ok = condition {
		fmt.Println("Success do while")
		count++
		fmt.Println(count)
		if count == 5 {
			condition = false
		}

	}

	// Second way to perform do-while loop
	// for {
	// 	work()
	// 	if !condition {
	// 			break
	// 	}
	// }
	fmt.Println("END")
}
