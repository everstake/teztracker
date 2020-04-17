// Code generated by go-swagger; DO NOT EDIT.

package operations_list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	models "github.com/everstake/teztracker/gen/models"
)

// GetDoubleEndorsementsListOKCode is the HTTP code returned for type GetDoubleEndorsementsListOK
const GetDoubleEndorsementsListOKCode int = 200

/*GetDoubleEndorsementsListOK Query compatibility endpoint for operations

swagger:response getDoubleEndorsementsListOK
*/
type GetDoubleEndorsementsListOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.OperationsRow `json:"body,omitempty"`
}

// NewGetDoubleEndorsementsListOK creates GetDoubleEndorsementsListOK with default headers values
func NewGetDoubleEndorsementsListOK() *GetDoubleEndorsementsListOK {

	return &GetDoubleEndorsementsListOK{}
}

// WithXTotalCount adds the xTotalCount to the get double endorsements list o k response
func (o *GetDoubleEndorsementsListOK) WithXTotalCount(xTotalCount int64) *GetDoubleEndorsementsListOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get double endorsements list o k response
func (o *GetDoubleEndorsementsListOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get double endorsements list o k response
func (o *GetDoubleEndorsementsListOK) WithPayload(payload []*models.OperationsRow) *GetDoubleEndorsementsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get double endorsements list o k response
func (o *GetDoubleEndorsementsListOK) SetPayload(payload []*models.OperationsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDoubleEndorsementsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.OperationsRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDoubleEndorsementsListBadRequestCode is the HTTP code returned for type GetDoubleEndorsementsListBadRequest
const GetDoubleEndorsementsListBadRequestCode int = 400

/*GetDoubleEndorsementsListBadRequest Bad request

swagger:response getDoubleEndorsementsListBadRequest
*/
type GetDoubleEndorsementsListBadRequest struct {
}

// NewGetDoubleEndorsementsListBadRequest creates GetDoubleEndorsementsListBadRequest with default headers values
func NewGetDoubleEndorsementsListBadRequest() *GetDoubleEndorsementsListBadRequest {

	return &GetDoubleEndorsementsListBadRequest{}
}

// WriteResponse to the client
func (o *GetDoubleEndorsementsListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetDoubleEndorsementsListNotFoundCode is the HTTP code returned for type GetDoubleEndorsementsListNotFound
const GetDoubleEndorsementsListNotFoundCode int = 404

/*GetDoubleEndorsementsListNotFound Not Found

swagger:response getDoubleEndorsementsListNotFound
*/
type GetDoubleEndorsementsListNotFound struct {
}

// NewGetDoubleEndorsementsListNotFound creates GetDoubleEndorsementsListNotFound with default headers values
func NewGetDoubleEndorsementsListNotFound() *GetDoubleEndorsementsListNotFound {

	return &GetDoubleEndorsementsListNotFound{}
}

// WriteResponse to the client
func (o *GetDoubleEndorsementsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
