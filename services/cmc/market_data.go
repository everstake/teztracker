package cmc

// MarketData is a Price and Price Change with json deserialization for USD .
type USDMarketData struct {
	Price          float64 `json:"current_price"`
	Price24hChange float64 `json:"price_change_24h"`
	MarketCap      float64 `json:"market_cap"`
	Volume         float64 `json:"total_volume"`
	Supply         float64 `json:"circulating_supply"`
}

// GetPrice returns the price in USD.
func (md USDMarketData) GetPrice() float64 {
	return md.Price
}

// GetPriceChange returns the price change during the last 24 hours in percents.
func (md USDMarketData) GetPriceChange() float64 {
	return md.Price24hChange
}
func (md USDMarketData) GetMarketCap() float64 {
	return md.MarketCap
}
func (md USDMarketData) GetVolume() float64 {
	return md.Volume
}
func (md USDMarketData) GetSupply() float64 {
	return md.Supply
}
