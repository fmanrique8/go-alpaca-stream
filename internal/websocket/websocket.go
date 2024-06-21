package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"syscall"
)

func ConnectAndSubscribe() {
	url := "wss://stream.data.alpaca.markets/v2/iex"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Error connecting to WebSocket:", err)
		return
	}
	defer conn.Close()

	authMsg := map[string]string{
		"action": "auth",
		"key":    os.Getenv("API_KEY"),
		"secret": os.Getenv("API_SECRET"),
	}
	if err := conn.WriteJSON(authMsg); err != nil {
		fmt.Println("Error sending auth message:", err)
		return
	}

	subMsg := map[string]interface{}{
		"action": "subscribe",
		"bars":   []string{"AAPL"},
	}
	if err := conn.WriteJSON(subMsg); err != nil {
		fmt.Println("Error sending subscribe message:", err)
		return
	}

	fmt.Println("Establishing Connection...")

	go handleMessages(conn)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	fmt.Println("Received interrupt signal, closing connection...")
}
