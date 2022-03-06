package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	bucket := "demo-serverless-service-serverlessdeploymentbuck-16xxmtgvcx133"
	tagName1 := "Cost Center"
	tagValue1 := "123456"
	tagName2 := "Stack"
	tagValue2 := "MyTestStack"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// s3 bucket service Instance
	svc := s3.New(sess)

	// Tagging s3 bucket with PutBucket Method
	putInput := &s3.PutBucketTaggingInput{
		Bucket: aws.String(bucket),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(tagName1),
					Value: aws.String(tagValue1),
				},
				{
					Key:   aws.String(tagName2),
					Value: aws.String(tagValue2),
				},
			},
		},
	}

	_, err = svc.PutBucketTagging(putInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// getting s3 bucket Tag aith GetBucket Method
	getInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucket),
	}

	result, err := svc.GetBucketTagging(getInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numTags := len(result.TagSet)
	if numTags > 0 {
		fmt.Println("Found", numTags, "Tag(s):")
		fmt.Println("")
		for _, t := range result.TagSet {
			fmt.Println(" Key: ", *t.Key)
			fmt.Println(" Value:", *t.Value)
			fmt.Println("")
		}
	} else {
		fmt.Println("Did not find any tags")
	}
}
