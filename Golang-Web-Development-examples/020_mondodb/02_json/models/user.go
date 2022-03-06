package models

type User struct { // keep Capital letter so it could be impoted outside
	Name   string `json:"Name"`
	Gender string `json:"gender"` 
	Age    int	  `json:"age"`
	Id     string `json:"id"`
}