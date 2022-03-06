package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"log"
	"fmt"
	"Golang-Web-Development/020_mondodb/04_controllers/models"
)

type UserController struct {} // empty struct

// with pointer you can access address of UserController struct.
// Then its Fields, elements are accesseble
// Important:- In GO-Lang everything work with value (pass by value)
func NewUserController() *UserController { // for inhancing speed
	return &UserController{} // As pointer stores the memory address, so address is returned to caller
}
/* also works
func NewUserController() UserController {
	return UserController{}
}
*/

// As GetUser attatched to UserController its funcanality available through UserController Struct
// GetUser and also access any value of UserController struct

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name: "James",
		Gender: "Male",
		Age: 54,
		Id: p.ByName("id"), // when link at index is triggered, /user/:id url is its target with id (9872309847), and getUser func is executed
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(uj)
	// fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{} // needed variable to store JSON data in GO

	// get JSON data from request body
	// encode/decode for sending/receiving JSON to/from a stream
	// it automatically decode JSON to GO and store data at memory adderss of the &u (VARIABLE) of type User struct, means automatically stored in struct( struct has json tags)
	json.NewDecoder(req.Body).Decode(&u) // JSON to GO

	// change Id of User struct
	u.Id ="007"

	// checking
	fmt.Println(u)

	// Again Marshal into JSON
	ujson, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", ujson)
}


func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
