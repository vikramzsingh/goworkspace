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
	episodesCollection := quickstartDatabase.Collection("episodes")

					// ****** INSERT ******//

	// To handle collection start creating data
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{"title","The Developer podcast"}, // key: value --> see in mongodb
		{"author","Vikram Singh"},
		{"tags", bson.A{"development", "programming", "codding"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastResult.InsertedID)

	// now we are going to insert many document
	// so we need []interface{}{ bson.D{} }
	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode #1"},
			{"description", "This is the first episode"},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode #2"},
			{"description", "This is the second episode"},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodeResult.InsertedIDs)

}
