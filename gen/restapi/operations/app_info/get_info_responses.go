// Code generated by go-swagger; DO NOT EDIT.

package app_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetInfoOKCode is the HTTP code returned for type GetInfoOK
const GetInfoOKCode int = 200

/*GetInfoOK Application info endpoint

swagger:response getInfoOK
*/
type GetInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Info `json:"body,omitempty"`
}

// NewGetInfoOK creates GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {

	return &GetInfoOK{}
}

// WithPayload adds the payload to the get info o k response
func (o *GetInfoOK) WithPayload(payload *models.Info) *GetInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get info o k response
func (o *GetInfoOK) SetPayload(payload *models.Info) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfoBadRequestCode is the HTTP code returned for type GetInfoBadRequest
const GetInfoBadRequestCode int = 400

/*GetInfoBadRequest Bad request

swagger:response getInfoBadRequest
*/
type GetInfoBadRequest struct {
}

// NewGetInfoBadRequest creates GetInfoBadRequest with default headers values
func NewGetInfoBadRequest() *GetInfoBadRequest {

	return &GetInfoBadRequest{}
}

// WriteResponse to the client
func (o *GetInfoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetInfoInternalServerErrorCode is the HTTP code returned for type GetInfoInternalServerError
const GetInfoInternalServerErrorCode int = 500

/*GetInfoInternalServerError Internal error

swagger:response getInfoInternalServerError
*/
type GetInfoInternalServerError struct {
}

// NewGetInfoInternalServerError creates GetInfoInternalServerError with default headers values
func NewGetInfoInternalServerError() *GetInfoInternalServerError {

	return &GetInfoInternalServerError{}
}

// WriteResponse to the client
func (o *GetInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
