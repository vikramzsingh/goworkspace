package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// "log"

	// "os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// func exitErrorf(msg string, args ...interface{}) {
// 	fmt.Fprintf(os.Stderr, msg+"\n", args...)
// 	os.Exit(1)
// }

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Bucket-Name and Key or object-Name is required")
	}

	bucket := os.Args[1]
	object := os.Args[2]
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-1"),
	// })

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	// result is a *s3.GetObjectOutput struct pointer
	// err is a error which can be cast to awserr.Error.
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// process SDK error
			fmt.Println("erty", awsErr.Error())
		}
	}

	sliceOfByte, _ := ioutil.ReadAll(result.Body)
	err = ioutil.WriteFile(object, sliceOfByte, 0644)
	if err != nil {
		panic(err)
	}

	// or

	// if err != nil {
	// 	// Casting to the awserr.Error type will allow you to inspect the error
	// 	// code returned by the service in code. The error code can be used
	// 	// to switch on context specific functionality. In this case a context
	// 	// specific error message is printed to the user based on the bucket
	// 	// and key existing.
	// 	//
	// 	// For information on other S3 API error codes see:
	// 	// http://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
	// 	if aerr, ok := err.(awserr.Error); ok {
	// 		switch aerr.Code() {
	// 		case s3.ErrCodeNoSuchBucket:
	// 			exitErrorf("bucket %s does not exist", os.Args[1])
	// 		case s3.ErrCodeNoSuchKey:
	// 			exitErrorf("object with key %s does not exist in bucket %s", os.Args[2], os.Args[1])
	// 		}
	// 	}
	// }

	fmt.Println(result)
}
