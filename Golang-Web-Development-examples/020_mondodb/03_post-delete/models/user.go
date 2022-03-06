package models

type User struct { // capital for export outside the package
	Name   string `json:"name"` // capital and small letters json tags does not affect.
	Gender string `json:"gender"`
	Age    int	   `json:"age"`
	Id     string  `json:"id"`
}