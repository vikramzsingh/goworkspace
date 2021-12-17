package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Podcast struct {
	//if value is nil or empty string,not included with omitempty
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string			   `json:"title" bson:"title,omitempty"`
	Author string			   `json:"author" bson:"author,omitempty"`
	Tags   []string			   `json:"tags" bson:"tags",omitempty`
}

type Episode struct {
	ID          primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	Podcast     primitive.ObjectID      `json:"podcast" bson:"podcast,omitempty"`
	Title       string                  `json:"title" bson:"title,omitempty"`
	Description string					 `json:"description" bson:"description,omitempty"`
	Duration    int32                   `json:"duration" bson:"duration,omitempty"`
}

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

	// ***** FIND query ***** //

	// Retrive manny
	var podcast []Podcast
	podcastCursor, err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	if err = podcastCursor.All(ctx, &podcast); err != nil {
		panic(err)
	}
	fmt.Println(podcast)
	fmt.Println(podcast[0].Author)
	fmt.Println(podcast[1].ID)
	fmt.Println(podcast[1].Title)

	defer podcastCursor.Close(ctx)
	/*
	// use bson.M{} to get all data from mongodb
	cursor, err := episodesCollection.Find(ctx, bson.M{}) // to retrive all data
	if err != nil {
		log.Fatal(err)
	}


	// to store fetched data
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(episodes)

	for _, episode := range episodes {
		fmt.Println(episode)
		fmt.Println(episode["title"])
	}




	for cursor.Next(ctx) {
		var episode bson.M
		// Decode will unmarshal the current document into val and return any errors from the unmarshalling process without any modification
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episode)
	}



	// filterCursor is cursor for all episode having duration 25
	filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration":25})
	if err != nil {
		log.Fatal(err)
	}

	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)

	opts := options.Find()
	opts.SetSort(bson.D{{"duration", -1}})// -1 for decending, 1 for assending
	sortCursor, err := episodesCollection.Find(ctx, bson.D{
		{"duration", bson.D{
			{"$gt", 24},
		}},
	}, opts)

	var episodeSorted []bson.M
	if err = sortCursor.All(ctx, &episodeSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodeSorted)
	*/
}



