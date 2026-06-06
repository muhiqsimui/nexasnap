package docs

import (
	"fmt"

	"github.com/muhiqsimui/nexasnap/generator/internal/config"
	"github.com/muhiqsimui/nexasnap/generator/internal/models"
	"github.com/muhiqsimui/nexasnap/generator/internal/writer"
)

func WriteCoinDocs(coin config.CoinDef, endpoint string) error {
	doc := models.EndpointDoc{
		Title:       fmt.Sprintf("%s (%s) Price", coin.Name, coin.Symbol),
		Slug:        coin.Slug,
		Category:    coin.Category,
		Endpoint:    endpoint,
		Method:      "GET",
		Description: fmt.Sprintf("Current %s price and market data in USD", coin.Name),
	}

	path := fmt.Sprintf("../generated/docs/%s.json", coin.Slug)
	return writer.WriteJSON(path, doc)
}

func WriteGlobalDocs() error {
	doc := models.EndpointDoc{
		Title:       "Global Cryptocurrency Market Data",
		Slug:        "global",
		Category:    "market",
		Endpoint:    "/api/v1/global.json",
		Method:      "GET",
		Description: "Global cryptocurrency market overview including total market cap, volume, and BTC/ETH dominance",
	}

	return writer.WriteJSON("../generated/docs/global.json", doc)
}

func WriteMarketsDocs() error {
	doc := models.EndpointDoc{
		Title:       "Cryptocurrency Markets",
		Slug:        "markets",
		Category:    "market",
		Endpoint:    "/api/v1/markets.json",
		Method:      "GET",
		Description: "Top 100 cryptocurrencies by market cap with price, volume, and market data",
	}

	return writer.WriteJSON("../generated/docs/markets.json", doc)
}

func WriteTrendingDocs() error {
	doc := models.EndpointDoc{
		Title:       "Trending Coins",
		Slug:        "trending",
		Category:    "market",
		Endpoint:    "/api/v1/trending.json",
		Method:      "GET",
		Description: "Currently trending coins on CoinGecko",
	}

	return writer.WriteJSON("../generated/docs/trending.json", doc)
}
