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

// NewGetBakersListParams creates a new GetBakersListParams object
// with the default values initialized.
func NewGetBakersListParams() GetBakersListParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(20)

		offsetDefault = int64(0)
	)

	return GetBakersListParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetBakersListParams contains all the bound params for the get bakers list operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBakersList
type GetBakersListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*favorites accounts
	  In: query
	  Collection Format: multi
	*/
	Favorites []string
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
// To ensure default values, the struct must have been initialized with NewGetBakersListParams() beforehand.
func (o *GetBakersListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFavorites, qhkFavorites, _ := qs.GetOK("favorites")
	if err := o.bindFavorites(qFavorites, qhkFavorites, route.Formats); err != nil {
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

// bindFavorites binds and validates array parameter Favorites from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetBakersListParams) bindFavorites(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	favoritesIC := rawData

	if len(favoritesIC) == 0 {
		return nil
	}

	var favoritesIR []string
	for _, favoritesIV := range favoritesIC {
		favoritesI := favoritesIV

		favoritesIR = append(favoritesIR, favoritesI)
	}

	o.Favorites = favoritesIR

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetBakersListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetBakersListParams()
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
func (o *GetBakersListParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", int64(*o.Limit), 500, false); err != nil {
		return err
	}

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetBakersListParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetBakersListParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetBakersListParams()
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
func (o *GetBakersListParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", int64(*o.Offset), 0, false); err != nil {
		return err
	}

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetBakersListParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}
