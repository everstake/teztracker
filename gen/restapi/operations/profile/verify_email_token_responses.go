// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// VerifyEmailTokenOKCode is the HTTP code returned for type VerifyEmailTokenOK
const VerifyEmailTokenOKCode int = 200

/*VerifyEmailTokenOK Verified user email token

swagger:response verifyEmailTokenOK
*/
type VerifyEmailTokenOK struct {
}

// NewVerifyEmailTokenOK creates VerifyEmailTokenOK with default headers values
func NewVerifyEmailTokenOK() *VerifyEmailTokenOK {

	return &VerifyEmailTokenOK{}
}

// WriteResponse to the client
func (o *VerifyEmailTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// VerifyEmailTokenBadRequestCode is the HTTP code returned for type VerifyEmailTokenBadRequest
const VerifyEmailTokenBadRequestCode int = 400

/*VerifyEmailTokenBadRequest Bad request

swagger:response verifyEmailTokenBadRequest
*/
type VerifyEmailTokenBadRequest struct {
}

// NewVerifyEmailTokenBadRequest creates VerifyEmailTokenBadRequest with default headers values
func NewVerifyEmailTokenBadRequest() *VerifyEmailTokenBadRequest {

	return &VerifyEmailTokenBadRequest{}
}

// WriteResponse to the client
func (o *VerifyEmailTokenBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// VerifyEmailTokenInternalServerErrorCode is the HTTP code returned for type VerifyEmailTokenInternalServerError
const VerifyEmailTokenInternalServerErrorCode int = 500

/*VerifyEmailTokenInternalServerError Internal error

swagger:response verifyEmailTokenInternalServerError
*/
type VerifyEmailTokenInternalServerError struct {
}

// NewVerifyEmailTokenInternalServerError creates VerifyEmailTokenInternalServerError with default headers values
func NewVerifyEmailTokenInternalServerError() *VerifyEmailTokenInternalServerError {

	return &VerifyEmailTokenInternalServerError{}
}

// WriteResponse to the client
func (o *VerifyEmailTokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
