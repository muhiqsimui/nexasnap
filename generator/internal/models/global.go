package models

type GlobalData struct {
	ActiveCryptocurrencies    int                `json:"active_cryptocurrencies"`
	TotalMarketCap            map[string]float64 `json:"total_market_cap"`
	TotalVolume               map[string]float64 `json:"total_volume"`
	MarketCapPercentage       map[string]float64 `json:"market_cap_percentage"`
	MarketCapChange24hUSD     float64            `json:"market_cap_change_percentage_24h_usd"`
	BTCDominance              float64            `json:"btc_dominance"`
	ETHDominance              float64            `json:"eth_dominance"`
}

type GlobalResponse struct {
	Data GlobalData `json:"data"`
}
