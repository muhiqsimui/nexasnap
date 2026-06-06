package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/muhiqsimui/nexasnap/generator/internal/models"
)

const baseURL = "https://api.coingecko.com/api/v3"

type Client struct {
	apiKey string
	http   *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http:   &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) do(req *http.Request, dst any) error {
	req.Header.Set("User-Agent", "NexaSnap/1.0")
	if c.apiKey != "" {
		req.Header.Set("x-cg-demo-api-key", c.apiKey)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(dst)
}

func (c *Client) get(path string, dst any) error {
	req, err := http.NewRequest(http.MethodGet, baseURL+path, nil)
	if err != nil {
		return err
	}
	return c.do(req, dst)
}

func (c *Client) FetchMarkets() ([]models.CoinData, error) {
	var raw []struct {
		ID                       string  `json:"id"`
		Symbol                   string  `json:"symbol"`
		Name                     string  `json:"name"`
		Image                    string  `json:"image"`
		CurrentPrice             float64 `json:"current_price"`
		MarketCap                float64 `json:"market_cap"`
		MarketCapRank            int     `json:"market_cap_rank"`
		TotalVolume              float64 `json:"total_volume"`
		High24h                  float64 `json:"high_24h"`
		Low24h                   float64 `json:"low_24h"`
		PriceChange24h           float64 `json:"price_change_24h"`
		PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
		CirculatingSupply        float64 `json:"circulating_supply"`
		TotalSupply              float64 `json:"total_supply"`
		MaxSupply                float64 `json:"max_supply"`
		Ath                      float64 `json:"ath"`
		AthDate                  string  `json:"ath_date"`
		Atl                      float64 `json:"atl"`
		AtlDate                  string  `json:"atl_date"`
		LastUpdated              string  `json:"last_updated"`
	}

	err := c.get("/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=24h", &raw)
	if err != nil {
		return nil, fmt.Errorf("fetch markets: %w", err)
	}

	coins := make([]models.CoinData, len(raw))
	for i, r := range raw {
		coins[i] = models.CoinData{
			ID:                       r.ID,
			Symbol:                   r.Symbol,
			Name:                     r.Name,
			Image:                    r.Image,
			CurrentPrice:             r.CurrentPrice,
			MarketCap:                r.MarketCap,
			MarketCapRank:            r.MarketCapRank,
			TotalVolume:              r.TotalVolume,
			High24h:                  r.High24h,
			Low24h:                   r.Low24h,
			PriceChange24h:           r.PriceChange24h,
			PriceChangePercentage24h: r.PriceChangePercentage24h,
			CirculatingSupply:        r.CirculatingSupply,
			TotalSupply:              r.TotalSupply,
			MaxSupply:                r.MaxSupply,
			Ath:                      r.Ath,
			AthDate:                  r.AthDate,
			Atl:                      r.Atl,
			AtlDate:                  r.AtlDate,
			LastUpdated:              r.LastUpdated,
		}
	}
	return coins, nil
}

func (c *Client) FetchTrending() ([]models.CoinData, error) {
	var raw struct {
		Coins []struct {
			Item struct {
				ID            string  `json:"id"`
				Name          string  `json:"name"`
				Symbol        string  `json:"symbol"`
				MarketCapRank int     `json:"market_cap_rank"`
				PriceBTC      float64 `json:"price_btc"`
				Score         int     `json:"score"`
				Slug          string  `json:"slug"`
			} `json:"item"`
		} `json:"coins"`
	}

	err := c.get("/search/trending", &raw)
	if err != nil {
		return nil, fmt.Errorf("fetch trending: %w", err)
	}

	coins := make([]models.CoinData, len(raw.Coins))
	for i, r := range raw.Coins {
		coins[i] = models.CoinData{
			ID:            r.Item.ID,
			Name:          r.Item.Name,
			Symbol:        r.Item.Symbol,
			MarketCapRank: r.Item.MarketCapRank,
		}
	}
	return coins, nil
}

func (c *Client) FetchGlobal() (*models.GlobalData, error) {
	var raw models.GlobalResponse
	err := c.get("/global", &raw)
	if err != nil {
		return nil, fmt.Errorf("fetch global: %w", err)
	}

	d := raw.Data
	d.BTCDominance = d.MarketCapPercentage["btc"]
	d.ETHDominance = d.MarketCapPercentage["eth"]
	return &d, nil
}
