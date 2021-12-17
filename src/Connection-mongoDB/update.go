package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	// How long we should wait before we throw some kind of timeout error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// we are connected verifying this.
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}

	// to close the connection
	defer client.Disconnect(ctx) // defer close connection at exit of main function

	/*
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(databases)
	*/

	// creating new database
	quickstartDatabase := client.Database("quickstart")             // Database access
	podcastsCollection := quickstartDatabase.Collection("podcasts") // Collection access
	//episodesCollection := quickstartDatabase.Collection("episodes")

						//***** UPDATE *****//
	// ObjectIDFromHex creates a new ObjectID from a hex string.
	id, _ := primitive.ObjectIDFromHex("5fcd822936cadb997a24b15d") //this is hash

	result, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id}, // selection criteria
		bson.D{ // field to update
			{"$set", bson.D{{"author", "Bobby"}}},
		},
		)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v documents!\n", result.ModifiedCount)

	result, err = podcastsCollection.UpdateMany(
		ctx,
		bson.M{"title": "The Developer podcast"}, // selection criteria
		bson.D{ // update fileds
			{"$set", bson.D{{"author", "Vikram Singh"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v documents!\n", result.ModifiedCount)
}

