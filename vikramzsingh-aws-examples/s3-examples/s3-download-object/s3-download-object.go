package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		exitErrorf("Bucket and item names required\nUsage: %s bucket_name item_name", os.Args[0])
	}

	bucket := os.Args[1]
	item := os.Args[2] // Object name to download

	// This Creates file for you
	file, err := os.Create(item)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", item, err)
	}

	// initilize session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	// create s3manager service to download file from AWS s3
	downloader := s3manager.NewDownloader(sess)

	// this method downloads file from AWS and write it to file (io.File),
	// returns number of bytes it has downloaded and error 
	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket), // bucket name from where file needs to download
		Key:    aws.String(item),   // key is Just file name that need to be download. 
	})

	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
