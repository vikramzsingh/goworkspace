package models

type User struct { // capital for expotrt outside
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int	  `json:age`
	Id     string `json:"id"`
}