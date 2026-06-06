package main

import (
	"fmt"
	"log"
	"time"

	"github.com/muhiqsimui/nexasnap/generator/internal/config"
	"github.com/muhiqsimui/nexasnap/generator/internal/docs"
	"github.com/muhiqsimui/nexasnap/generator/internal/fetcher"
	"github.com/muhiqsimui/nexasnap/generator/internal/models"
	"github.com/muhiqsimui/nexasnap/generator/internal/scraper"
	"github.com/muhiqsimui/nexasnap/generator/internal/writer"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	client := fetcher.NewClient(cfg.CoinGeckoAPIKey)
	now := time.Now().UTC()

	featuredCoins := config.GetCoins()
	coinIDs := make([]string, len(featuredCoins))
	for i, c := range featuredCoins {
		coinIDs[i] = c.CoinGeckoID
	}

	markets, err := client.FetchMarkets()
	if err != nil {
		log.Fatalf("markets: %v", err)
	}
	log.Printf("fetched %d markets", len(markets))

	trending, err := client.FetchTrending()
	if err != nil {
		log.Printf("warning: trending fetch failed: %v", err)
		trending = []models.CoinData{}
	}
	log.Printf("fetched %d trending", len(trending))

	global, err := client.FetchGlobal()
	if err != nil {
		log.Printf("warning: global fetch failed: %v", err)
		global = &models.GlobalData{}
	}
	log.Printf("fetched global data")

	marketMap := make(map[string]models.CoinData, len(markets))
	for _, m := range markets {
		marketMap[m.ID] = m
	}

	for _, coin := range featuredCoins {
		data, ok := marketMap[coin.CoinGeckoID]
		if !ok {
			log.Printf("warning: no market data for %s, skipping", coin.CoinGeckoID)
			continue
		}

		apiPath := fmt.Sprintf("../generated/api/v1/crypto/%s.json", coin.Slug)
		err = writer.WriteJSON(apiPath, models.APIResponse{
			Success: true,
			Source:  "CoinGecko",
			Updated: now,
			Data:    data,
		})
		if err != nil {
			log.Fatalf("write %s: %v", coin.Slug, err)
		}

		err = docs.WriteCoinDocs(coin, fmt.Sprintf("/api/v1/crypto/%s.json", coin.Slug))
		if err != nil {
			log.Fatalf("doc %s: %v", coin.Slug, err)
		}

		log.Printf("wrote %s with price $%.2f", coin.Name, data.CurrentPrice)
	}

	err = writer.WriteJSON("../generated/api/v1/markets.json", models.APIResponse{
		Success: true,
		Source:  "CoinGecko",
		Updated: now,
		Data:    markets,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := docs.WriteMarketsDocs(); err != nil {
		log.Fatal(err)
	}
	log.Println("wrote markets endpoint")

	err = writer.WriteJSON("../generated/api/v1/trending.json", models.APIResponse{
		Success: true,
		Source:  "CoinGecko",
		Updated: now,
		Data:    trending,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := docs.WriteTrendingDocs(); err != nil {
		log.Fatal(err)
	}
	log.Println("wrote trending endpoint")

	err = writer.WriteJSON("../generated/api/v1/global.json", models.APIResponse{
		Success: true,
		Source:  "CoinGecko",
		Updated: now,
		Data:    global,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := docs.WriteGlobalDocs(); err != nil {
		log.Fatal(err)
	}
	log.Println("wrote global endpoint")

	newsResult, err := scraper.ScrapeCryptoNews()
	if err != nil {
		log.Printf("warning: scraper failed: %v", err)
	} else {
		err = writer.WriteJSON("../generated/api/v1/news.json", models.APIResponse{
			Success: true,
			Source:  "CoinTelegraph",
			Updated: now,
			Data:    newsResult,
		})
		if err != nil {
			log.Fatal(err)
		}

		err = writer.WriteJSON("../generated/docs/news.json", models.EndpointDoc{
			Title:       "Crypto News Headlines",
			Slug:        "news",
			Category:    "news",
			Endpoint:    "/api/v1/news.json",
			Method:      "GET",
			Description: "Latest cryptocurrency news headlines from CoinTelegraph",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("wrote news endpoint")
	}

	log.Println("generation complete")
}
