package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "Movies"
	minRating := 4.0
	year := 2019

	// Create the Expression to fill the input struct with.
	// Get all movies in that year; we'll pull out those with a higher rating later
	filt := expression.Name("year").Equal(expression.Value(year))
	// Or we could get by ratings and pull out those with the right year later
	// filt := expression.Name("info.rating").GreaterThan(expression.Value(min_rating))
	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("title"), expression.Name("year"),
		expression.Name("rating"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	numItems := 0
	for _, i := range result.Items {
		item := Item{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}
		// Which ones had a higher rating than minimum?
		if item.Rating > minRating {
			// Or it we had filtered by rating previously:
			// if item.Year == year {
			numItems++
			fmt.Println("Title: ", item.Title)
			fmt.Println("Rating:", item.Rating)
			fmt.Println()
		}
	}
	fmt.Println("Found", numItems, "movie(s) with a rating above", minRating, "in", year)
}
