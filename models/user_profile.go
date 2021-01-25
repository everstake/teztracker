package models

import (
	"errors"
	"github.com/guregu/null"
	"time"
)

var AccountNotFoundErr = errors.New("account not found")
var UserLimitReachedErr = errors.New("limit reached")

type User struct {
	AccountID string
	Username  string
	Email     string
	Verified  bool
	CreatedAt time.Time
}

type UserAddress struct {
	AccountID           string
	Address             string
	DelegationsEnabled  bool
	InTransfersEnabled  bool
	OutTransfersEnabled bool
}

type UserAddressWithBalance struct {
	UserAddress
	Balance null.Int
}

type UserNote struct {
	ID          uint64
	AccountID   string
	Address     string
	Alias       string
	Tag         string
	Description string
}

type UserNoteWithBalance struct {
	UserNote
	Balance null.Int
}

type UserAddressWithEmail struct {
	Email string
	UserAddress
}

type EmailVerification struct {
	AccountID string
	Email     string
	Token     string
	Verified  bool
	Sent      bool
	CreatedAt time.Time
}
