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
	if len(os.Args) != 2 {
		exitErrorf("Bucket name required\nUsage: %s bucket_name", os.Args[0])
	}

	bucket := os.Args[1]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}
	for _, item := range resp.Contents {
		fmt.Println("Name: ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size: ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}
