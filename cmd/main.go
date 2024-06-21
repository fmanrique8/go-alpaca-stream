package main

import (
	"go-alpaca-stream/internal/config"
	"go-alpaca-stream/internal/websocket"
)

func main() {
	config.LoadEnv()                // Load environment variables
	websocket.ConnectAndSubscribe() // Connect to WebSocket and subscribe to AAPL data
}
