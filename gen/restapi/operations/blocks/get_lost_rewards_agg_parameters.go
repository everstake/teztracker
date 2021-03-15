// Code generated by go-swagger; DO NOT EDIT.

package blocks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLostRewardsAggParams creates a new GetLostRewardsAggParams object
// no default values defined in spec.
func NewGetLostRewardsAggParams() GetLostRewardsAggParams {

	return GetLostRewardsAggParams{}
}

// GetLostRewardsAggParams contains all the bound params for the get lost rewards agg operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLostRewardsAgg
type GetLostRewardsAggParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Network string
	/*
	  Required: true
	  In: query
	*/
	Period string
	/*
	  Required: true
	  In: path
	*/
	Platform string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLostRewardsAggParams() beforehand.
func (o *GetLostRewardsAggParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	qPeriod, qhkPeriod, _ := qs.GetOK("period")
	if err := o.bindPeriod(qPeriod, qhkPeriod, route.Formats); err != nil {
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
func (o *GetLostRewardsAggParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}

// bindPeriod binds and validates parameter Period from query.
func (o *GetLostRewardsAggParams) bindPeriod(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("period", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("period", "query", raw); err != nil {
		return err
	}

	o.Period = raw

	if err := o.validatePeriod(formats); err != nil {
		return err
	}

	return nil
}

// validatePeriod carries on validations for parameter Period
func (o *GetLostRewardsAggParams) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Enum("period", "query", o.Period, []interface{}{"day", "week", "month"}); err != nil {
		return err
	}

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetLostRewardsAggParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}
