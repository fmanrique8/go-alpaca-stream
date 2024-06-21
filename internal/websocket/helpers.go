package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func PrintBar(bar Bar) {
	if bar.Symbol == "" && bar.Open == 0 && bar.Close == 0 && bar.High == 0 && bar.Low == 0 && bar.Volume == 0 {
		return // Skip empty records
	}
	fmt.Printf("Symbol: %s | Open: %.2f | High: %.2f | Low: %.2f | Close: %.2f | Volume: %d | Timestamp (UTC): %s\n",
		bar.Symbol, bar.Open, bar.High, bar.Low, bar.Close, bar.Volume, bar.Timestamp.Format("2006-01-02 15:04:05"))
}

func handleMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		var bars []Bar
		if err := json.Unmarshal(message, &bars); err != nil {
			fmt.Println("Error unmarshalling message:", err)
			return
		}
		for _, bar := range bars {
			PrintBar(bar)
		}
	}
}
