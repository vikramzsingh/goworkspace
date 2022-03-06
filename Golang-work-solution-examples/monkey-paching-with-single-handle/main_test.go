package main

import (
	"log"
	"monkey-paching-with-single-handle/ota"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iot"
)

func TestMain(t *testing.T) {
	// Here cleanUp done.
	oldGetOTAUpdate := ota.GetOTAUpdateFunc
	defer func() { ota.GetOTAUpdateFunc = oldGetOTAUpdate }() // clearing request made on ota.GetOTAUpdateFunc
	log.Println("ns")
	ota.GetOTAUpdateFunc = func(updateId string) (*iot.GetOTAUpdateOutput, error) {
		log.Println("-----Success------")
		return &iot.GetOTAUpdateOutput{
			OtaUpdateInfo: &iot.OTAUpdateInfo{
				OtaUpdateFiles: []*iot.OTAUpdateFile{
					{
						Attributes: map[string]*string{
							"Ptype": aws.String("XTR"),
						},
					},
				},
			},
		}, nil
	}

	myFunc("bsjnaj")
}
