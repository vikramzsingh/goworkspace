// Print every rune code point of the uppercase alphabet three times. Your output should look like this:
//65
//	U+0041 'A'
//	U+0041 'A'
//U+0041 'A'
//66
//	U+0042 'B'
//	U+0042 'B'
//	U+0042 'B'
// … through the rest of the alphabet characters

package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			// UTF-8 code or rune
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}
