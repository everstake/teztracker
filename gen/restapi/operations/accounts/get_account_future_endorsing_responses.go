// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetAccountFutureEndorsingOKCode is the HTTP code returned for type GetAccountFutureEndorsingOK
const GetAccountFutureEndorsingOKCode int = 200

/*GetAccountFutureEndorsingOK Query compatibility endpoint for account future baking

swagger:response getAccountFutureEndorsingOK
*/
type GetAccountFutureEndorsingOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AccountEndorsingRow `json:"body,omitempty"`
}

// NewGetAccountFutureEndorsingOK creates GetAccountFutureEndorsingOK with default headers values
func NewGetAccountFutureEndorsingOK() *GetAccountFutureEndorsingOK {

	return &GetAccountFutureEndorsingOK{}
}

// WithPayload adds the payload to the get account future endorsing o k response
func (o *GetAccountFutureEndorsingOK) WithPayload(payload []*models.AccountEndorsingRow) *GetAccountFutureEndorsingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account future endorsing o k response
func (o *GetAccountFutureEndorsingOK) SetPayload(payload []*models.AccountEndorsingRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountFutureEndorsingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AccountEndorsingRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountFutureEndorsingBadRequestCode is the HTTP code returned for type GetAccountFutureEndorsingBadRequest
const GetAccountFutureEndorsingBadRequestCode int = 400

/*GetAccountFutureEndorsingBadRequest Bad request

swagger:response getAccountFutureEndorsingBadRequest
*/
type GetAccountFutureEndorsingBadRequest struct {
}

// NewGetAccountFutureEndorsingBadRequest creates GetAccountFutureEndorsingBadRequest with default headers values
func NewGetAccountFutureEndorsingBadRequest() *GetAccountFutureEndorsingBadRequest {

	return &GetAccountFutureEndorsingBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountFutureEndorsingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAccountFutureEndorsingNotFoundCode is the HTTP code returned for type GetAccountFutureEndorsingNotFound
const GetAccountFutureEndorsingNotFoundCode int = 404

/*GetAccountFutureEndorsingNotFound Not Found

swagger:response getAccountFutureEndorsingNotFound
*/
type GetAccountFutureEndorsingNotFound struct {
}

// NewGetAccountFutureEndorsingNotFound creates GetAccountFutureEndorsingNotFound with default headers values
func NewGetAccountFutureEndorsingNotFound() *GetAccountFutureEndorsingNotFound {

	return &GetAccountFutureEndorsingNotFound{}
}

// WriteResponse to the client
func (o *GetAccountFutureEndorsingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
