// Loop_init_condition_post video 56
/*
Different way of using For-Loop
package main
import "fmt"
func main() {
The most basic type, with a single condition.

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
A classic initial/condition/after for loop.

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

    for {
        fmt.Println("loop")
        break
    }
You can also continue to the next iteration of the loop.

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
*/

package main

import "fmt"

func main() {

	/* syntax of for-loop
	for init; condition; post {}
	*/

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

}
