package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	// create flags for emailID, userPool or userName
	emailIDPtr := flag.String("e", "", "The email address of the user")
	userPoolIDPtr := flag.String("p", "", "The ID of the user pool")
	userNamePtr := flag.String("n", "", "The name of the user")

	flag.Parse()


    if *emailIDPtr == "" || *userPoolIDPtr == "" || *userNamePtr == "" {
        fmt.Println("You must supply an email address, user pool ID, and user name")
        fmt.Println("Usage: go run CreateUser.go -e EMAIL-ADDRESS -p USER-POOL-ID -n USER-NAME")
        os.Exit(1)
    }
	
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))

	svc := cognitoidentityprovider.New(sess)

	newUserData := &cognitoidentityprovider.AdminCreateUserInput{
		DesiredDeliveryMediums: []*string{
			aws.String("EMAIL"),
		},
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name: aws.String("email"),
				Value: aws.String(*emailIDPtr),
			},
		},
		UserPoolId: aws.String(*userPoolIDPtr),
		Username: aws.String(*userNamePtr),
	}

	// Or you also use these functions
	// newUserData.SetUserPoolId(*userPoolIDPtr)
	// newUserData.SetUsername(*userNamePtr)

	_, err := svc.AdminCreateUser(newUserData) 
	if err != nil {
        fmt.Println("Got error creating user:", err)
    }

}