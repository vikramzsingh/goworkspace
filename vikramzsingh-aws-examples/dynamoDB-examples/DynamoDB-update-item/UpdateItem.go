package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// initilize session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	// Create service to access DynamoDB
	svc := dynamodb.New(sess)

	tableName := "Movies"
	movieName := "Avengers"
	movieYear := "2018"
	movieRating := "0.5"

	input := &dynamodb.UpdateItemInput{
		// Takes values that we are about to Update
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(movieRating),
			},
		},
		// Table Name where update is required
		TableName: aws.String(tableName),
		// To Identify/Find out which OBJECT needs update
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String(movieYear),
			},
			"title": {
				S: aws.String(movieName),
			},
		},
		// Its has Multiple options you can check.
		ReturnValues: aws.String("UPDATED_NEW"),
		// Field Name which we are about to update
		UpdateExpression: aws.String("set rating = :r"),
	}

	// execyte update function
	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}
	fmt.Println("Successfully updated '" + movieName + "' (" + movieYear + ") rating to " + movieRating)
}
