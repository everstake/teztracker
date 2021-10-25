// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNFTContractDistributionParams creates a new GetNFTContractDistributionParams object
// no default values defined in spec.
func NewGetNFTContractDistributionParams() GetNFTContractDistributionParams {

	return GetNFTContractDistributionParams{}
}

// GetNFTContractDistributionParams contains all the bound params for the get n f t contract distribution operation
// typically these are obtained from a http.Request
//
// swagger:parameters getNFTContractDistribution
type GetNFTContractDistributionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ContractID string
	/*
	  Required: true
	  In: path
	*/
	Network string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNFTContractDistributionParams() beforehand.
func (o *GetNFTContractDistributionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rContractID, rhkContractID, _ := route.Params.GetOK("contract_id")
	if err := o.bindContractID(rContractID, rhkContractID, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindContractID binds and validates parameter ContractID from path.
func (o *GetNFTContractDistributionParams) bindContractID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ContractID = raw

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetNFTContractDistributionParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}