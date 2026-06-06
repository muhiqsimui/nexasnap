package models

type CoinData struct {
	ID                       string  `json:"id"`
	Symbol                   string  `json:"symbol"`
	Name                     string  `json:"name"`
	Image                    string  `json:"image,omitempty"`
	CurrentPrice             float64 `json:"current_price"`
	MarketCap                float64 `json:"market_cap"`
	MarketCapRank            int     `json:"market_cap_rank"`
	TotalVolume              float64 `json:"total_volume"`
	High24h                  float64 `json:"high_24h"`
	Low24h                   float64 `json:"low_24h"`
	PriceChange24h           float64 `json:"price_change_24h"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	CirculatingSupply        float64 `json:"circulating_supply"`
	TotalSupply              float64 `json:"total_supply,omitempty"`
	MaxSupply                float64 `json:"max_supply,omitempty"`
	Ath                      float64 `json:"ath"`
	AthDate                  string  `json:"ath_date"`
	Atl                      float64 `json:"atl"`
	AtlDate                  string  `json:"atl_date"`
	LastUpdated              string  `json:"last_updated"`
}
