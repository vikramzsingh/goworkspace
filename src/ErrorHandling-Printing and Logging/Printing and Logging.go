// ErrorHandling_Printing_and_Logging

package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
		//log.Println("error happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}
