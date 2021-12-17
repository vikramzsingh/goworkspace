package main

import (
	"encoding/json"
	"fmt"
	"log"

)

type thumbnail struct {
	URL    string `json:"Url"`
	Height int    `json:"Height"`
	Width  int    `json:"Width"`
}

type img struct {
	Width     int       `json:"Width"`
	Height    int       `json:"Height"`
	Title     string    `json:"Title"`
	Thumbnail thumbnail `json:"Thumbnail"`
	Animated  bool      `json:"Animated"`
	IDs       []int     `json:"IDs"`
}

func main() {
	var data img

	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	// JSON to Go
	err := json.Unmarshal([]byte(rcvd), &data) // return error
	if err != nil {
		log.Println(err)
	}

	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	fmt.Println(data.Thumbnail.URL)
}
