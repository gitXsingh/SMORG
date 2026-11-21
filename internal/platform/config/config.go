package config

import (
	"os"
	"time"
)

type Config struct {
	Address           string
	ReadHeaderTimeout time.Duration
}

func Load() Config {
	address := os.Getenv("APP_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	return Config{Address: address, ReadHeaderTimeout: 5 * time.Second}
}
