// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBakersDelegatorsParams creates a new GetBakersDelegatorsParams object
// no default values defined in spec.
func NewGetBakersDelegatorsParams() GetBakersDelegatorsParams {

	return GetBakersDelegatorsParams{}
}

// GetBakersDelegatorsParams contains all the bound params for the get bakers delegators operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBakersDelegators
type GetBakersDelegatorsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Network string
	/*
	  Required: true
	  In: path
	*/
	Platform string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBakersDelegatorsParams() beforehand.
func (o *GetBakersDelegatorsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

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

// bindNetwork binds and validates parameter Network from path.
func (o *GetBakersDelegatorsParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetBakersDelegatorsParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}
