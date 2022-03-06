package main

import (
	"errors"
	// "fmt"
	"log"
	// "mywater-backend-utility-lib/lib/go/dynamouserdao"
	// "os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	// "flag"
	// "fmt"
	// "os"
)

type UserPrefInfo struct {
	Email         string
	Phone         string
	Ios           string
	Android       string
	FName         string
	Lname         string
	EmailVerified bool
	NickName      string
	SDNickName    string
	LangPref      string
	TimeZone      string
}


func main() {
    // userPoolIDPtr := flag.String("p", "", "The ID of the user pool")

    // flag.Parse()

    // if *userPoolIDPtr == "" {
    //     fmt.Println("You must supply a user pool ID")
    //     fmt.Println("Usage: go run CreateUser.go -p USER-POOL-ID")
    //     os.Exit(1)
    // }

    // Initialize a session that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
    }))

    sci := cognitoidentityprovider.New(sess)

    s := getUserLanguage("1630852224123#general", sci)

    log.Println(s)


    // GetAdminGetUser("1630852224123#general", sci)

    // results, err := cognitoClient.ListUsers(
    //     &cognitoidentityprovider.ListUsersInput {
    //         UserPoolId: userPoolIDPtr,
	// 	},
	// )
    // if err != nil {
    //     fmt.Println("Got error listing users")
    //     os.Exit(1)
    // }

    // // Show their names an email addresses
    // for _, user := range results.Users {
    //     attributes := user.Attributes

    //     for _, a := range attributes {
    //         if *a.Name == "name" {
    //             fmt.Println("Name:  " + *a.Value)
    //         } else if *a.Name == "email" {
    //             fmt.Println("Email: " + *a.Value)
    //         }
    //     }

    //     fmt.Println("")
    // }
}


func GetAdminGetUser(userId string, sci *cognitoidentityprovider.CognitoIdentityProvider) (*UserPrefInfo, error) {

	s := strings.Split(userId, "#")
	if s == nil || len(s) != 2 {
		return nil, errors.New("user id split issue occurred")
	}
	params := &cognitoidentityprovider.AdminGetUserInput{
		Username:   aws.String(s[1]),
		UserPoolId: aws.String(s[0]),
	}

	res, err := sci.AdminGetUser(params)
    // fmt.Println("Results:", res)
	if err != nil {
        // log.Println("Error", err)
		if aerr, ok := err.(awserr.Error); ok {
            // fmt.Println("OK error:",aerr)
            // fmt.Println("-----------", aerr.Code())
			switch aerr.Code() {
			case cognitoidentityprovider.ErrCodeResourceNotFoundException:
				log.Println(cognitoidentityprovider.ErrCodeResourceNotFoundException, aerr.Error())
			case cognitoidentityprovider.ErrCodeNotAuthorizedException:
				log.Println(cognitoidentityprovider.ErrCodeNotAuthorizedException, aerr.Error())
			case cognitoidentityprovider.ErrCodeUserNotFoundException:
				log.Println(cognitoidentityprovider.ErrCodeUserNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				log.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
		return nil, err
	}
	userInfo := UserPrefInfo{}
	for _, attr := range res.UserAttributes {
		val := *attr.Value
		switch *attr.Name {
		case "custom:iosTokens":
			userInfo.Ios = val
		case "custom:androidTokens":
			userInfo.Android = val
		case "phone_number":
			userInfo.Phone = val
		case "email":
			userInfo.Email = val
		case "given_name":
			userInfo.FName = val
		case "name":
			userInfo.Lname = val
		case "custom:timezone":
			userInfo.TimeZone = val
		case "email_verified":
			emailverified, err := strconv.ParseBool(val)
			if err == nil {
				userInfo.EmailVerified = emailverified
			}
		case "locale":
			userInfo.LangPref = val
		default:
		}
	}
	return &userInfo, nil
}

func getUserLanguage(userId string, sci *cognitoidentityprovider.CognitoIdentityProvider) string {
	if len(userId) > 0 {
		userDetails, _ := GetAdminGetUser(userId, sci)
        // if err != nil {
        //     fmt.Println("Error occured here:", err)
        // }
		log.Println("userDetails are ", userDetails)
		if userDetails != nil && len(userDetails.LangPref) > 0 {
			return userDetails.LangPref
		}
	}
	return "en_us"
}
