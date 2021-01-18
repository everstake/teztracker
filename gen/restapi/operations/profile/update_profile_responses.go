// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateProfileOKCode is the HTTP code returned for type UpdateProfileOK
const UpdateProfileOKCode int = 200

/*UpdateProfileOK Update user profile

swagger:response updateProfileOK
*/
type UpdateProfileOK struct {
}

// NewUpdateProfileOK creates UpdateProfileOK with default headers values
func NewUpdateProfileOK() *UpdateProfileOK {

	return &UpdateProfileOK{}
}

// WriteResponse to the client
func (o *UpdateProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateProfileBadRequestCode is the HTTP code returned for type UpdateProfileBadRequest
const UpdateProfileBadRequestCode int = 400

/*UpdateProfileBadRequest Bad request

swagger:response updateProfileBadRequest
*/
type UpdateProfileBadRequest struct {
}

// NewUpdateProfileBadRequest creates UpdateProfileBadRequest with default headers values
func NewUpdateProfileBadRequest() *UpdateProfileBadRequest {

	return &UpdateProfileBadRequest{}
}

// WriteResponse to the client
func (o *UpdateProfileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateProfileInternalServerErrorCode is the HTTP code returned for type UpdateProfileInternalServerError
const UpdateProfileInternalServerErrorCode int = 500

/*UpdateProfileInternalServerError Internal error

swagger:response updateProfileInternalServerError
*/
type UpdateProfileInternalServerError struct {
}

// NewUpdateProfileInternalServerError creates UpdateProfileInternalServerError with default headers values
func NewUpdateProfileInternalServerError() *UpdateProfileInternalServerError {

	return &UpdateProfileInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateProfileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}