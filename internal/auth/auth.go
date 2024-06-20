package auth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// GetAccount makes a GET request to the Alpaca API to verify connection.
func GetAccount() {
	url := "https://" + os.Getenv("APISERVER_DOMAIN") + "/v2/account"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("APCA-API-KEY-ID", os.Getenv("API_KEY"))
	req.Header.Set("APCA-API-SECRET-KEY", os.Getenv("API_SECRET"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("Error closing response body: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
}
