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

// GetDoubleBakingsListOKCode is the HTTP code returned for type GetDoubleBakingsListOK
const GetDoubleBakingsListOKCode int = 200

/*GetDoubleBakingsListOK Query compatibility endpoint for operations

swagger:response getDoubleBakingsListOK
*/
type GetDoubleBakingsListOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.OperationsRow `json:"body,omitempty"`
}

// NewGetDoubleBakingsListOK creates GetDoubleBakingsListOK with default headers values
func NewGetDoubleBakingsListOK() *GetDoubleBakingsListOK {

	return &GetDoubleBakingsListOK{}
}

// WithXTotalCount adds the xTotalCount to the get double bakings list o k response
func (o *GetDoubleBakingsListOK) WithXTotalCount(xTotalCount int64) *GetDoubleBakingsListOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get double bakings list o k response
func (o *GetDoubleBakingsListOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get double bakings list o k response
func (o *GetDoubleBakingsListOK) WithPayload(payload []*models.OperationsRow) *GetDoubleBakingsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get double bakings list o k response
func (o *GetDoubleBakingsListOK) SetPayload(payload []*models.OperationsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDoubleBakingsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetDoubleBakingsListBadRequestCode is the HTTP code returned for type GetDoubleBakingsListBadRequest
const GetDoubleBakingsListBadRequestCode int = 400

/*GetDoubleBakingsListBadRequest Bad request

swagger:response getDoubleBakingsListBadRequest
*/
type GetDoubleBakingsListBadRequest struct {
}

// NewGetDoubleBakingsListBadRequest creates GetDoubleBakingsListBadRequest with default headers values
func NewGetDoubleBakingsListBadRequest() *GetDoubleBakingsListBadRequest {

	return &GetDoubleBakingsListBadRequest{}
}

// WriteResponse to the client
func (o *GetDoubleBakingsListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetDoubleBakingsListNotFoundCode is the HTTP code returned for type GetDoubleBakingsListNotFound
const GetDoubleBakingsListNotFoundCode int = 404

/*GetDoubleBakingsListNotFound Not Found

swagger:response getDoubleBakingsListNotFound
*/
type GetDoubleBakingsListNotFound struct {
}

// NewGetDoubleBakingsListNotFound creates GetDoubleBakingsListNotFound with default headers values
func NewGetDoubleBakingsListNotFound() *GetDoubleBakingsListNotFound {

	return &GetDoubleBakingsListNotFound{}
}

// WriteResponse to the client
func (o *GetDoubleBakingsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
