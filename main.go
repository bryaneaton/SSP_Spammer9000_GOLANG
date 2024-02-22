package main

import (
	"SSP_Spammer9000_GOLANG/src"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// main function is the entry point of the program
func main() {
	// Check if the number of arguments is not equal to 2
	if len(os.Args) != 2 {
		fmt.Println("Please provide a number.")
		os.Exit(1)
	}
	println("Creating " + os.Args[1] + " Random Security Plan for " + os.Getenv("REGSCALE_DOMAIN"))

	// Convert the argument to an integer
	argument, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid number.")
		os.Exit(1)
	}

	// Loop for the number of times specified by the argument
	for i := 0; i < argument; i++ {
		// Create a new SSP with a random word as the system name
		ssp := src.NewSSP(getRandomWord())
		// Post the SSP to the specified URL
		postSSP(ssp)
	}

}

// postSSP function sends a POST request to the specified URL with the SSP as the body
func postSSP(ssp *src.SSP) error {
	// Get the URL from the system environment
	url := os.Getenv("REGSCALE_DOMAIN")
	// Marshal the SSP into JSON
	jsonData, err := json.Marshal(ssp)
	if err != nil {
		return err
	}
	// Create a new POST request with the JSON as the body
	req, err := http.NewRequest("POST", url+"/api/securityplans", bytes.NewBuffer(jsonData))
	// Set the headers for the request
	req.Header.Set("Authorization", os.Getenv("REGSCALE_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	// Close the request body
	defer req.Body.Close()

	// Check if the response status is not OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post SSP, status code: %d", resp.StatusCode)
	}

	return nil
}

// getRandomWord function fetches a random word from an API
func getRandomWord() string {
	// Send a GET request to the API
	resp, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		log.Fatal(err)
	}
	// Close the response body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the response body into a slice of strings
	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the slice is not empty
	if len(words) > 0 {
		return words[0]
	}

	return ""
}
