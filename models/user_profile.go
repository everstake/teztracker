package models

import (
	"errors"
	"time"
)

var AccountNotFoundErr = errors.New("account not found")
var UserLimitReached = errors.New("limit reached")

type User struct {
	AccountID string
	Email     string
	CreatedAt time.Time
}

type UserAddress struct {
	AccountID          string
	Address            string
	DelegationsEnabled bool
	TransfersEnabled   bool
}

type UserNote struct {
	ID        uint64
	AccountID string
	Text      string
	Alias     string
}

type UserAndAddress struct {
	Email string
	UserAddress
}
