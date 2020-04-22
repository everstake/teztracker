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

// ChartsData charts data
// swagger:model ChartsData
type ChartsData struct {

	// activations
	Activations int64 `json:"activations,omitempty"`

	// average delay
	AverageDelay float64 `json:"averageDelay,omitempty"`

	// bakers
	Bakers int64 `json:"bakers,omitempty"`

	// block priority counter
	BlockPriorityCounter *BlockPriorityCounter `json:"blockPriorityCounter,omitempty"`

	// blocks
	Blocks int64 `json:"blocks,omitempty"`

	// delegation volume
	DelegationVolume int64 `json:"delegationVolume,omitempty"`

	// fees
	Fees int64 `json:"fees,omitempty"`

	// operations
	Operations int64 `json:"operations,omitempty"`

	// timestamp
	// Required: true
	Timestamp *int64 `json:"timestamp"`

	// transaction volume
	TransactionVolume int64 `json:"transactionVolume,omitempty"`
}

// Validate validates this charts data
func (m *ChartsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockPriorityCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartsData) validateBlockPriorityCounter(formats strfmt.Registry) error {

	if swag.IsZero(m.BlockPriorityCounter) { // not required
		return nil
	}

	if m.BlockPriorityCounter != nil {
		if err := m.BlockPriorityCounter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blockPriorityCounter")
			}
			return err
		}
	}

	return nil
}

func (m *ChartsData) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChartsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChartsData) UnmarshalBinary(b []byte) error {
	var res ChartsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
