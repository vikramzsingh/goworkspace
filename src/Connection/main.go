package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
type Trainer struct {
	Name string
	Age int
	City string
}
*/

func main() {
	// set client options
	// The Client Options are used to set the connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// connect to mongodb
	//you must pass a context and option.ClientOption object to mongo.connect()
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// To close the connection at Exit of main
	defer client.Disconnect(context.TODO())

	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to MongoDB!")
	fmt.Println(databases)
}
