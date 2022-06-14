/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// hitmeCmd represents the hitme command
var hitmeCmd = &cobra.Command{
	Use:   "hitme",
	Short: "Get a random dad joke",
	Long:  `Gets a random dad joke from https://icanhazdadjokes.com and prints it to the screen`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(hitmeCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com"
	responseBytes := getJokeData(url)

	recievedJoke := Joke{}
	if err := json.Unmarshal(responseBytes, &recievedJoke); err != nil {
		log.Printf("Failed to unmarshal response: %v", err)
	}

	fmt.Println(string(recievedJoke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)

	if err != nil {
		log.Printf("Could not request a dad joke: %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "dadjokeCLI Go tutorial - learning Go")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Failed to make request to API: %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body: %v", err)
	}

	return responseBytes
}
