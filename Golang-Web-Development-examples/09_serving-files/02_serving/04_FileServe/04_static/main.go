package main

import (
	"log"
	"net/http"
)

func main() {
	// special case must required index.html page
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
