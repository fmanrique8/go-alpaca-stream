package auth

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestLoadEnv(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	if os.Getenv("API_KEY") == "" {
		t.Error("Expected API_KEY to be set")
	}

	if os.Getenv("API_SECRET") == "" {
		t.Error("Expected API_SECRET to be set")
	}

	if os.Getenv("APISERVER_DOMAIN") == "" {
		t.Error("Expected APISERVER_DOMAIN to be set")
	}

	if os.Getenv("STREAM_URL") == "" {
		t.Error("Expected STREAM_URL to be set")
	}
}
