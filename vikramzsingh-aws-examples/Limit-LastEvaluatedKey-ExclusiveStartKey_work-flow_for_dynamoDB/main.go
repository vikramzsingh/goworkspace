// pagination demo
package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type OtaJob struct {
	Id              string      `json:"id"`
	Type            string      `json:"type"`
	PType           string      `json:"device_type"`
	SType           string      `json:"stype"`
	FirmwareVersion string      `json:"fw_version"`
	FirmwareUrl     string      `json:"fw_url"`
	FirmwareApiKey  string      `json:"fw_apikey"`
	Status          string      `json:"status"`
	Checksum        string      `json:"chksum"`
	CreatedOn       int64       `json:"created_on"`
	ModifiedOn      int64       `json:"modified_on"`
	CreatedBy       string      `json:"created_by"`
	ModifiedBy      string      `json:"modified_by"`
	CompletedOn     int64       `json:"completed_on"`
	Archived        int         `json:"archived"`
	CancelledBy     string      `json:"cancelled_by"`
	CancelledOn     int64       `json:"cancelled_on"`
	TargetSelection string      `json:"targetSelection"`
	TargetType      string      `json:"targetType"`
	TargetValue     []string    `json:"targetValue"`
	JobArn          string      `json:"job_arn"`
	ArchivedBy      string      `json:"archived_by"`
	ArchivedOn      int64       `json:"archived_on"`
	ProcessDetails  interface{} `json:"jobProcessDetails"`
}


type OtaList struct {
	Items []OtaJob `json:"items"`
	LeKey string          `json:"leKey"`
}

// type data interface{}

var result *dynamodb.QueryOutput

func main() {
	// Initilise session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))

	tableName := "ota_job"
	// movieYear := "2020"

	// create a DynamoDB service
	svc := dynamodb.New(sess)

	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		IndexName: aws.String("device_type-archived-index"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":device_type": {
				S: aws.String("IF31"),
			},
		},
		KeyConditionExpression: aws.String("device_type = :device_type"),
		ScanIndexForward:       aws.Bool(false),
		Limit:                  aws.Int64(20),
	
		// without IndexName
		// IndexName: aws.String("title-plot-index"),
		// ExpressionAttributeNames: map[string]*string{
		// 	"#year": aws.String("year"),
		// 	"#title": aws.String("title"),
		// },
		// ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
		// 	":year": {
		// 		N: aws.String("2018"),
		// 	},
		// 	":title": {
		// 		S: aws.String("Avengers"),
		// 	},
		// },
		// KeyConditionExpression: aws.String(" #title = :title "),
		// ScanIndexForward: aws.Bool(false), // default true means asc order, when false means decs order
		// Limit: aws.Int64(4),
		// ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
		// 	"archived": {
		// 		N: aws.String("0"),
		// 	},
		// 	"device_type": {
		// 		S: aws.String("IF31"),
		// 	},
		// 	"id": {
		// 		S: aws.String("1626951512021"),
		// 	},
		// },
	}

	input.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
		"archived": {
			N: aws.String("0"),
		},
		"device_type": {
			S: aws.String("IF31"),
		},
		"id": {
			S: aws.String("1626951512021"),
		},
	}

	fmt.Println("ExclusiveStart Key:", input.ExclusiveStartKey)


	result, err := svc.Query(input)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	items := []OtaJob{}
	for _, i := range result.Items {
		item := OtaJob{}
		err := dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			log.Println("Got error unmarshalling otaJob:", err.Error())
			continue
		}
		fmt.Println(item)
		items = append(items, item)
	}

	fmt.Println("Items", items)
	fmt.Println("LastEvaluatedKey:", result.LastEvaluatedKey)

	// for result.LastEvaluatedKey != nil {
	// 	fmt.Println("-----------------While Enter-------------")
	// 	// key := result.LastEvaluatedKey
	// 	// input.ExclusiveStartKey = key

	// 	result, err = svc.Query(input)
	// 	if err != nil {
	// 		log.Fatalf("Got error calling GetItem: %s", err)
	// 	}
	// 	items := []Item{}
	// 	for _, i := range result.Items {
	// 		item := Item{}
	// 		err := dynamodbattribute.UnmarshalMap(i, &item)
	// 		if err != nil {
	// 			log.Println("Got error unmarshalling otaJob:", err.Error())
	// 			continue
	// 		}
	// 		fmt.Println(item)
	// 		items = append(items, item)
	// 	}
	// 	fmt.Println("Items in while loop:", items)
	// 	fmt.Println("LastEvaluatedKey:", result.LastEvaluatedKey)
	// }
	// fmt.Println("----------------while exit---------------")
}
