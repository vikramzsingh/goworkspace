package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Year  int    `json:"year"`
	Title string `json:"title"`
}

func getItems() []Item {
	sliceOfbytes, err := ioutil.ReadFile("./movies_data.json")
	if err != nil {
		fmt.Println(err)
	}

	var items []Item
	_ = json.Unmarshal(sliceOfbytes, &items)
	return items
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// getItems from JSON file
	items := getItems()
	fmt.Println(items)

	tableName := "Movies"

	for _, item := range items {
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			fmt.Println(err)
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(tableName),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			fmt.Println(err)
		}

		year := strconv.Itoa(item.Year)
		fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)
	}

}
