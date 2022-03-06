package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Year   int     `json:"year"`
	Title  string  `json:"title"`
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}
 

func main() {
	// Initilize AWS session with valid Credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	// item to be Inserted in DynamoDB
	item := Item{
		Year:  2018, 
		Title: "Avengers",
		Plot: "SCI-fci",
		Rating: 4.0,
	}

	// Befor inserting into DynamoDB,
	// convert Item struct into map[string]*dynamodb.AttributeValue,
	// because PutItemInput{} has field named as Item that accepts data of type map[string]*dynamodb.AttributeValue,
	// In all MarshalMap takes a struct and gives back a map of type map[string]*dynamodb.AttributeValue,   
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	// Table name
	tableName := "Movies"
	input := &dynamodb.PutItemInput{
		Item:      av, // this field requires data of type == map[string]*dynamodb.AttributeValue
		TableName: aws.String(tableName),
	}

	// execute PutItem method to insert data into DynamoDB
	_, err = svc.PutItem(input) 
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	year := strconv.Itoa(item.Year)
	fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)
}
