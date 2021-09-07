// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

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

// NewGetNFTContractTokensListParams creates a new GetNFTContractTokensListParams object
// with the default values initialized.
func NewGetNFTContractTokensListParams() GetNFTContractTokensListParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(20)

		offsetDefault = int64(0)
	)

	return GetNFTContractTokensListParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetNFTContractTokensListParams contains all the bound params for the get n f t contract tokens list operation
// typically these are obtained from a http.Request
//
// swagger:parameters getNFTContractTokensList
type GetNFTContractTokensListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ContractID string
	/*
	  Maximum: 300
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*
	  Required: true
	  In: path
	*/
	Network string
	/*
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Offset *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNFTContractTokensListParams() beforehand.
func (o *GetNFTContractTokensListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rContractID, rhkContractID, _ := route.Params.GetOK("contract_id")
	if err := o.bindContractID(rContractID, rhkContractID, route.Formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindContractID binds and validates parameter ContractID from path.
func (o *GetNFTContractTokensListParams) bindContractID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ContractID = raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetNFTContractTokensListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetNFTContractTokensListParams()
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
func (o *GetNFTContractTokensListParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", int64(*o.Limit), 300, false); err != nil {
		return err
	}

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetNFTContractTokensListParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetNFTContractTokensListParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetNFTContractTokensListParams()
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
func (o *GetNFTContractTokensListParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", int64(*o.Offset), 0, false); err != nil {
		return err
	}

	return nil
}
