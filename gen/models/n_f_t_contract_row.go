// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NFTContractRow n f t contract row
// swagger:model NFTContractRow
type NFTContractRow struct {

	// address
	Address string `json:"address,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nfts number
	NftsNumber int64 `json:"nfts_number,omitempty"`

	// operations number
	OperationsNumber int64 `json:"operations_number,omitempty"`
}

// Validate validates this n f t contract row
func (m *NFTContractRow) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NFTContractRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NFTContractRow) UnmarshalBinary(b []byte) error {
	var res NFTContractRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
