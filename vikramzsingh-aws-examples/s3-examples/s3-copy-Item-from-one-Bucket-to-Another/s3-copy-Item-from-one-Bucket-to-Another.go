package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

func main() {

	if len(os.Args) != 4 {
		exitErrorf("Bucket, item, and other bucket names required\nUsage: go run s3_copy_object bucket item other-bucket")
	}

	bucket := os.Args[1]
	item := os.Args[2]
	other := os.Args[3]

	source := bucket + "/" + item
	fmt.Println(source)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	// Copy the item
	_, err := svc.CopyObject(&s3.CopyObjectInput{
		Bucket: aws.String(other), // Name od destination bucket
		CopySource: aws.String(url.PathEscape(source)), //
		Key: aws.String(item),
	})
	if err != nil {
		exitErrorf("Unable to copy item from bucket %q to bucket %q, %v", bucket, other, err)
	}

	fmt.Printf("Item %q successfully copied from bucket %q to bucket %q\n", item, bucket, other)
}
