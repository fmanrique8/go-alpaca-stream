package main

import (
	"go-alpaca-stream/internal"
	"go-alpaca-stream/internal/provider"
	"go-alpaca-stream/internal/provider/clients/alpaca/crypto"
	"log"
	"sync"
)

func main() {
	internal.LoadEnv()

	var wg sync.WaitGroup
	wg.Add(1)

	// Initialize the crypto clients as a provider
	var cryptoProvider provider.Provider = &crypto.Client{}

	// Connect to the provider
	err := cryptoProvider.Connect()
	if err != nil {
		log.Fatalf("Provider connection error: %v", err)
	}

	// Subscribe to channels
	err = cryptoProvider.Subscribe([]string{"BTC/USD"})
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
