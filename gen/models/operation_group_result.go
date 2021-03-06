// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OperationGroupResult operation group result
// swagger:model OperationGroupResult
type OperationGroupResult struct {

	// operation group
	// Required: true
	OperationGroup *OperationGroupsRow `json:"operation_group"`

	// operations
	// Required: true
	Operations []*OperationsRow `json:"operations"`
}

// Validate validates this operation group result
func (m *OperationGroupResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperationGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationGroupResult) validateOperationGroup(formats strfmt.Registry) error {

	if err := validate.Required("operation_group", "body", m.OperationGroup); err != nil {
		return err
	}

	if m.OperationGroup != nil {
		if err := m.OperationGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation_group")
			}
			return err
		}
	}

	return nil
}

func (m *OperationGroupResult) validateOperations(formats strfmt.Registry) error {

	if err := validate.Required("operations", "body", m.Operations); err != nil {
		return err
	}

	for i := 0; i < len(m.Operations); i++ {
		if swag.IsZero(m.Operations[i]) { // not required
			continue
		}

		if m.Operations[i] != nil {
			if err := m.Operations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationGroupResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationGroupResult) UnmarshalBinary(b []byte) error {
	var res OperationGroupResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
