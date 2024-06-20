package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// GetAccount makes a GET request to the Alpaca API to verify connection.
func GetAccount() {
	url := "https://" + os.Getenv("APISERVER_DOMAIN") + "/v2/account"
	req, err := CreateRequest(url)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer CloseResponseBody(resp)

	body, err := ReadResponseBody(resp)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", body)
}
