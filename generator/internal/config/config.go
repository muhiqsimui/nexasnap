package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CoinGeckoAPIKey string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		CoinGeckoAPIKey: os.Getenv("COINGECKO_API_KEY"),
	}, nil
}