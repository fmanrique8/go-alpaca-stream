package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getAccount() {
	// Make a GET request to the Alpaca API to verify connection
	url := "https://" + os.Getenv("APISERVER_DOMAIN") + "/v2/account"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers with API keys
	req.Header.Set("APCA-API-KEY-ID", os.Getenv("API_KEY"))
	req.Header.Set("APCA-API-SECRET-KEY", os.Getenv("API_SECRET"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print response
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
}

func main() {
	loadEnv()    // Load environment variables
	getAccount() // Make a GET request to verify connection
}
