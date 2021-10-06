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

// NFTTokenRow n f t token row
// swagger:model NFTTokenRow
type NFTTokenRow struct {

	// amount
	Amount int64 `json:"amount,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// created at
	CreatedAt int64 `json:"created_at,omitempty"`

	// decimals
	// Required: true
	Decimals *int64 `json:"decimals"`

	// description
	Description string `json:"description,omitempty"`

	// ipfs source
	IpfsSource string `json:"ipfs_source,omitempty"`

	// is for sale
	IsForSale bool `json:"is_for_sale,omitempty"`

	// issued by
	IssuedBy string `json:"issued_by,omitempty"`

	// last active at
	LastActiveAt int64 `json:"last_active_at,omitempty"`

	// last price
	LastPrice int64 `json:"last_price,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// token id
	TokenID int64 `json:"token_id,omitempty"`
}

// Validate validates this n f t token row
func (m *NFTTokenRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecimals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NFTTokenRow) validateDecimals(formats strfmt.Registry) error {

	if err := validate.Required("decimals", "body", m.Decimals); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NFTTokenRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NFTTokenRow) UnmarshalBinary(b []byte) error {
	var res NFTTokenRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
