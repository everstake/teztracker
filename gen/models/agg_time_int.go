// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AggTimeInt agg time int
// swagger:model AggTimeInt
type AggTimeInt struct {

	// date
	Date int64 `json:"date,omitempty"`

	// value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this agg time int
func (m *AggTimeInt) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggTimeInt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggTimeInt) UnmarshalBinary(b []byte) error {
	var res AggTimeInt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
