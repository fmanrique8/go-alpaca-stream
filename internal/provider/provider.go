package provider

// Provider interface defines methods for market data providers.
type Provider interface {
	Connect() error
	Subscribe(channels []string) error
	HandleMessages() error
}
