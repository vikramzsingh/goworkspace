package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// creating DynamoDB client
	// Expose http req. and res.
	svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	// Add "CustomHeader" header with value of 10
	svc.Handlers.Send.PushFront(func(r *request.Request) {
		r.HTTPRequest.Header.Set("CustomHeader", fmt.Sprintf("%d", 10))
	})

	// Call ListTables just to see HTTP request/response
	// The request should have the CustomHeader set to 10
	_, _ = svc.ListTables(&dynamodb.ListTablesInput{})
}
