package ota

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
)
// sess := session.Must(session.NewSessionWithOptions(session.Options{
// 	SharedConfigState: session.SharedConfigEnable,
// 	AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
// }))

var svc = iot.New(session.New())

var GetOTAUpdateFunc = GetOTAUpdate

func GetOTAUpdate(updateId string) (*iot.GetOTAUpdateOutput, error) {

	input := &iot.GetOTAUpdateInput{
		OtaUpdateId: aws.String(updateId),
	}
	return svc.GetOTAUpdate(input)
}