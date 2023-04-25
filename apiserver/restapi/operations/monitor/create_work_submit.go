// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// CreateWorkSubmitHandlerFunc turns a function with the right signature into a create work submit handler
type CreateWorkSubmitHandlerFunc func(CreateWorkSubmitParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateWorkSubmitHandlerFunc) Handle(params CreateWorkSubmitParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateWorkSubmitHandler interface for that can handle valid create work submit params
type CreateWorkSubmitHandler interface {
	Handle(CreateWorkSubmitParams, *v1.Principal) middleware.Responder
}

// NewCreateWorkSubmit creates a new http.Handler for the create work submit operation
func NewCreateWorkSubmit(ctx *middleware.Context, handler CreateWorkSubmitHandler) *CreateWorkSubmit {
	return &CreateWorkSubmit{Context: ctx, Handler: handler}
}

/* CreateWorkSubmit swagger:route POST /v1/workSubmit monitor createWorkSubmit

创建作业提交信息

创建作业提交信息

*/
type CreateWorkSubmit struct {
	Context *middleware.Context
	Handler CreateWorkSubmitHandler
}

func (o *CreateWorkSubmit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateWorkSubmitParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *v1.Principal
	if uprinc != nil {
		principal = uprinc.(*v1.Principal) // this is really a v1.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}