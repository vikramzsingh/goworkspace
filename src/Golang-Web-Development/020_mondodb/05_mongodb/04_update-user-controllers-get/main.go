package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
	"fmt"
	"Golang-Web-Development/020_mondodb/05_mongodb/04_update-user-controllers-get/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb://localhost") // connection to mongodb
	if err != nil {
		fmt.Println(err)
	}
	return s
}