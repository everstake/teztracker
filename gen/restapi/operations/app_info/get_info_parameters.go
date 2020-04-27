// Code generated by go-swagger; DO NOT EDIT.

package app_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInfoParams creates a new GetInfoParams object
// no default values defined in spec.
func NewGetInfoParams() GetInfoParams {

	return GetInfoParams{}
}

// GetInfoParams contains all the bound params for the get info operation
// typically these are obtained from a http.Request
//
// swagger:parameters getInfo
type GetInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

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
// To ensure default values, the struct must have been initialized with NewGetInfoParams() beforehand.
func (o *GetInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
func (o *GetInfoParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	if err := o.validateNetwork(formats); err != nil {
		return err
	}

	return nil
}

// validateNetwork carries on validations for parameter Network
func (o *GetInfoParams) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Enum("network", "path", o.Network, []interface{}{"mainnet", "carthagenet"}); err != nil {
		return err
	}

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetInfoParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	if err := o.validatePlatform(formats); err != nil {
		return err
	}

	return nil
}

// validatePlatform carries on validations for parameter Platform
func (o *GetInfoParams) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Enum("platform", "path", o.Platform, []interface{}{"tezos"}); err != nil {
		return err
	}

	return nil
}
