// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetAccountTotalEndorsingingListOKCode is the HTTP code returned for type GetAccountTotalEndorsingingListOK
const GetAccountTotalEndorsingingListOKCode int = 200

/*GetAccountTotalEndorsingingListOK Query compatibility endpoint for account baking

swagger:response getAccountTotalEndorsingingListOK
*/
type GetAccountTotalEndorsingingListOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccountEndorsingRow `json:"body,omitempty"`
}

// NewGetAccountTotalEndorsingingListOK creates GetAccountTotalEndorsingingListOK with default headers values
func NewGetAccountTotalEndorsingingListOK() *GetAccountTotalEndorsingingListOK {

	return &GetAccountTotalEndorsingingListOK{}
}

// WithPayload adds the payload to the get account total endorsinging list o k response
func (o *GetAccountTotalEndorsingingListOK) WithPayload(payload *models.AccountEndorsingRow) *GetAccountTotalEndorsingingListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account total endorsinging list o k response
func (o *GetAccountTotalEndorsingingListOK) SetPayload(payload *models.AccountEndorsingRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountTotalEndorsingingListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAccountTotalEndorsingingListBadRequestCode is the HTTP code returned for type GetAccountTotalEndorsingingListBadRequest
const GetAccountTotalEndorsingingListBadRequestCode int = 400

/*GetAccountTotalEndorsingingListBadRequest Bad request

swagger:response getAccountTotalEndorsingingListBadRequest
*/
type GetAccountTotalEndorsingingListBadRequest struct {
}

// NewGetAccountTotalEndorsingingListBadRequest creates GetAccountTotalEndorsingingListBadRequest with default headers values
func NewGetAccountTotalEndorsingingListBadRequest() *GetAccountTotalEndorsingingListBadRequest {

	return &GetAccountTotalEndorsingingListBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountTotalEndorsingingListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAccountTotalEndorsingingListNotFoundCode is the HTTP code returned for type GetAccountTotalEndorsingingListNotFound
const GetAccountTotalEndorsingingListNotFoundCode int = 404

/*GetAccountTotalEndorsingingListNotFound Not Found

swagger:response getAccountTotalEndorsingingListNotFound
*/
type GetAccountTotalEndorsingingListNotFound struct {
}

// NewGetAccountTotalEndorsingingListNotFound creates GetAccountTotalEndorsingingListNotFound with default headers values
func NewGetAccountTotalEndorsingingListNotFound() *GetAccountTotalEndorsingingListNotFound {

	return &GetAccountTotalEndorsingingListNotFound{}
}

// WriteResponse to the client
func (o *GetAccountTotalEndorsingingListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}