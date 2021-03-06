// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetProfileOKCode is the HTTP code returned for type GetProfileOK
const GetProfileOKCode int = 200

/*GetProfileOK Successful Operation

swagger:response getProfileOK
*/
type GetProfileOK struct {
}

// NewGetProfileOK creates GetProfileOK with default headers values
func NewGetProfileOK() *GetProfileOK {

	return &GetProfileOK{}
}

// WriteResponse to the client
func (o *GetProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetProfileNoContentCode is the HTTP code returned for type GetProfileNoContent
const GetProfileNoContentCode int = 204

/*GetProfileNoContent profile not found

swagger:response getProfileNoContent
*/
type GetProfileNoContent struct {
}

// NewGetProfileNoContent creates GetProfileNoContent with default headers values
func NewGetProfileNoContent() *GetProfileNoContent {

	return &GetProfileNoContent{}
}

// WriteResponse to the client
func (o *GetProfileNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetProfileInternalServerErrorCode is the HTTP code returned for type GetProfileInternalServerError
const GetProfileInternalServerErrorCode int = 500

/*GetProfileInternalServerError internal error

swagger:response getProfileInternalServerError
*/
type GetProfileInternalServerError struct {
}

// NewGetProfileInternalServerError creates GetProfileInternalServerError with default headers values
func NewGetProfileInternalServerError() *GetProfileInternalServerError {

	return &GetProfileInternalServerError{}
}

// WriteResponse to the client
func (o *GetProfileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
