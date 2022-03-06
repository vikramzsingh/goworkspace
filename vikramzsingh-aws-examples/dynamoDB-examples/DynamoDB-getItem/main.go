package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	Year  int    `json:"year"`
	Title string `json:"title"`
}

type UserNotifyPref struct {
	Id                    string          `json:"id"`
	UserId                string          `json:"userId"`
	Pref                  map[string]bool `json:"pref"`
	LangPref              string          `json:"langPref"`
	Ptype                 string          `json:"ptype"`
	Stype                 string          `json:"stype"`
	Ntype                 string          `json:"ntype"`
	SMSVerificationStatus int             `json:"smsVerificationStatus"`
}


// type data interface{}

func main() {
	// Initilise session 
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))

	// create a DynamoDB service
	svc := dynamodb.New(sess)

	// npref, err := GetByPrimaryKey("us-west-2_259lyYUwA#40ba33ca-3ca1-4422-97fd-466b92a280a9#news", svc)
	// if err != nil {
	// 	fmt.Println("A")
	// 	fmt.Println(err)
	// }

	// fmt.Println("result:=", npref)

	// getDeviceNotifyTemplate(npref)

	
	// tableName := "Movies"
	// movieYear := "2020"

	// result, err := svc.GetItem(&dynamodb.GetItemInput{
	params := &dynamodb.ScanInput{
			TableName: aws.String("product_config"),
	}

	result, err := svc.Scan(params)
	if err != nil {
	log.Fatalf("Query API call failed: %s", err)
	}

	fmt.Println("RESULT:-", result)
	

	// if err != nil {
	// 	log.Fatalf("Got error calling GetItem: %s", err)
	// }

	// if result.Item == nil {
	// 	msg := "Could not find"
	// 	return nil, errors.New(msg)
	// }

	// Output form result (FIELD Item) is of type == map[string]*dynamodb.AttributeValue,
	// with the below UnmarshalMap we converted this into type == Item
	// item := Item{}
	// err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	// }

	// fmt.Println("Found item:")
	// fmt.Println("Year: ", item.Year)
	// fmt.Println("Title: ", item.Title)
}


// example code:-

// func GetByPrimarykey(alarmId string) *entity.Alarm {

// 	input := &dynamodb.QueryInput{
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":aid": {
// 				S: aws.String(alarmId),
// 			},
// 		},
// 		KeyConditionExpression: aws.String("id = :aid"),
// 		TableName:              aws.String(TABLE_NAME),
// 	}
// 	result, err := dynamo.ExecuteQuery(input)
// 	if err != nil {
// 		return nil
// 	}
// 	for _, i := range result.Items {
// 		item := entity.Alarm{}
// 		err = dynamodbattribute.UnmarshalMap(i, &item)
// 		if err != nil {
// 			log.Println("Got error GetByPrimarykey for Alarm:", err.Error())
// 			continue
// 		}
// 		return &item
// 	}
// 	return nil
// }

// func GetByPrimaryKey(id string, svc *dynamodb.DynamoDB) (*UserNotifyPref, error) {
// 	input := &dynamodb.QueryInput{
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":u1": {
// 				S: aws.String(id),
// 			},
// 		},
// 		KeyConditionExpression: aws.String("id = :u1"),
// 		TableName:              aws.String("user_notify_pref"),
// 	}

// 	result, err := svc.Query(input)
// 	fmt.Println("result----------------", result)
// 	if err != nil {
// 		fmt.Println("-----------------------qytquwy--------------------")
// 		return nil, err
// 	}
// 	for _, i := range result.Items {
// 		item := UserNotifyPref{}
// 		fmt.Println("Unmarshal check")
// 		err := dynamodbattribute.UnmarshalMap(i, &item)
// 		if err != nil {
// 			log.Println("Got error unmarshalling GetByPrimaryKey:", err.Error())
// 			continue
// 		}
// 		return &item, nil
// 	}

// 	fmt.Println("end func")
// 	return nil, nil	
// }


// func getDeviceNotifyTemplate(pref *UserNotifyPref) {
// 	fmt.Println("Pref", pref)
// }
