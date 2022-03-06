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

	if len(os.Args) != 3 {
		exitErrorf("Bucket and object name required\nUsage: %s bucket_name object_name",
			os.Args[0])
	}

	bucket := os.Args[1]
	obj := os.Args[2]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket), // bucket Name
		Key:    aws.String(obj),    // object Name or file Name to delete
	})
	if err != nil {
		exitErrorf("Unable to delete object %q from bucket %q, %v", obj, bucket, err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for object %q to be deleted, %v", obj, err)
	}

	fmt.Printf("Object %q successfully deleted\n", obj)
}
