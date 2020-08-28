// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/everstake/teztracker/gen/models"
)

// GetAssetsListOKCode is the HTTP code returned for type GetAssetsListOK
const GetAssetsListOKCode int = 200

/*GetAssetsListOK Query compatibility endpoint for tokens list

swagger:response getAssetsListOK
*/
type GetAssetsListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TokenAssetRow `json:"body,omitempty"`
}

// NewGetAssetsListOK creates GetAssetsListOK with default headers values
func NewGetAssetsListOK() *GetAssetsListOK {

	return &GetAssetsListOK{}
}

// WithPayload adds the payload to the get assets list o k response
func (o *GetAssetsListOK) WithPayload(payload []*models.TokenAssetRow) *GetAssetsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get assets list o k response
func (o *GetAssetsListOK) SetPayload(payload []*models.TokenAssetRow) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAssetsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.TokenAssetRow, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAssetsListBadRequestCode is the HTTP code returned for type GetAssetsListBadRequest
const GetAssetsListBadRequestCode int = 400

/*GetAssetsListBadRequest Bad request

swagger:response getAssetsListBadRequest
*/
type GetAssetsListBadRequest struct {
}

// NewGetAssetsListBadRequest creates GetAssetsListBadRequest with default headers values
func NewGetAssetsListBadRequest() *GetAssetsListBadRequest {

	return &GetAssetsListBadRequest{}
}

// WriteResponse to the client
func (o *GetAssetsListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetAssetsListNotFoundCode is the HTTP code returned for type GetAssetsListNotFound
const GetAssetsListNotFoundCode int = 404

/*GetAssetsListNotFound Not Found

swagger:response getAssetsListNotFound
*/
type GetAssetsListNotFound struct {
}

// NewGetAssetsListNotFound creates GetAssetsListNotFound with default headers values
func NewGetAssetsListNotFound() *GetAssetsListNotFound {

	return &GetAssetsListNotFound{}
}

// WriteResponse to the client
func (o *GetAssetsListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
