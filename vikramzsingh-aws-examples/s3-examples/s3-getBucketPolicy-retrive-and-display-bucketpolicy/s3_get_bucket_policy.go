package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("bucket name required", filepath.Base(os.Args[0]))
	}
	bucket := os.Args[1]

	//load credentails
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
		SharedConfigState:       session.SharedConfigEnable,
	}))

	// create s3 service
	svc := s3.New(sess)

	result, err := svc.GetBucketPolicy(&s3.GetBucketPolicyInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		fmt.Println("GetBucketpolicy error!!!", err)
		return
	}

	out := bytes.Buffer{}
	policyStr := aws.StringValue(result.Policy)
	if err := json.Indent(&out, []byte(policyStr), "", " "); err != nil {
		fmt.Println("Intent error!!", err)
		return
	}

	fmt.Println(out.String())
}

// To run Program:-
// go run s3_get_bucket_policy.go <BUCKET_NAME>

// THIS PROGRAM TO DEMONSTRATE HOW TO GET BUCKET POLICY
// There are other programs like PUT-BUCKET-POLICY, DELETE-BUCKET-POLICY, TYR THEM>>>> 