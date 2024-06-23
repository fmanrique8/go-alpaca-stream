package main

import (
	"go-alpaca-stream/internal/config"
	"go-alpaca-stream/internal/provider"
	"go-alpaca-stream/internal/provider/clients/alpaca/crypto"
	"log"
	"sync"
)

func main() {
	config.LoadEnv()

	var wg sync.WaitGroup
	wg.Add(1)

	// Initialize the crypto client as a provider
	var cryptoProvider provider.Provider = &crypto.Client{}

	// Connect to the provider
	err := cryptoProvider.Connect()
	if err != nil {
		log.Fatalf("Provider connection error: %v", err)
	}

	// Subscribe to multiple channels
	cryptos := []string{"BTC/USD", "ETH/USD", "LTC/USD"}
	err = cryptoProvider.Subscribe(cryptos)
	if err != nil {
		log.Fatalf("Provider subscription error: %v", err)
	}

	// Handle messages
	go func() {
		defer wg.Done()
		err = cryptoProvider.HandleMessages()
		if err != nil {
			log.Fatalf("Provider handle messages error: %v", err)
		}
	}()

	// Wait for the message handling goroutine to finish
	wg.Wait()
}
