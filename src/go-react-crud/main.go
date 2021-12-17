package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	//"strconv"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User struct
type User struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name,omitempty"`
	Age     int                `json:"age" bson:"age,omitempty"`
	Address string             `json:"address" bson:"address,omitempty"`
}

// UserSave struct
type UserSave struct {
	//ID 			primitive.ObjectID `json:"_id,omitempty" bson="_id,omitempty"`
	FirstName string `json:"fname" bson:"fname,omitempty"`
	LastName  string `json:"lname" bson:"lname,omitempty"`
	Age       int    `json:"age" bson:"age,omitempty"`
	MobileNo  string `json:"mobileno" bson:"mobileno,omitempty"`
	Dob       string `json:"dob" bson:"dob,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Address   string `json:"address" bson:"address,omitempty"`
}

// UserGet for fetching user
type UserGet struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson="_id,omitempty"`
	FirstName string             `json:"fname" bson:"fname,omitempty"`
	LastName  string             `json:"lname" bson:"lname,omitempty"`
	Age       int                `json:"age" bson:"age,omitempty"`
	MobileNo  string             `json:"mobileno" bson:"mobileno,omitempty"`
	Dob       string             `json:"dob" bson:"dob,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Address   string             `json:"address" bson:"address,omitempty"`
}

// Database points to DB
var Database *mongo.Database
var ctx = context.TODO()

// var x = func(h http.Handler)http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Header.Get("Access-Control-Request-Method") != "" {
// 			// Set CORS headers
// 			header := w.Header()
// 			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
// 			header.Set("Access-Control-Allow-Origin", "*")
// 		}
// 		if r.Method == http.MethodOptions{
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}else{
// 			h.ServeHTTP(w,r)
// 		}

// 		// Adjust status code to 204
// 	})
// }

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}

	Database = client.Database("web_development")
	fmt.Println("connected to mongoDB at", client)
}

func main() {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "application/json", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/createuser", CreateUser).Methods("POST")
	router.HandleFunc("/getuser", GetUser).Methods("GET")
	router.HandleFunc("/updateuser/{email}", UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteuser/{email}", DeleteUser).Methods("DELETE")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

// Index handler
func Index(w http.ResponseWriter, req *http.Request) {

	result := Database.Collection("user")
	var data []User
	cur, err := result.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(w, `{error: %s}`, err.Error())
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var d User
		cur.Decode(&d)
		data = append(data, d)
	}
	resp, _ := json.Marshal(data)
	w.Write(resp)
	w.WriteHeader(http.StatusOK)

}

// CreateUser creates user
func CreateUser(w http.ResponseWriter, req *http.Request) {
	// handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// handlers.AllowedMethods([]string{"GET", "POST", "PUT","DELETE"})
	// handlers.AllowedOrigins([]string{"* "})
	//w.Header().Set("Content-Type", "application/json")

	if req.Method == http.MethodPost {
		// fmt.Println("Method is POST")
		// data := req.Body
		// data1 := req.Response
		// fmt.Println("Body data:", data)
		// fmt.Println("response data:", data1)

		var user UserSave
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			fmt.Println(err)
		}
		result, err := Database.Collection("users").InsertOne(ctx, user)
		if err != nil {
			fmt.Fprintf(w, `{error: %s}`, err.Error())
			return
		}
		fmt.Println("Inserted:", result)
	}
}

// GetUser for fetchin user
func GetUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("context-type", "application/json")
	var users []UserGet
	cur, err := Database.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user UserGet
		cur.Decode(&user)
		users = append(users, user)
	}
	resp, _ := json.Marshal(users)
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}

// UpdateUser for upadting
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	email := vars["email"]

	var user UserGet
	json.NewDecoder(req.Body).Decode(&user)
	fname := user.FirstName
	fmt.Println(fname)

	result, err := Database.Collection("users").UpdateOne(
		ctx,
		bson.M{"email": email},
		bson.D{
			{"$set", bson.D{{"fname", fname}}},
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("No. of Record Updated:", result.UpsertedCount)
}

// DeleteUser for upadting
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	email := vars["email"]
	result, err := Database.Collection("users").DeleteOne(
		ctx,
		bson.M{"email": email},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("No. of Record Deleted:", result.DeletedCount)
}
