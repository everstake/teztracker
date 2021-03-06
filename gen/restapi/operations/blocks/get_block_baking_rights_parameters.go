// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBlockBakingRightsParams creates a new GetBlockBakingRightsParams object
// no default values defined in spec.
func NewGetBlockBakingRightsParams() GetBlockBakingRightsParams {

	return GetBlockBakingRightsParams{}
}

// GetBlockBakingRightsParams contains all the bound params for the get block baking rights operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBlockBakingRights
type GetBlockBakingRightsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Hash string
	/*Not used
	  Required: true
	  In: path
	*/
	Network string
	/*Not used
	  Required: true
	  In: path
	*/
	Platform string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBlockBakingRightsParams() beforehand.
func (o *GetBlockBakingRightsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rHash, rhkHash, _ := route.Params.GetOK("hash")
	if err := o.bindHash(rHash, rhkHash, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	rPlatform, rhkPlatform, _ := route.Params.GetOK("platform")
	if err := o.bindPlatform(rPlatform, rhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHash binds and validates parameter Hash from path.
func (o *GetBlockBakingRightsParams) bindHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Hash = raw

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetBlockBakingRightsParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetBlockBakingRightsParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}
