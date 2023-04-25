// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetHomeWorkHandlerFunc turns a function with the right signature into a get home work handler
type GetHomeWorkHandlerFunc func(GetHomeWorkParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHomeWorkHandlerFunc) Handle(params GetHomeWorkParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetHomeWorkHandler interface for that can handle valid get home work params
type GetHomeWorkHandler interface {
	Handle(GetHomeWorkParams, *v1.Principal) middleware.Responder
}

// NewGetHomeWork creates a new http.Handler for the get home work operation
func NewGetHomeWork(ctx *middleware.Context, handler GetHomeWorkHandler) *GetHomeWork {
	return &GetHomeWork{Context: ctx, Handler: handler}
}

/* GetHomeWork swagger:route GET /v1/homeWork/{id} monitor getHomeWork

获取申报信息

获取申报信息

*/
type GetHomeWork struct {
	Context *middleware.Context
	Handler GetHomeWorkHandler
}

func (o *GetHomeWork) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHomeWorkParams()
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