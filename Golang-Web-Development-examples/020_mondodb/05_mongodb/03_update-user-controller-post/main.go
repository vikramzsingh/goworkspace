package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Golang-Web-Development/020_mondodb/05_mongodb/03_update-user-controller-post/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/used/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

// connection to mongo
func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb://localhost") // s contains connection or *mgo.Session
	if err != nil {
		panic(err)
	}

	return s
}