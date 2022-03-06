package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func filterMethods(methods []string) []string {
	filtered := make([]string, 0, len(methods))
	for _, m := range methods {
		v := strings.ToUpper(m)
		switch v {
		case "POST", "GET", "PUT", "PATCH", "DELETE":
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	bucketPtr := flag.String("b", "", "Bucket to set CORS, (required)")

	flag.Parse()

	if *bucketPtr == "" {
		exitErrorf("-b <bucket> Bucket name required")
	}

	methods := filterMethods(flag.Args())

	fmt.Println("Methods function:", methods)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))
	svc := s3.New(sess)

	rule := s3.CORSRule{
		AllowedHeaders: aws.StringSlice([]string{"Authorization"}),
		AllowedOrigins: aws.StringSlice([]string{"*"}),
		MaxAgeSeconds:  aws.Int64(3000),
		// Add HTTP methods CORS request that were specified in the CLI.
		AllowedMethods: aws.StringSlice(methods),
	}

	params := s3.PutBucketCorsInput{
		Bucket: bucketPtr,
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{&rule},
		},
	}

	_, err := svc.PutBucketCors(&params)
	if err != nil {
		// Print the error message
		exitErrorf("Unable to set Bucket %q's CORS, %v", *bucketPtr, err)
	}
	// Print the updated CORS config for the bucket
	fmt.Printf("Updated bucket %q CORS for %v\n", *bucketPtr, methods)
}
