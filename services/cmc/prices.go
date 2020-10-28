package cmc

type Prices struct {
	USD float64 `json:"usd"`
	EUR float64 `json:"eur"`
	GBP float64 `json:"gbp"`
	CNY float64 `json:"cny"`
}

func (p Prices) GetUSD() float64 {
	return p.USD
}

func (p Prices) GetEUR() float64 {
	return p.EUR
}

func (p Prices) GetGBP() float64 {
	return p.GBP
}

func (p Prices) GetCNY() float64 {
	return p.CNY
}
