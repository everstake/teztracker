// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HoldingPoint holding point
// swagger:model HoldingPoint
type HoldingPoint struct {

	// amount
	Amount int64 `json:"amount,omitempty"`

	// count
	Count int64 `json:"count,omitempty"`

	// percent
	Percent float64 `json:"percent,omitempty"`
}

// Validate validates this holding point
func (m *HoldingPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HoldingPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HoldingPoint) UnmarshalBinary(b []byte) error {
	var res HoldingPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
