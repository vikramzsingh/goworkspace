package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Output struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string  		  `json:"body"`
	IsBase64Encoded bool     	      `json:"isBase64Encoded"`
}

func Handler() (Output, error) {
	link := "https://www.pexels.com/photo/two-yellow-labrador-retriever-puppies-1108099/"
	img, _ := getImage(link)
	
	o := Output {
		StatusCode: 200,
		Headers: map[string]string{
			"content-type": "media/jpg",
		},
		Body: base64.StdEncoding.EncodeToString(img),
		IsBase64Encoded: true, // for files or images ture is required
	}

	return o, err
}

func getImage(url string) ([]bytes, error) {
	resp, err := http.Get(url)
	if err != nil {
		return bytes, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {
	lambda.Start(Handler)
}