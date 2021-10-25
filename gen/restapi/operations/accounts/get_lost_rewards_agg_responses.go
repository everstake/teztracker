// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetLostRewardsAggOKCode is the HTTP code returned for type GetLostRewardsAggOK
const GetLostRewardsAggOKCode int = 200

/*GetLostRewardsAggOK Get amount of lost rewards agg by period

swagger:response getLostRewardsAggOK
*/
type GetLostRewardsAggOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AggTimeInt `json:"body,omitempty"`
}

// NewGetLostRewardsAggOK creates GetLostRewardsAggOK with default headers values
func NewGetLostRewardsAggOK() *GetLostRewardsAggOK {

	return &GetLostRewardsAggOK{}
}

// WithPayload adds the payload to the get lost rewards agg o k response
func (o *GetLostRewardsAggOK) WithPayload(payload []*models.AggTimeInt) *GetLostRewardsAggOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lost rewards agg o k response
func (o *GetLostRewardsAggOK) SetPayload(payload []*models.AggTimeInt) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLostRewardsAggOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AggTimeInt, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetLostRewardsAggBadRequestCode is the HTTP code returned for type GetLostRewardsAggBadRequest
const GetLostRewardsAggBadRequestCode int = 400

/*GetLostRewardsAggBadRequest Bad request

swagger:response getLostRewardsAggBadRequest
*/
type GetLostRewardsAggBadRequest struct {
}

// NewGetLostRewardsAggBadRequest creates GetLostRewardsAggBadRequest with default headers values
func NewGetLostRewardsAggBadRequest() *GetLostRewardsAggBadRequest {

	return &GetLostRewardsAggBadRequest{}
}

// WriteResponse to the client
func (o *GetLostRewardsAggBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetLostRewardsAggNotFoundCode is the HTTP code returned for type GetLostRewardsAggNotFound
const GetLostRewardsAggNotFoundCode int = 404

/*GetLostRewardsAggNotFound Not Found

swagger:response getLostRewardsAggNotFound
*/
type GetLostRewardsAggNotFound struct {
}

// NewGetLostRewardsAggNotFound creates GetLostRewardsAggNotFound with default headers values
func NewGetLostRewardsAggNotFound() *GetLostRewardsAggNotFound {

	return &GetLostRewardsAggNotFound{}
}

// WriteResponse to the client
func (o *GetLostRewardsAggNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetLostRewardsAggInternalServerErrorCode is the HTTP code returned for type GetLostRewardsAggInternalServerError
const GetLostRewardsAggInternalServerErrorCode int = 500

/*GetLostRewardsAggInternalServerError Internal error

swagger:response getLostRewardsAggInternalServerError
*/
type GetLostRewardsAggInternalServerError struct {
}

// NewGetLostRewardsAggInternalServerError creates GetLostRewardsAggInternalServerError with default headers values
func NewGetLostRewardsAggInternalServerError() *GetLostRewardsAggInternalServerError {

	return &GetLostRewardsAggInternalServerError{}
}

// WriteResponse to the client
func (o *GetLostRewardsAggInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}