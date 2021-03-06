// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ForgetPasswordOKCode is the HTTP code returned for type ForgetPasswordOK
const ForgetPasswordOKCode int = 200

/*ForgetPasswordOK successful operation

swagger:response forgetPasswordOK
*/
type ForgetPasswordOK struct {
}

// NewForgetPasswordOK creates ForgetPasswordOK with default headers values
func NewForgetPasswordOK() *ForgetPasswordOK {

	return &ForgetPasswordOK{}
}

// WriteResponse to the client
func (o *ForgetPasswordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ForgetPasswordBadRequestCode is the HTTP code returned for type ForgetPasswordBadRequest
const ForgetPasswordBadRequestCode int = 400

/*ForgetPasswordBadRequest bad request

swagger:response forgetPasswordBadRequest
*/
type ForgetPasswordBadRequest struct {
}

// NewForgetPasswordBadRequest creates ForgetPasswordBadRequest with default headers values
func NewForgetPasswordBadRequest() *ForgetPasswordBadRequest {

	return &ForgetPasswordBadRequest{}
}

// WriteResponse to the client
func (o *ForgetPasswordBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ForgetPasswordInternalServerErrorCode is the HTTP code returned for type ForgetPasswordInternalServerError
const ForgetPasswordInternalServerErrorCode int = 500

/*ForgetPasswordInternalServerError internal error

swagger:response forgetPasswordInternalServerError
*/
type ForgetPasswordInternalServerError struct {
}

// NewForgetPasswordInternalServerError creates ForgetPasswordInternalServerError with default headers values
func NewForgetPasswordInternalServerError() *ForgetPasswordInternalServerError {

	return &ForgetPasswordInternalServerError{}
}

// WriteResponse to the client
func (o *ForgetPasswordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
