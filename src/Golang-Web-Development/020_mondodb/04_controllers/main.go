package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Golang-Web-Development/020_mondodb/04_controllers/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser) // route
	r.DELETE("/delete/:id", uc.DeleteUser) // route plus params
	http.ListenAndServe("localhost:8080", r)
}