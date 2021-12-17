//Loop_Printing_ascii video 60

package main

import "fmt"

func main() {
	// as for ascii character value we used 33('!') to 122('z')
	for i := 33; i <= 122; i++ {
		// here %v prints value of i in loop
		// here %T prints its type
		// %#X for hexa-decimal value
		// %U prints only unicode(UTF-8 code)
		// But %#U prints both unicode and Character
		fmt.Printf("%v%T\t\t%#x\t\t%#U\n", i, i, i, i)
	}
}
