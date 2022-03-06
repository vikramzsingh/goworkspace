package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo) // closest match run
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")                                     // retriving value q from FORM
	w.Header().Set("Content-Type", "text/html; charset=-utf-8") // fro writing some html

	h := req.Header // you can use further function like Get("Accept")
	fmt.Fprintln(w, h)

	// closest match run
	io.WriteString(w, `
	<form method="POST">
		<input type="text" name="q">
		<input type="submit">
   </form>
   <br>`+v)
}
