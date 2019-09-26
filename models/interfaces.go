package models

// MarketInfo is the interface getting prices and price changes.
type MarketInfo interface {
	GetPrice() float64
	GetPriceChange() float64
}
