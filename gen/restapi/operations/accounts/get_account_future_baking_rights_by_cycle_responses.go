// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	models "github.com/everstake/teztracker/gen/models"
)

// GetAccountFutureBakingRightsByCycleOKCode is the HTTP code returned for type GetAccountFutureBakingRightsByCycleOK
const GetAccountFutureBakingRightsByCycleOKCode int = 200

/*GetAccountFutureBakingRightsByCycleOK Query compatibility endpoint for account future baking

swagger:response getAccountFutureBakingRightsByCycleOK
*/
type GetAccountFutureBakingRightsByCycleOK struct {
	/*The total number of data entries.

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.BakingRightsRow `json:"body,omitempty"`
}

// NewGetAccountFutureBakingRightsByCycleOK creates GetAccountFutureBakingRightsByCycleOK with default headers values
func NewGetAccountFutureBakingRightsByCycleOK() *GetAccountFutureBakingRightsByCycleOK {

	return &GetAccountFutureBakingRightsByCycleOK{}
}

// WithXTotalCount adds the xTotalCount to the get account future baking rights by cycle o k response
func (o *GetAccountFutureBakingRightsByCycleOK) WithXTotalCount(xTotalCount int64) *GetAccountFutureBakingRightsByCycleOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the get account future baking rights by cycle o k response
func (o *GetAccountFutureBakingRightsByCycleOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the get account future baking rights by cycle o k response
func (o *GetAccountFutureBakingRightsByCycleOK) WithPayload(payload []*models.BakingRightsRow) *GetAccountFutureBakingRightsByCycleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account future baking rights by cycle o k response
func (o *GetAccountFutureBakingRightsByCycleOK) SetPayload(payload []*models.BakingRightsRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountFutureBakingRightsByCycleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.BakingRightsRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAccountFutureBakingRightsByCycleBadRequestCode is the HTTP code returned for type GetAccountFutureBakingRightsByCycleBadRequest
const GetAccountFutureBakingRightsByCycleBadRequestCode int = 400

/*GetAccountFutureBakingRightsByCycleBadRequest Bad request

swagger:response getAccountFutureBakingRightsByCycleBadRequest
*/
type GetAccountFutureBakingRightsByCycleBadRequest struct {
}

// NewGetAccountFutureBakingRightsByCycleBadRequest creates GetAccountFutureBakingRightsByCycleBadRequest with default headers values
func NewGetAccountFutureBakingRightsByCycleBadRequest() *GetAccountFutureBakingRightsByCycleBadRequest {

	return &GetAccountFutureBakingRightsByCycleBadRequest{}
}

// WriteResponse to the client
func (o *GetAccountFutureBakingRightsByCycleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAccountFutureBakingRightsByCycleNotFoundCode is the HTTP code returned for type GetAccountFutureBakingRightsByCycleNotFound
const GetAccountFutureBakingRightsByCycleNotFoundCode int = 404

/*GetAccountFutureBakingRightsByCycleNotFound Not Found

swagger:response getAccountFutureBakingRightsByCycleNotFound
*/
type GetAccountFutureBakingRightsByCycleNotFound struct {
}

// NewGetAccountFutureBakingRightsByCycleNotFound creates GetAccountFutureBakingRightsByCycleNotFound with default headers values
func NewGetAccountFutureBakingRightsByCycleNotFound() *GetAccountFutureBakingRightsByCycleNotFound {

	return &GetAccountFutureBakingRightsByCycleNotFound{}
}

// WriteResponse to the client
func (o *GetAccountFutureBakingRightsByCycleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}