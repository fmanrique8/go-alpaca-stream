package crypto

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go-alpaca-stream/internal/provider"
	"os"
)

// Client represents an Alpaca crypto client.
type Client struct {
	Conn *websocket.Conn
}

// Connect establishes a connection to the Alpaca crypto WebSocket.
func (c *Client) Connect() error {
	url := os.Getenv("CRYPTO_STREAM_URL")
	if url == "" {
		return fmt.Errorf("CRYPTO_STREAM_URL not set in environment variables")
	}

	var err error
	c.Conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("error connecting to WebSocket: %v", err)
	}

	return nil
}

// Subscribe subscribes to the given channels.
func (c *Client) Subscribe(channels []string) error {
	// Authenticate with a message
	authMsg := map[string]string{
		"action": "auth",
		"key":    os.Getenv("API_KEY"),
		"secret": os.Getenv("API_SECRET"),
	}

	if err := c.Conn.WriteJSON(authMsg); err != nil {
		return fmt.Errorf("error sending auth message: %v", err)
	}

	// Subscribe to the channels
	subMsg := map[string]interface{}{
		"action": "subscribe",
		"bars":   channels,
	}

	if err := c.Conn.WriteJSON(subMsg); err != nil {
		return fmt.Errorf("error sending subscribe message: %v", err)
	}

	return nil
}

// HandleMessages handles incoming messages from the WebSocket.
func (c *Client) HandleMessages() error {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("error reading message: %v", err)
		}

		// Print the raw message
		fmt.Printf("Received raw message: %s\n", message)
	}
}

var _ provider.Provider = (*Client)(nil)
