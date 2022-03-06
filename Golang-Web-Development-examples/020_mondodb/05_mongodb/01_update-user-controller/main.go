package main

import (
	"Golang-Web-Development/020_mondodb/05_mongodb/01_update-user-controller/controllers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	fmt.Println("This is UserController:", uc)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// connect to our local mongo
	// or its connection to data base of mongo
	s, err := mgo.Dial("mongodb://localhost") // returns *mgo.Session, error

	// check if connection error, is mongo running ?
	if err != nil {
		panic(err)
	}
	fmt.Println("This is *mgo.Session", s)
	return s
}
