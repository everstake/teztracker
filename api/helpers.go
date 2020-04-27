package api

import "github.com/everstake/teztracker/models"

import "strings"

import "fmt"

type Limits struct {
	limit  uint
	offset uint
}

func (l *Limits) Limit() uint {
	return l.limit
}
func (l *Limits) Offset() uint {
	return l.offset
}

func NewLimiter(limit, offset *int64) *Limits {
	var l, o uint
	if limit != nil {
		l = uint(*limit)
	}
	if offset != nil {
		o = uint(*offset)
	}
	return &Limits{limit: l, offset: o}
}

func ToNetwork(net string) (models.Network, error) {
	switch strings.ToLower(net) {
	case "main", "mainnet":
		return models.NetworkMain, nil
	case "carthage", "carthagenet":
		return models.NetworkCarthage, nil
	}

	return "", fmt.Errorf("not supported network")
}
