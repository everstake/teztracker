// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountsRow accounts row
// swagger:model AccountsRow
type AccountsRow struct {

	// account Id
	// Required: true
	AccountID *string `json:"accountId"`

	// account name
	AccountName string `json:"accountName,omitempty"`

	// baker info
	BakerInfo *BakerInfo `json:"bakerInfo,omitempty"`

	// balance
	// Required: true
	Balance *int64 `json:"balance"`

	// block Id
	// Required: true
	BlockID *string `json:"blockId"`

	// block level
	// Required: true
	BlockLevel *int64 `json:"blockLevel"`

	// counter
	// Required: true
	Counter *int64 `json:"counter"`

	// created at
	CreatedAt int64 `json:"createdAt,omitempty"`

	// delegate setable
	// Required: true
	DelegateSetable *bool `json:"delegateSetable"`

	// delegate value
	DelegateValue string `json:"delegateValue,omitempty"`

	// last active
	LastActive int64 `json:"lastActive,omitempty"`

	// manager
	// Required: true
	Manager *string `json:"manager"`

	// operations
	Operations int64 `json:"operations,omitempty"`

	// revealed
	Revealed bool `json:"revealed,omitempty"`

	// script
	Script string `json:"script,omitempty"`

	// spendable
	// Required: true
	Spendable *bool `json:"spendable"`

	// storage
	Storage string `json:"storage,omitempty"`

	// transactions
	Transactions int64 `json:"transactions,omitempty"`
}

// Validate validates this accounts row
func (m *AccountsRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBakerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegateSetable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpendable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsRow) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateBakerInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.BakerInfo) { // not required
		return nil
	}

	if m.BakerInfo != nil {
		if err := m.BakerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bakerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *AccountsRow) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateBlockID(formats strfmt.Registry) error {

	if err := validate.Required("blockId", "body", m.BlockID); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateBlockLevel(formats strfmt.Registry) error {

	if err := validate.Required("blockLevel", "body", m.BlockLevel); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateCounter(formats strfmt.Registry) error {

	if err := validate.Required("counter", "body", m.Counter); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateDelegateSetable(formats strfmt.Registry) error {

	if err := validate.Required("delegateSetable", "body", m.DelegateSetable); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateManager(formats strfmt.Registry) error {

	if err := validate.Required("manager", "body", m.Manager); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateSpendable(formats strfmt.Registry) error {

	if err := validate.Required("spendable", "body", m.Spendable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountsRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountsRow) UnmarshalBinary(b []byte) error {
	var res AccountsRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
