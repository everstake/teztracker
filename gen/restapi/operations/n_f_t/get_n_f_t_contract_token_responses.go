// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetNFTContractTokenOKCode is the HTTP code returned for type GetNFTContractTokenOK
const GetNFTContractTokenOKCode int = 200

/*GetNFTContractTokenOK Query compatibility endpoint for NFT token

swagger:response getNFTContractTokenOK
*/
type GetNFTContractTokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.NFTTokenRow `json:"body,omitempty"`
}

// NewGetNFTContractTokenOK creates GetNFTContractTokenOK with default headers values
func NewGetNFTContractTokenOK() *GetNFTContractTokenOK {

	return &GetNFTContractTokenOK{}
}

// WithPayload adds the payload to the get n f t contract token o k response
func (o *GetNFTContractTokenOK) WithPayload(payload *models.NFTTokenRow) *GetNFTContractTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get n f t contract token o k response
func (o *GetNFTContractTokenOK) SetPayload(payload *models.NFTTokenRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNFTContractTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNFTContractTokenBadRequestCode is the HTTP code returned for type GetNFTContractTokenBadRequest
const GetNFTContractTokenBadRequestCode int = 400

/*GetNFTContractTokenBadRequest Bad request

swagger:response getNFTContractTokenBadRequest
*/
type GetNFTContractTokenBadRequest struct {
}

// NewGetNFTContractTokenBadRequest creates GetNFTContractTokenBadRequest with default headers values
func NewGetNFTContractTokenBadRequest() *GetNFTContractTokenBadRequest {

	return &GetNFTContractTokenBadRequest{}
}

// WriteResponse to the client
func (o *GetNFTContractTokenBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetNFTContractTokenNotFoundCode is the HTTP code returned for type GetNFTContractTokenNotFound
const GetNFTContractTokenNotFoundCode int = 404

/*GetNFTContractTokenNotFound Not Found

swagger:response getNFTContractTokenNotFound
*/
type GetNFTContractTokenNotFound struct {
}

// NewGetNFTContractTokenNotFound creates GetNFTContractTokenNotFound with default headers values
func NewGetNFTContractTokenNotFound() *GetNFTContractTokenNotFound {

	return &GetNFTContractTokenNotFound{}
}

// WriteResponse to the client
func (o *GetNFTContractTokenNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetNFTContractTokenInternalServerErrorCode is the HTTP code returned for type GetNFTContractTokenInternalServerError
const GetNFTContractTokenInternalServerErrorCode int = 500

/*GetNFTContractTokenInternalServerError Internal error

swagger:response getNFTContractTokenInternalServerError
*/
type GetNFTContractTokenInternalServerError struct {
}

// NewGetNFTContractTokenInternalServerError creates GetNFTContractTokenInternalServerError with default headers values
func NewGetNFTContractTokenInternalServerError() *GetNFTContractTokenInternalServerError {

	return &GetNFTContractTokenInternalServerError{}
}

// WriteResponse to the client
func (o *GetNFTContractTokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}