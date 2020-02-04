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
	tezosPriceURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=tezos&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=24h"
	cacheTTL      = 30 * time.Second
	marketInfoKey = "market_info"
)

type tezosMarketData struct {
	Tezos USDMarketData `json:"tezos"`
}

// CoinGecko is market data provider.
type CoinGecko struct {
	Cache *cache.Cache
}

func NewCoinGecko() *CoinGecko {
	return &CoinGecko{cache.New(cacheTTL, cacheTTL)}
}

// GetTezosMarketData gets the tezos price and price change from CoinGecko API.
func (c CoinGecko) GetTezosMarketData() (md models.MarketInfo, err error) {
	if marketData, isFound := c.Cache.Get(marketInfoKey); isFound {
		return marketData.(models.MarketInfo), nil
	}

	cg := coingecko.NewClient(nil)
	b, err := cg.MakeReq(tezosPriceURL)
	if err != nil {
		return nil, err
	}
	var tmd []USDMarketData
	err = json.Unmarshal(b, &tmd)
	if err != nil {
		return nil, err
	}
	if len(tmd) != 1 {
		return nil, fmt.Errorf("got enexpected number of entries")
	}

	//Save into cache error can be skipped
	c.Cache.Add(marketInfoKey, tmd[0], cacheTTL)

	return tmd[0], nil
}
