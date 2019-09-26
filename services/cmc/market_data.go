package cmc

// MarketData is a Price and Price Change with json deserialization for USD .
type USDMarketData struct {
	Price          float64 `json:"usd"`
	Price24hChange float64 `json:"usd_24h_change"`
}

// GetPrice returns the price in USD.
func (md *USDMarketData) GetPrice() float64 {
	return md.Price
}

// GetPriceChange returns the price change during the last 24 hours in percents.
func (md *USDMarketData) GetPriceChange() float64 {
	return md.Price24hChange
}
