package models

// MarketInfo is the interface getting prices and price changes.
type MarketInfo interface {
	GetPrices() Prices
	GetPriceChange() float64
	GetMarketCap() float64
	GetVolume() float64
	GetSupply() float64
}

type Prices interface {
	GetUSD() float64
	GetEUR() float64
	GetGBP() float64
	GetCNY() float64
}
