// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInactiveAccountsAggCountParams creates a new GetInactiveAccountsAggCountParams object
// no default values defined in spec.
func NewGetInactiveAccountsAggCountParams() GetInactiveAccountsAggCountParams {

	return GetInactiveAccountsAggCountParams{}
}

// GetInactiveAccountsAggCountParams contains all the bound params for the get inactive accounts agg count operation
// typically these are obtained from a http.Request
//
// swagger:parameters getInactiveAccountsAggCount
type GetInactiveAccountsAggCountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	From *int64
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
	/*
	  In: query
	*/
	To *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetInactiveAccountsAggCountParams() beforehand.
func (o *GetInactiveAccountsAggCountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

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

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetInactiveAccountsAggCountParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("from", "query", "int64", raw)
	}
	o.From = &value

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetInactiveAccountsAggCountParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetInactiveAccountsAggCountParams) bindPeriod(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetInactiveAccountsAggCountParams) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Enum("period", "query", o.Period, []interface{}{"day", "week", "month"}); err != nil {
		return err
	}

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetInactiveAccountsAggCountParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetInactiveAccountsAggCountParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("to", "query", "int64", raw)
	}
	o.To = &value

	return nil
}