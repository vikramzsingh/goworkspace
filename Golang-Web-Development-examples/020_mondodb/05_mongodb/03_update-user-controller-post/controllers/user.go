package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"Golang-Web-Development/020_mondodb/05_mongodb/03_update-user-controller-post/models"
	"log"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s} // return type is pointer, and pointer stores memory address.
								// Hence &UserController{s} is returned
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}

	// at this point no data is fetched from request body, so struct is empty
	fmt.Println("data in u: ", u) // empty struct
	fmt.Println("Id without changed", u.Id)
	// get json from request body and Decode into Go
	err := json.NewDecoder(req.Body).Decode(&u)
	/*if err != nil { // Its throwing error invalid objectId in json, as 777- ObjectIdhex not used
		fmt.Println(err)
	}*/

	// create a bson ID
	u.Id = bson.NewObjectId()
	fmt.Println("changed Id: ", u.Id)

	// func CollectionNames:- attached to *Database
	// returns []strings, error
	// []strings contains collections
	sx, err := uc.session.DB("store").CollectionNames()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("List of collections:-")
	for _, c := range sx {
		fmt.Println(c)
	}

	// store user in mongodb
	uc.session.DB("store").C("customers").Insert(u) // directly storing in mongoDB in Json document

	// Marshal into JSON
	// Go to JSON
	bs, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintln(w, "%s\n", bs)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}