package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ConnectAndSubscribe() {
	url := "wss://stream.data.alpaca.markets/v2/iex"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer conn.Close()

	authMsg := map[string]string{
		"action": "auth",
		"key":    os.Getenv("API_KEY"),
		"secret": os.Getenv("API_SECRET"),
	}
	log.Println("Sending auth message:", authMsg)
	if err := conn.WriteJSON(authMsg); err != nil {
		log.Fatal("Error sending auth message:", err)
	}

	subMsg := map[string]interface{}{
		"action": "subscribe",
		"bars":   []string{"AAPL"},
	}
	log.Println("Sending subscribe message:", subMsg)
	if err := conn.WriteJSON(subMsg); err != nil {
		log.Fatal("Error sending subscribe message:", err)
	}

	go handleMessages(conn)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	fmt.Println("Received interrupt signal, closing connection...")
}
