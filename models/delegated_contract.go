package models

type DelegatedContract struct {
	AccountID     string   `json:"account_id"`
	Account       *Account `json:"account"` // This line is infered from column name "account_id".
	DelegateValue string   `json:"delegate_value"`
}
