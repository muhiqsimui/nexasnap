package config

type CoinDef struct {
	CoinGeckoID string
	Slug        string
	Name        string
	Symbol      string
	Category    string
}

func GetCoins() []CoinDef {
	return []CoinDef{
		{CoinGeckoID: "bitcoin", Slug: "bitcoin", Name: "Bitcoin", Symbol: "BTC", Category: "crypto"},
		{CoinGeckoID: "ethereum", Slug: "ethereum", Name: "Ethereum", Symbol: "ETH", Category: "crypto"},
		{CoinGeckoID: "tether", Slug: "tether", Name: "Tether", Symbol: "USDT", Category: "crypto"},
		{CoinGeckoID: "ripple", Slug: "ripple", Name: "XRP", Symbol: "XRP", Category: "crypto"},
		{CoinGeckoID: "binancecoin", Slug: "binance-coin", Name: "BNB", Symbol: "BNB", Category: "crypto"},
		{CoinGeckoID: "solana", Slug: "solana", Name: "Solana", Symbol: "SOL", Category: "crypto"},
		{CoinGeckoID: "usd-coin", Slug: "usd-coin", Name: "USD Coin", Symbol: "USDC", Category: "crypto"},
		{CoinGeckoID: "cardano", Slug: "cardano", Name: "Cardano", Symbol: "ADA", Category: "crypto"},
		{CoinGeckoID: "dogecoin", Slug: "dogecoin", Name: "Dogecoin", Symbol: "DOGE", Category: "crypto"},
		{CoinGeckoID: "sui", Slug: "sui", Name: "Sui", Symbol: "SUI", Category: "crypto"},
		{CoinGeckoID: "tron", Slug: "tron", Name: "TRON", Symbol: "TRX", Category: "crypto"},
		{CoinGeckoID: "the-open-network", Slug: "toncoin", Name: "Toncoin", Symbol: "TON", Category: "crypto"},
		{CoinGeckoID: "avalanche-2", Slug: "avalanche", Name: "Avalanche", Symbol: "AVAX", Category: "crypto"},
		{CoinGeckoID: "chainlink", Slug: "chainlink", Name: "Chainlink", Symbol: "LINK", Category: "crypto"},
		{CoinGeckoID: "polygon-ecosystem-token", Slug: "pol", Name: "POL", Symbol: "POL", Category: "crypto"},
		{CoinGeckoID: "litecoin", Slug: "litecoin", Name: "Litecoin", Symbol: "LTC", Category: "crypto"},
		{CoinGeckoID: "shiba-inu", Slug: "shiba-inu", Name: "Shiba Inu", Symbol: "SHIB", Category: "crypto"},
		{CoinGeckoID: "bitcoin-cash", Slug: "bitcoin-cash", Name: "Bitcoin Cash", Symbol: "BCH", Category: "crypto"},
		{CoinGeckoID: "monero", Slug: "monero", Name: "Monero", Symbol: "XMR", Category: "crypto"},
		{CoinGeckoID: "stellar", Slug: "stellar", Name: "Stellar", Symbol: "XLM", Category: "crypto"},
		{CoinGeckoID: "hedera-hashgraph", Slug: "hedera", Name: "Hedera", Symbol: "HBAR", Category: "crypto"},
		{CoinGeckoID: "crypto-com-chain", Slug: "cronos", Name: "Cronos", Symbol: "CRO", Category: "crypto"},
		{CoinGeckoID: "cosmos", Slug: "cosmos", Name: "Cosmos", Symbol: "ATOM", Category: "crypto"},
		{CoinGeckoID: "ethereum-classic", Slug: "ethereum-classic", Name: "Ethereum Classic", Symbol: "ETC", Category: "crypto"},
		{CoinGeckoID: "dai", Slug: "dai", Name: "Dai", Symbol: "DAI", Category: "crypto"},
		{CoinGeckoID: "filecoin", Slug: "filecoin", Name: "Filecoin", Symbol: "FIL", Category: "crypto"},
	}
}
