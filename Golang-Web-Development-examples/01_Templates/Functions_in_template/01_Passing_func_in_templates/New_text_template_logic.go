package main

import (
	"os"
	"text/template"
)

func main() {
	tpleg := template.Must(template.New("somet").Parse("here is text in the template"))
	tpleg.ExecuteTemplate(os.Stdout, "somet", nil)
}

// I used Parse, I have not used ParseFiles, ParseGlob
// so in case of Parse there is not file(html, gohtml) which going to call out during execution
// So we use New before Parse our data into string "somte" to execute code
