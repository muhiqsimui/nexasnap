package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PriceResponse map[string]struct {
	USD float64 `json:"usd"`
}

func FetchPrices() (PriceResponse, error) {

	url :=
		"https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd"

		client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"User-Agent",
		"NexaSnap/1.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	return nil, fmt.Errorf(
		"unexpected status code: %d",
		resp.StatusCode,
	)
}

	var result PriceResponse

	err = json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}