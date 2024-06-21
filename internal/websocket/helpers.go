package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Bar struct {
	Symbol    string    `json:"S"`
	Open      float64   `json:"o"`
	Close     float64   `json:"c"`
	High      float64   `json:"h"`
	Low       float64   `json:"l"`
	Volume    int       `json:"v"`
	Timestamp time.Time `json:"t"`
}

func (b *Bar) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if val, ok := raw["S"].(string); ok {
		b.Symbol = val
	}
	if val, ok := raw["o"].(float64); ok {
		b.Open = val
	}
	if val, ok := raw["c"].(float64); ok {
		b.Close = val
	}
	if val, ok := raw["h"].(float64); ok {
		b.High = val
	}
	if val, ok := raw["l"].(float64); ok {
		b.Low = val
	}
	if val, ok := raw["v"].(float64); ok {
		b.Volume = int(val)
	}
	if val, ok := raw["t"].(string); ok {
		timestamp, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return err
		}
		b.Timestamp = timestamp
	}

	return nil
}

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
