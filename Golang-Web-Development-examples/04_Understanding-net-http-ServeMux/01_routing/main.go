package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		fmt.Fprintln(res, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}
