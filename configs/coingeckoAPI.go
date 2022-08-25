package configs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const APIEndpoint = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd"

type Coin struct {
	ID            string  `json:"id"`
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	CurrentPrice  float64 `json:"current_price"`
	MarketCapRank int     `json:"market_cap_rank"`
	LastUpdated   string  `json:"last_updated"`
}

func GetMarketCoins() ([]Coin, error) {

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.Get(APIEndpoint)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req.Request)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		return nil, errors.New("API Request Error")
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	var coins []Coin

	err = json.Unmarshal(body, &coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}

func GetSpecificCoins(coins []Coin) ([]Coin, error) {

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	ids := ""
	for i := range coins {
		ids = ids + string(coins[i].ID) + ","
	}
	CoinsEndpoint := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=" + string(ids)

	req, err := http.Get(CoinsEndpoint)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req.Request)
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		return nil, errors.New("API Request Error")
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	var new_coins []Coin

	err = json.Unmarshal(body, &new_coins)
	if err != nil {
		return nil, err
	}

	return new_coins, nil
}
