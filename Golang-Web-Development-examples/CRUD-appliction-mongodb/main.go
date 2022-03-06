package main

import (
	"Golang-Web-Development/CRUD-appliction-mongodb/controllers"
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	uc := controllers.NewUserController(getConnection())
	http.HandleFunc("/", uc.Index)
	http.HandleFunc("/createuser", uc.CreateUser)
	http.HandleFunc("/getuser", uc.GetUser)
	http.ListenAndServe("localhost:8080", nil)
}

func getConnection() (*mongo.Client, context.Context) {
	// connection
	// return *mongo.client, error
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}

	//Context, how long we wait before throwing error of timeout
	ctx := context.TODO() // context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// To close connection
	//defer client.Disconnect(ctx)

	// we are connected verifying this
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}

	return client, ctx
}
