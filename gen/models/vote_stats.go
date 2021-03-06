// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VoteStats vote stats
// swagger:model VoteStats
type VoteStats struct {

	// num voters
	NumVoters int64 `json:"numVoters,omitempty"`

	// num voters total
	NumVotersTotal int64 `json:"numVotersTotal,omitempty"`

	// votes available
	VotesAvailable int64 `json:"votesAvailable,omitempty"`

	// votes cast
	VotesCast int64 `json:"votesCast,omitempty"`
}

// Validate validates this vote stats
func (m *VoteStats) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VoteStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoteStats) UnmarshalBinary(b []byte) error {
	var res VoteStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
