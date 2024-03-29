// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetNFTContractTokensOKCode is the HTTP code returned for type GetNFTContractTokensOK
const GetNFTContractTokensOKCode int = 200

/*GetNFTContractTokensOK Query compatibility endpoint for NFT tokens

swagger:response getNFTContractTokensOK
*/
type GetNFTContractTokensOK struct {

	/*
	  In: Body
	*/
	Payload []*models.NFTTokenRow `json:"body,omitempty"`
}

// NewGetNFTContractTokensOK creates GetNFTContractTokensOK with default headers values
func NewGetNFTContractTokensOK() *GetNFTContractTokensOK {

	return &GetNFTContractTokensOK{}
}

// WithPayload adds the payload to the get n f t contract tokens o k response
func (o *GetNFTContractTokensOK) WithPayload(payload []*models.NFTTokenRow) *GetNFTContractTokensOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get n f t contract tokens o k response
func (o *GetNFTContractTokensOK) SetPayload(payload []*models.NFTTokenRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNFTContractTokensOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.NFTTokenRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetNFTContractTokensBadRequestCode is the HTTP code returned for type GetNFTContractTokensBadRequest
const GetNFTContractTokensBadRequestCode int = 400

/*GetNFTContractTokensBadRequest Bad request

swagger:response getNFTContractTokensBadRequest
*/
type GetNFTContractTokensBadRequest struct {
}

// NewGetNFTContractTokensBadRequest creates GetNFTContractTokensBadRequest with default headers values
func NewGetNFTContractTokensBadRequest() *GetNFTContractTokensBadRequest {

	return &GetNFTContractTokensBadRequest{}
}

// WriteResponse to the client
func (o *GetNFTContractTokensBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetNFTContractTokensNotFoundCode is the HTTP code returned for type GetNFTContractTokensNotFound
const GetNFTContractTokensNotFoundCode int = 404

/*GetNFTContractTokensNotFound Not Found

swagger:response getNFTContractTokensNotFound
*/
type GetNFTContractTokensNotFound struct {
}

// NewGetNFTContractTokensNotFound creates GetNFTContractTokensNotFound with default headers values
func NewGetNFTContractTokensNotFound() *GetNFTContractTokensNotFound {

	return &GetNFTContractTokensNotFound{}
}

// WriteResponse to the client
func (o *GetNFTContractTokensNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetNFTContractTokensInternalServerErrorCode is the HTTP code returned for type GetNFTContractTokensInternalServerError
const GetNFTContractTokensInternalServerErrorCode int = 500

/*GetNFTContractTokensInternalServerError Internal error

swagger:response getNFTContractTokensInternalServerError
*/
type GetNFTContractTokensInternalServerError struct {
}

// NewGetNFTContractTokensInternalServerError creates GetNFTContractTokensInternalServerError with default headers values
func NewGetNFTContractTokensInternalServerError() *GetNFTContractTokensInternalServerError {

	return &GetNFTContractTokensInternalServerError{}
}

// WriteResponse to the client
func (o *GetNFTContractTokensInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
