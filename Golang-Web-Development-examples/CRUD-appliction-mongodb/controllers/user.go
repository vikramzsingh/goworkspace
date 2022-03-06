package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"html/template"
	"Golang-Web-Development/CRUD-appliction-mongodb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"encoding/json"
)
var TPL *template.Template
var ObjID = map[string]string{}

func init()  {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type UserController struct {
	client *mongo.Client
	ctx    context.Context
}

func NewUserController(c *mongo.Client, ctx context.Context) *UserController {
	return &UserController{c, ctx}
}


// INDEX
func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	// It is always Get Method
	if req.Method == http.MethodGet {
		err := TPL.ExecuteTemplate(w, "index.gohtml", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound) // 404
			return
		}
	}
}

// CREATE USER
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// Process form Submission

		fname := req.FormValue("fname")
		lname := req.FormValue("lname")
		email := req.FormValue("email")
		contactno := req.FormValue("contactno")
		dob := req.FormValue("dob")

		oid := primitive.NewObjectID() // create new objectID

		// store object id into map
		u := models.User{oid,fname, lname, email, contactno, dob}

		// stroe in mongodb
		storeDatabase := uc.client.Database("store") //attached to *client, return *Database
		usersCollection := storeDatabase.Collection("users") //attached to *Database, return *collection
		resultCursor, err := usersCollection.InsertOne(uc.ctx, &u) // attached to *Collection, return cursor/ documents
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Context-Type", "application/json")
		w.WriteHeader(http.StatusOK) //200
		fmt.Fprintf(w, "Data Inserted:\n %v\n", resultCursor.InsertedID)

		// marshal into json
		uj, err := json.Marshal(u)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "Json Document Inserted: \n%s ", uj)

		//convert objectId to Hex representation
		hexId := oid.Hex()
		fmt.Println(hexId) // Hex represented id

		//fmt.Println(oid.String()) // just string for storing

		//convert Hex representation ObjectId
		//objId, err := primitive.ObjectIDFromHex(hexId)
		//fmt.Println(objId)

		ObjID["id"] = hexId

		return
	}

	err := TPL.ExecuteTemplate(w, "CreateUser.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}
}

// GET USER
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request) {

		hexId := ObjID["id"] // get Hash Id from map
		oid, _ := primitive.ObjectIDFromHex(hexId)
		fmt.Println("oid: ", oid)
		// user struct
		u := models.User{}

		// get user from db
		storeDatabase := uc.client.Database("store")
		userCollection := storeDatabase.Collection("users")
		err := userCollection.FindOne(uc.ctx, bson.M{"_id":oid}).Decode(&u)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Context-Type", "application/json")
		w.WriteHeader(http.StatusOK) //200
		fmt.Fprintln(w, u)
		return
}
/*
func (uc UserController) UpdateUser(res http.ResponseWriter, req *http.Request) {
	hexid := ObjID["id"]
	oid, _ := primitive.ObjectIDFromHex(hexid)

	storeDatabase := uc.client.Database("store")
	userCollection := storeDatabase.Collection("users")
	userCollection.UpdateOne()
}
*/