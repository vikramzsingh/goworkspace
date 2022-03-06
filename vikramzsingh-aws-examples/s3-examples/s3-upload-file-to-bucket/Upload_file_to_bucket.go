package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	// User need to pass bucket name from CLI,
	// OR you can hard-code here
	if len(os.Args) != 3 {
		exitErrorf("bucket and file name required\nUsage: %s bucket_name filename", os.Args[0])
	}

	bucket := os.Args[1]
	filename := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}
	defer file.Close()

	// initilize session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	// Create s3manager service to upload file 
	uploader := s3manager.NewUploader(sess)

	// s3manager service provide method to your file in AWS s3  
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket), // This field is Bucket Name
		Key:    aws.String(filename), // This field is your file-name that you want to UPLOAD
		Body:   file, // This field is of type == io.Reader, it takes readable content of the file, this can read from img or text files
	})

	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}
	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}
