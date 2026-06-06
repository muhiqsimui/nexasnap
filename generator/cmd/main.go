package main

import (
	"log"
	"time"

	"github.com/muhiqsimui/nexasnap/generator/internal/config"
	"github.com/muhiqsimui/nexasnap/generator/internal/fetcher"
	"github.com/muhiqsimui/nexasnap/generator/internal/writer"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("apikey:", cfg.CoinGeckoAPIKey)

	prices, err := fetcher.FetchPrices()
	if err != nil {
		log.Fatal(err)
	}

	err = writer.WriteJSON(
		"../generated/api/v1/crypto/bitcoin.json",
		map[string]any{
			"success": true,
			"source":  "CoinGecko",
			"updated": time.Now().UTC(),
			"data":    prices["bitcoin"],
		},
	)

	err = writer.WriteJSON(
		"../generated/docs/bitcoin.json",
		map[string]any{
			"title":       "Bitcoin Price",
			"slug":        "bitcoin",
			"category":    "crypto",
			"endpoint":    "/api/v1/crypto/bitcoin.json",
			"method":      "GET",
			"description": "Current Bitcoin price in USD",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	err = writer.WriteJSON(
		"../generated/api/v1/crypto/ethereum.json",
		map[string]any{
			"success": true,
			"source":  "CoinGecko",
			"updated": time.Now().UTC(),
			"data":    prices["ethereum"],
		},
	)

	err = writer.WriteJSON(
		"../generated/docs/ethereum.json",
		map[string]any{
			"title":       "Ethereum Price",
			"slug":        "ethereum",
			"category":    "crypto",
			"endpoint":    "/api/v1/crypto/ethereum.json",
			"method":      "GET",
			"description": "Current Ethereum price in USD",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("generation complete")
}