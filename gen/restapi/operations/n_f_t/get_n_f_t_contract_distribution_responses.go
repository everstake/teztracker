// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetNFTContractDistributionOKCode is the HTTP code returned for type GetNFTContractDistributionOK
const GetNFTContractDistributionOKCode int = 200

/*GetNFTContractDistributionOK Query compatibility endpoint for NFT contract

swagger:response getNFTContractDistributionOK
*/
type GetNFTContractDistributionOK struct {

	/*
	  In: Body
	*/
	Payload *models.NFTContractDistribution `json:"body,omitempty"`
}

// NewGetNFTContractDistributionOK creates GetNFTContractDistributionOK with default headers values
func NewGetNFTContractDistributionOK() *GetNFTContractDistributionOK {

	return &GetNFTContractDistributionOK{}
}

// WithPayload adds the payload to the get n f t contract distribution o k response
func (o *GetNFTContractDistributionOK) WithPayload(payload *models.NFTContractDistribution) *GetNFTContractDistributionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get n f t contract distribution o k response
func (o *GetNFTContractDistributionOK) SetPayload(payload *models.NFTContractDistribution) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNFTContractDistributionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNFTContractDistributionBadRequestCode is the HTTP code returned for type GetNFTContractDistributionBadRequest
const GetNFTContractDistributionBadRequestCode int = 400

/*GetNFTContractDistributionBadRequest Bad request

swagger:response getNFTContractDistributionBadRequest
*/
type GetNFTContractDistributionBadRequest struct {
}

// NewGetNFTContractDistributionBadRequest creates GetNFTContractDistributionBadRequest with default headers values
func NewGetNFTContractDistributionBadRequest() *GetNFTContractDistributionBadRequest {

	return &GetNFTContractDistributionBadRequest{}
}

// WriteResponse to the client
func (o *GetNFTContractDistributionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetNFTContractDistributionNotFoundCode is the HTTP code returned for type GetNFTContractDistributionNotFound
const GetNFTContractDistributionNotFoundCode int = 404

/*GetNFTContractDistributionNotFound Not Found

swagger:response getNFTContractDistributionNotFound
*/
type GetNFTContractDistributionNotFound struct {
}

// NewGetNFTContractDistributionNotFound creates GetNFTContractDistributionNotFound with default headers values
func NewGetNFTContractDistributionNotFound() *GetNFTContractDistributionNotFound {

	return &GetNFTContractDistributionNotFound{}
}

// WriteResponse to the client
func (o *GetNFTContractDistributionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetNFTContractDistributionInternalServerErrorCode is the HTTP code returned for type GetNFTContractDistributionInternalServerError
const GetNFTContractDistributionInternalServerErrorCode int = 500

/*GetNFTContractDistributionInternalServerError Internal error

swagger:response getNFTContractDistributionInternalServerError
*/
type GetNFTContractDistributionInternalServerError struct {
}

// NewGetNFTContractDistributionInternalServerError creates GetNFTContractDistributionInternalServerError with default headers values
func NewGetNFTContractDistributionInternalServerError() *GetNFTContractDistributionInternalServerError {

	return &GetNFTContractDistributionInternalServerError{}
}

// WriteResponse to the client
func (o *GetNFTContractDistributionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}