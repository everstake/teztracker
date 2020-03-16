// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BakersRow bakers row
// swagger:model BakersRow
type BakersRow struct {

	// account Id
	AccountID string `json:"accountId,omitempty"`

	// baker info
	BakerInfo *BakerInfo `json:"bakerInfo,omitempty"`
}

// Validate validates this bakers row
func (m *BakersRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBakerInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BakersRow) validateBakerInfo(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BakersRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BakersRow) UnmarshalBinary(b []byte) error {
	var res BakersRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
