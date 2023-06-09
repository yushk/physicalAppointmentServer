// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// DeleteWorkHandlerFunc turns a function with the right signature into a delete work handler
type DeleteWorkHandlerFunc func(DeleteWorkParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteWorkHandlerFunc) Handle(params DeleteWorkParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteWorkHandler interface for that can handle valid delete work params
type DeleteWorkHandler interface {
	Handle(DeleteWorkParams, *v1.Principal) middleware.Responder
}

// NewDeleteWork creates a new http.Handler for the delete work operation
func NewDeleteWork(ctx *middleware.Context, handler DeleteWorkHandler) *DeleteWork {
	return &DeleteWork{Context: ctx, Handler: handler}
}

/* DeleteWork swagger:route DELETE /v1/work/{id} monitor deleteWork

删除作业信息

删除作业信息

*/
type DeleteWork struct {
	Context *middleware.Context
	Handler DeleteWorkHandler
}

func (o *DeleteWork) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteWorkParams()
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
