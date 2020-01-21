package cmc

import (
	"encoding/json"
	"fmt"

	"github.com/bullblock-io/tezTracker/models"
	coingecko "github.com/superoo7/go-gecko/v3"
)

const tezosPriceURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=tezos&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=24h"

type tezosMarketData struct {
	Tezos USDMarketData `json:"tezos"`
}

// CoinGecko is market data provider.
type CoinGecko struct{}

// GetTezosMarketData gets the tezos price and price change from CoinGecko API.
// TODO: some caching layer should be implemented.
func (CoinGecko) GetTezosMarketData() (md models.MarketInfo, err error) {
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

	return &tmd[0], nil
}
