// Introduction to packages video 26
package main

import (
	"fmt"
)

/*
func main() {
	n, err := fmt.Println("Hello World") // here err stores notthing so it err prints nil, which means no error
	fmt.Println(n)// n stores number of character in string"Hello World"
	fmt.Println(err) // if error raise it prints, here no error occured so it printer nil


	n, err := os.Open("file.ext")// here err stores notthing so it err prints nil, which means no error
	fmt.Println(n) // here n stores notthing so n is nil
	fmt.Println("Error Details:",err)

}
*/

// or

func main() {
	n, _ := fmt.Println("Hello World") // underscore character used to throw away returns // or discard it, although Println function works, and only data stored in _(blank identifire) is discarded

	// Important:-
	// _ throw away return (which is Number of character in string"Hello World"==12 ) or throw away first thing in data
	// n also store Number of character in the string

	fmt.Println(n) //2fmt.Println(_) if you use this as a variable, the message pop-up: cannot use _ as value

	//fmt.Println(a ...interface{}) (n int, err error)
	// as Println function return two values: n int, err error
	// n int to store Number of Character/variable
	// err error stores error if any error Occur
	// Hence, we need two variable n, _
	// n to store number of charcters (Acc. to Println() function)
	// _ to store other value which is error in this case( Acc. to Println() function).

	// also adjust n, _ acc to returned values of Println() function (n int, err error)
	// i.e n, _ check adjustment similarity (n int, err error)
}
