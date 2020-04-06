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

// NewGetAccountBakedBlocksListParams creates a new GetAccountBakedBlocksListParams object
// with the default values initialized.
func NewGetAccountBakedBlocksListParams() GetAccountBakedBlocksListParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(20)

		offsetDefault = int64(0)
	)

	return GetAccountBakedBlocksListParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetAccountBakedBlocksListParams contains all the bound params for the get account baked blocks list operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAccountBakedBlocksList
type GetAccountBakedBlocksListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	AccountID string
	/*
	  Required: true
	  In: path
	*/
	CycleID int64
	/*
	  Maximum: 500
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*Not used
	  Required: true
	  In: path
	*/
	Network string
	/*Offset
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*Not used
	  Required: true
	  In: path
	*/
	Platform string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAccountBakedBlocksListParams() beforehand.
func (o *GetAccountBakedBlocksListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rAccountID, rhkAccountID, _ := route.Params.GetOK("accountId")
	if err := o.bindAccountID(rAccountID, rhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	rCycleID, rhkCycleID, _ := route.Params.GetOK("cycleId")
	if err := o.bindCycleID(rCycleID, rhkCycleID, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
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

// bindAccountID binds and validates parameter AccountID from path.
func (o *GetAccountBakedBlocksListParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AccountID = raw

	return nil
}

// bindCycleID binds and validates parameter CycleID from path.
func (o *GetAccountBakedBlocksListParams) bindCycleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("cycleId", "path", "int64", raw)
	}
	o.CycleID = value

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetAccountBakedBlocksListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAccountBakedBlocksListParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetAccountBakedBlocksListParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", int64(*o.Limit), 500, false); err != nil {
		return err
	}

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetAccountBakedBlocksListParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetAccountBakedBlocksListParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAccountBakedBlocksListParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetAccountBakedBlocksListParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", int64(*o.Offset), 0, false); err != nil {
		return err
	}

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetAccountBakedBlocksListParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}
