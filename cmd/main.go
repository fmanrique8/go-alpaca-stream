package main

import (
	"go-alpaca-stream/internal/auth"
	"go-alpaca-stream/internal/config"
)

func main() {
	config.LoadEnv()  // Load environment variables
	auth.GetAccount() // Make a GET request to verify connection
}
