package cmc

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/everstake/teztracker/models"
	coingecko "github.com/superoo7/go-gecko/v3"

	"github.com/patrickmn/go-cache"
)

const (
	tezosPriceURL          = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=tezos&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=24h"
	tezosDenominationRates = "https://api.coingecko.com/api/v3/simple/price?ids=tezos&vs_currencies=usd,eur,gbp,cny"
	cacheTTL               = 5 * time.Minute
	marketInfoKey          = "market_info"
	denominationRatesKey   = "denomination_rates"
)

type tezosMarketData struct {
	USDMarketData
	Prices
}

func (t tezosMarketData) GetPrices() models.Prices {
	return t.Prices
}

// CoinGecko is market data provider.
type CoinGecko struct {
	Cache *cache.Cache
}

func NewCoinGecko() *CoinGecko {
	return &CoinGecko{cache.New(cacheTTL, cacheTTL)}
}

// GetTezosMarketData gets the tezos prices and price change from CoinGecko API.
func (c CoinGecko) GetTezosMarketData() (md models.MarketInfo, err error) {

	usdMarket, err := c.GetTezosUSDMarketData()
	if err != nil {
		return nil, err
	}

	prices, err := c.GetTezosDenominationRates()
	if err != nil {
		return nil, err
	}

	return tezosMarketData{
		USDMarketData: usdMarket,
		Prices:        prices,
	}, nil
}

func (c CoinGecko) GetTezosUSDMarketData() (md USDMarketData, err error) {
	if marketData, isFound := c.Cache.Get(marketInfoKey); isFound {
		return marketData.(USDMarketData), nil
	}

	cg := coingecko.NewClient(nil)
	b, err := cg.MakeReq(tezosPriceURL)
	if err != nil {
		return md, err
	}
	var tmd []USDMarketData
	err = json.Unmarshal(b, &tmd)
	if err != nil {
		return md, err
	}
	if len(tmd) != 1 {
		return md, fmt.Errorf("got enexpected number of entries")
	}

	//Save into cache error can be skipped
	c.Cache.Add(marketInfoKey, tmd[0], cacheTTL)

	return tmd[0], nil
}

func (c CoinGecko) GetTezosDenominationRates() (p Prices, err error) {

	if marketData, isFound := c.Cache.Get(denominationRatesKey); isFound {
		return marketData.(Prices), nil
	}

	cg := coingecko.NewClient(nil)
	b, err := cg.MakeReq(tezosDenominationRates)
	if err != nil {
		return p, err
	}

	resp := map[string]Prices{}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return p, err
	}

	p, ok := resp["tezos"]
	if !ok {
		return
	}

	//Save into cache error can be skipped
	c.Cache.Add(denominationRatesKey, p, cacheTTL)

	return p, nil

}
