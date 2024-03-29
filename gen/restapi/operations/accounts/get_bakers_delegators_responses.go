// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetBakersDelegatorsOKCode is the HTTP code returned for type GetBakersDelegatorsOK
const GetBakersDelegatorsOKCode int = 200

/*GetBakersDelegatorsOK Get number of baker delegators

swagger:response getBakersDelegatorsOK
*/
type GetBakersDelegatorsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.BakerDelegators `json:"body,omitempty"`
}

// NewGetBakersDelegatorsOK creates GetBakersDelegatorsOK with default headers values
func NewGetBakersDelegatorsOK() *GetBakersDelegatorsOK {

	return &GetBakersDelegatorsOK{}
}

// WithPayload adds the payload to the get bakers delegators o k response
func (o *GetBakersDelegatorsOK) WithPayload(payload []*models.BakerDelegators) *GetBakersDelegatorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bakers delegators o k response
func (o *GetBakersDelegatorsOK) SetPayload(payload []*models.BakerDelegators) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBakersDelegatorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.BakerDelegators, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBakersDelegatorsBadRequestCode is the HTTP code returned for type GetBakersDelegatorsBadRequest
const GetBakersDelegatorsBadRequestCode int = 400

/*GetBakersDelegatorsBadRequest Bad request

swagger:response getBakersDelegatorsBadRequest
*/
type GetBakersDelegatorsBadRequest struct {
}

// NewGetBakersDelegatorsBadRequest creates GetBakersDelegatorsBadRequest with default headers values
func NewGetBakersDelegatorsBadRequest() *GetBakersDelegatorsBadRequest {

	return &GetBakersDelegatorsBadRequest{}
}

// WriteResponse to the client
func (o *GetBakersDelegatorsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBakersDelegatorsNotFoundCode is the HTTP code returned for type GetBakersDelegatorsNotFound
const GetBakersDelegatorsNotFoundCode int = 404

/*GetBakersDelegatorsNotFound Not Found

swagger:response getBakersDelegatorsNotFound
*/
type GetBakersDelegatorsNotFound struct {
}

// NewGetBakersDelegatorsNotFound creates GetBakersDelegatorsNotFound with default headers values
func NewGetBakersDelegatorsNotFound() *GetBakersDelegatorsNotFound {

	return &GetBakersDelegatorsNotFound{}
}

// WriteResponse to the client
func (o *GetBakersDelegatorsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBakersDelegatorsInternalServerErrorCode is the HTTP code returned for type GetBakersDelegatorsInternalServerError
const GetBakersDelegatorsInternalServerErrorCode int = 500

/*GetBakersDelegatorsInternalServerError Internal error

swagger:response getBakersDelegatorsInternalServerError
*/
type GetBakersDelegatorsInternalServerError struct {
}

// NewGetBakersDelegatorsInternalServerError creates GetBakersDelegatorsInternalServerError with default headers values
func NewGetBakersDelegatorsInternalServerError() *GetBakersDelegatorsInternalServerError {

	return &GetBakersDelegatorsInternalServerError{}
}

// WriteResponse to the client
func (o *GetBakersDelegatorsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
