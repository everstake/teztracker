// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetBakersHoldingOKCode is the HTTP code returned for type GetBakersHoldingOK
const GetBakersHoldingOKCode int = 200

/*GetBakersHoldingOK Get bakers holding points

swagger:response getBakersHoldingOK
*/
type GetBakersHoldingOK struct {

	/*
	  In: Body
	*/
	Payload []*models.HoldingPoint `json:"body,omitempty"`
}

// NewGetBakersHoldingOK creates GetBakersHoldingOK with default headers values
func NewGetBakersHoldingOK() *GetBakersHoldingOK {

	return &GetBakersHoldingOK{}
}

// WithPayload adds the payload to the get bakers holding o k response
func (o *GetBakersHoldingOK) WithPayload(payload []*models.HoldingPoint) *GetBakersHoldingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bakers holding o k response
func (o *GetBakersHoldingOK) SetPayload(payload []*models.HoldingPoint) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBakersHoldingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.HoldingPoint, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBakersHoldingBadRequestCode is the HTTP code returned for type GetBakersHoldingBadRequest
const GetBakersHoldingBadRequestCode int = 400

/*GetBakersHoldingBadRequest Bad request

swagger:response getBakersHoldingBadRequest
*/
type GetBakersHoldingBadRequest struct {
}

// NewGetBakersHoldingBadRequest creates GetBakersHoldingBadRequest with default headers values
func NewGetBakersHoldingBadRequest() *GetBakersHoldingBadRequest {

	return &GetBakersHoldingBadRequest{}
}

// WriteResponse to the client
func (o *GetBakersHoldingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBakersHoldingNotFoundCode is the HTTP code returned for type GetBakersHoldingNotFound
const GetBakersHoldingNotFoundCode int = 404

/*GetBakersHoldingNotFound Not Found

swagger:response getBakersHoldingNotFound
*/
type GetBakersHoldingNotFound struct {
}

// NewGetBakersHoldingNotFound creates GetBakersHoldingNotFound with default headers values
func NewGetBakersHoldingNotFound() *GetBakersHoldingNotFound {

	return &GetBakersHoldingNotFound{}
}

// WriteResponse to the client
func (o *GetBakersHoldingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBakersHoldingInternalServerErrorCode is the HTTP code returned for type GetBakersHoldingInternalServerError
const GetBakersHoldingInternalServerErrorCode int = 500

/*GetBakersHoldingInternalServerError Internal error

swagger:response getBakersHoldingInternalServerError
*/
type GetBakersHoldingInternalServerError struct {
}

// NewGetBakersHoldingInternalServerError creates GetBakersHoldingInternalServerError with default headers values
func NewGetBakersHoldingInternalServerError() *GetBakersHoldingInternalServerError {

	return &GetBakersHoldingInternalServerError{}
}

// WriteResponse to the client
func (o *GetBakersHoldingInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
