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

// UserNoteWithBalance user note with balance
// swagger:model UserNoteWithBalance
type UserNoteWithBalance struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// alias
	// Required: true
	Alias *string `json:"alias"`

	// balance
	// Required: true
	Balance *int64 `json:"balance"`

	// description
	// Required: true
	Description *string `json:"description"`

	// tag
	// Required: true
	Tag *string `json:"tag"`
}

// Validate validates this user note with balance
func (m *UserNoteWithBalance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserNoteWithBalance) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *UserNoteWithBalance) validateAlias(formats strfmt.Registry) error {

	if err := validate.Required("alias", "body", m.Alias); err != nil {
		return err
	}

	return nil
}

func (m *UserNoteWithBalance) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *UserNoteWithBalance) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *UserNoteWithBalance) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserNoteWithBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserNoteWithBalance) UnmarshalBinary(b []byte) error {
	var res UserNoteWithBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
