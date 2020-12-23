// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserAddress user address
// swagger:model UserAddress
type UserAddress struct {

	// address
	Address string `json:"address,omitempty"`

	// delegations enabled
	DelegationsEnabled bool `json:"delegations:enabled,omitempty"`

	// transfers enabled
	TransfersEnabled bool `json:"transfers_enabled,omitempty"`
}

// Validate validates this user address
func (m *UserAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAddress) UnmarshalBinary(b []byte) error {
	var res UserAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}