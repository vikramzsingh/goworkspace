package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	// user need to pass bucket name form CLI in this program,
	// you also can hard-code name here
	if len(os.Args) != 2 {
		exitErrorf("Bucket name missing!\nUsage: %s bucket_name", os.Args[0])
	}

	bucket := os.Args[1]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))

	// Craete s3 service
	svc := s3.New(sess)

	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		// This field takes bucket-name that we are about to create
		Bucket: aws.String(bucket),
	})

	if err != nil {
		exitErrorf("Unable to create bucket %q, %v", bucket, err)
	}
	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucket)
	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be created, %v", bucket)
	}
	fmt.Printf("Bucket %q successfully created\n", bucket)
}
