package auth

import (
	"io"
	"log"
	"net/http"
	"os"
)

// CreateRequest creates a new HTTP GET request with the necessary headers.
func CreateRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("APCA-API-KEY-ID", os.Getenv("API_KEY"))
	req.Header.Set("APCA-API-SECRET-KEY", os.Getenv("API_SECRET"))
	return req, nil
}

// ReadResponseBody reads the response body and returns it as a string.
func ReadResponseBody(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// CloseResponseBody closes the response body and logs any error.
func CloseResponseBody(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		log.Fatalf("Error closing response body: %v", err)
	}
}
