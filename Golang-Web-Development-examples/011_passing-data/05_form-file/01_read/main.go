package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo) // clossest matching run
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil) // using DefaultServeMux
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost { // MethodPost    = "POST"  --> const from http-package, equal to "POST" method

		// open file as well as retrive file from FORM
		f, h, err := req.FormFile("q") // not FormValue, it is FormFile for catching file form FORM
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your Information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr:", err)

		bs, err := ioutil.ReadAll(f) // require r io.reader
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // err.Error() prinmts whatever the error is.
			return
		}
		s = string(bs) // byte-type is converted into string, and stored into s VARIABLE (which is outside this block, accessible to the whole function)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// clossest matching run
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
