// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// NewGetAssetReportParams creates a new GetAssetReportParams object
// no default values defined in spec.
func NewGetAssetReportParams() GetAssetReportParams {

	return GetAssetReportParams{}
}

// GetAssetReportParams contains all the bound params for the get asset report operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAssetReport
type GetAssetReportParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	AssetID string
	/*
	  Required: true
	  In: query
	*/
	From int64
	/*Not used
	  Required: true
	  In: path
	*/
	Network string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	OperationType []string
	/*Not used
	  Required: true
	  In: path
	*/
	Platform string
	/*
	  Required: true
	  In: query
	*/
	To int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAssetReportParams() beforehand.
func (o *GetAssetReportParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rAssetID, rhkAssetID, _ := route.Params.GetOK("assetId")
	if err := o.bindAssetID(rAssetID, rhkAssetID, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	qOperationType, qhkOperationType, _ := qs.GetOK("operation_type")
	if err := o.bindOperationType(qOperationType, qhkOperationType, route.Formats); err != nil {
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

// bindAssetID binds and validates parameter AssetID from path.
func (o *GetAssetReportParams) bindAssetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AssetID = raw

	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetAssetReportParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("from", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("from", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("from", "query", "int64", raw)
	}
	o.From = value

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetAssetReportParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}

// bindOperationType binds and validates array parameter OperationType from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAssetReportParams) bindOperationType(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	operationTypeIC := rawData

	if len(operationTypeIC) == 0 {
		return nil
	}

	var operationTypeIR []string
	for _, operationTypeIV := range operationTypeIC {
		operationTypeI := operationTypeIV

		operationTypeIR = append(operationTypeIR, operationTypeI)
	}

	o.OperationType = operationTypeIR

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetAssetReportParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetAssetReportParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("to", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("to", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("to", "query", "int64", raw)
	}
	o.To = value

	return nil
}
