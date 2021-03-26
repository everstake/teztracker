package models

type HoldingPoint struct {
	Percent float64 `json:"percent"`
	Amount  int64   `json:"amount"`
	Count   int64   `json:"count"`
}
