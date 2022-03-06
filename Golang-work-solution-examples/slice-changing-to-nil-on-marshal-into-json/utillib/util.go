package utillib

type Exp struct {
	Name string `json:"name"`
	Age  int 	`json:"age"`
}

type Globle struct {
	Hits []Exp
}