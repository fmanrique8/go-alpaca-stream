package provider

import "time"

// Provider interface defines methods for market data providers.
type Provider interface {
	Connect() error
	Subscribe(channels []string) error
	HandleMessages() error
}

// Bar represents a market data bar.
type Bar struct {
	Symbol    string    `json:"S"`
	Open      float64   `json:"o"`
	High      float64   `json:"h"`
	Low       float64   `json:"l"`
	Close     float64   `json:"c"`
	Timestamp time.Time `json:"t"`
}
