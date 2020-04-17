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

// AccountEndorsingRow account endorsing row
// swagger:model AccountEndorsingRow
type AccountEndorsingRow struct {

	// cycle
	// Required: true
	Cycle *int64 `json:"cycle"`

	// missed
	// Required: true
	Missed *int64 `json:"missed"`

	// rewards
	// Required: true
	Rewards *int64 `json:"rewards"`

	// slots
	// Required: true
	Slots *int64 `json:"slots"`

	// status
	Status string `json:"status,omitempty"`

	// total deposit
	// Required: true
	TotalDeposit *int64 `json:"totalDeposit"`
}

// Validate validates this account endorsing row
func (m *AccountEndorsingRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDeposit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountEndorsingRow) validateCycle(formats strfmt.Registry) error {

	if err := validate.Required("cycle", "body", m.Cycle); err != nil {
		return err
	}

	return nil
}

func (m *AccountEndorsingRow) validateMissed(formats strfmt.Registry) error {

	if err := validate.Required("missed", "body", m.Missed); err != nil {
		return err
	}

	return nil
}

func (m *AccountEndorsingRow) validateRewards(formats strfmt.Registry) error {

	if err := validate.Required("rewards", "body", m.Rewards); err != nil {
		return err
	}

	return nil
}

func (m *AccountEndorsingRow) validateSlots(formats strfmt.Registry) error {

	if err := validate.Required("slots", "body", m.Slots); err != nil {
		return err
	}

	return nil
}

func (m *AccountEndorsingRow) validateTotalDeposit(formats strfmt.Registry) error {

	if err := validate.Required("totalDeposit", "body", m.TotalDeposit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountEndorsingRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountEndorsingRow) UnmarshalBinary(b []byte) error {
	var res AccountEndorsingRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
