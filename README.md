
# Go Alpaca Stream

## Overview

Go Alpaca Stream is a Go application designed to connect to the Alpaca cryptocurrency market data stream. It retrieves real-time market data bars and prints the details to the console. This application is useful for financial analysts, traders, and developers interested in crypto market data.

## Prerequisites

- Go 1.16 or higher installed
- Alpaca API key and secret

## Quick Start

1. **Set up environment variables:**

   Create a `.env` file in the project directory with the following content:
   ```
   ALPACA_API_KEY=your_api_key
   ALPACA_API_SECRET=your_api_secret
   STREAM_URL=wss://stream.data.alpaca.markets/v2/iex
   CRYPTO_STREAM_URL=wss://stream.data.alpaca.markets/v1beta3/crypto/us
   ```

2. **Install dependencies:**

   Run the following command:
   ```sh
   go mod tidy
   ```

3. **Run the application:**

   Execute the following command:
   ```sh
   go run cmd/main.go
   ```

   This will connect to the Alpaca crypto market data stream and print the received market data bars to the console.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License.
