// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/everstake/teztracker/gen/models"
)

// NewCreateOrUpdateUserAddressParams creates a new CreateOrUpdateUserAddressParams object
// no default values defined in spec.
func NewCreateOrUpdateUserAddressParams() CreateOrUpdateUserAddressParams {

	return CreateOrUpdateUserAddressParams{}
}

// CreateOrUpdateUserAddressParams contains all the bound params for the create or update user address operation
// typically these are obtained from a http.Request
//
// swagger:parameters createOrUpdateUserAddress
type CreateOrUpdateUserAddressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	Address string
	/*
	  In: body
	*/
	Data *models.UserAddress
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateOrUpdateUserAddressParams() beforehand.
func (o *CreateOrUpdateUserAddressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAddress(r.Header[http.CanonicalHeaderKey("address")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserAddress
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("data", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAddress binds and validates parameter Address from header.
func (o *CreateOrUpdateUserAddressParams) bindAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("address", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("address", "header", raw); err != nil {
		return err
	}

	o.Address = raw

	return nil
}
