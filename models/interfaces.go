package models

// MarketInfo is the interface getting prices and price changes.
type MarketInfo interface {
	GetPrice() float64
	GetPriceChange() float64
	GetMarketCap() float64
	GetVolume() float64
	GetSupply() float64
}
