/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long: `This command fetches a random dad joke form api`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type joke struct {
	ID string `json="id"`
	Joke string `json="joke"`
	Status int `json="status"`
}

func getRandomJoke() {
	url := "http://icanhazdadjoke.com/"
	responseByte := getJokeData(url)

	joke := joke{}
	//var joke joke

	if err := json.Unmarshal(responseByte, &joke); err != nil { // json to go
		log.Print("Could not unmarshal response - %v", err)
	}

	fmt.Println(string(responseByte))
	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest( // create a new request
		http.MethodGet,
		baseAPI,
		nil,
		)
	if err != nil {
		log.Printf("Could not request a dadjoke - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/examole/cobra-dadjoke)")

	response, err := http.DefaultClient.Do(request) // getting response data
	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body - %v", err)
	}

	return responseByte
}
