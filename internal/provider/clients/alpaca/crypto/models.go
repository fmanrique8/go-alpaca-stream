package crypto

import (
	"encoding/json"
	"time"
)

// Bar represents a market data bar for crypto.
type Bar struct {
	Symbol    string    `json:"S"`
	Open      float64   `json:"o"`
	High      float64   `json:"h"`
	Low       float64   `json:"l"`
	Close     float64   `json:"c"`
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
	if val, ok := raw["t"].(string); ok {
		timestamp, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return err
		}
		b.Timestamp = timestamp
	}

	return nil
}
