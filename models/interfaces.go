package models

// MarketInfo is the interface getting prices and price changes.
type MarketInfo interface {
	GetPrice() float64
	GetPriceChange() float64
	GetMarketCap() float64
	GetVolume() float64
	GetSupply() float64
}

// MarketDataProvider is an interface for getting actual price and price changes.
type MarketDataProvider interface {
	GetTezosMarketData(curr string) (md MarketInfo, err error)
}
