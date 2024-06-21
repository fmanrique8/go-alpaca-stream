package main

import (
	"go-alpaca-stream/internal/auth"
	"go-alpaca-stream/internal/config"
	"go-alpaca-stream/internal/websocket"
)

func main() {
	config.LoadEnv()                // Load environment variables
	auth.GetAccount()               // Make a GET request to verify connection
	websocket.ConnectAndSubscribe() // Connect to WebSocket and subscribe to AAPL data
}
