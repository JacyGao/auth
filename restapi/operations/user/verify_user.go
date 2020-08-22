// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VerifyUserHandlerFunc turns a function with the right signature into a verify user handler
type VerifyUserHandlerFunc func(VerifyUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyUserHandlerFunc) Handle(params VerifyUserParams) middleware.Responder {
	return fn(params)
}

// VerifyUserHandler interface for that can handle valid verify user params
type VerifyUserHandler interface {
	Handle(VerifyUserParams) middleware.Responder
}

// NewVerifyUser creates a new http.Handler for the verify user operation
func NewVerifyUser(ctx *middleware.Context, handler VerifyUserHandler) *VerifyUser {
	return &VerifyUser{Context: ctx, Handler: handler}
}

/*VerifyUser swagger:route PUT /user/verify user verifyUser

Verify a user

This can only be done by the logged in user.

*/
type VerifyUser struct {
	Context *middleware.Context
	Handler VerifyUserHandler
}

func (o *VerifyUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}