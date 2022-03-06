package main

import (
	"log"
	"monkey-paching-with-single-handle/ota"

)

func main() {
	updateId := "qwerty"
	myFunc(updateId)
	log.Println("---END---")
}

func myFunc(updateId string) {
	// res, err := ota.GetOTAUpdate(updateId) // do not call function directly, call through global variable, see below line.
	res, err := ota.GetOTAUpdateFunc(updateId)
	if err != nil {
		log.Println("error:", err)
	}
	log.Println("RESP:", res)
	// return res
}

