// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetWorkSubmitsHandlerFunc turns a function with the right signature into a get work submits handler
type GetWorkSubmitsHandlerFunc func(GetWorkSubmitsParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWorkSubmitsHandlerFunc) Handle(params GetWorkSubmitsParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetWorkSubmitsHandler interface for that can handle valid get work submits params
type GetWorkSubmitsHandler interface {
	Handle(GetWorkSubmitsParams, *v1.Principal) middleware.Responder
}

// NewGetWorkSubmits creates a new http.Handler for the get work submits operation
func NewGetWorkSubmits(ctx *middleware.Context, handler GetWorkSubmitsHandler) *GetWorkSubmits {
	return &GetWorkSubmits{Context: ctx, Handler: handler}
}

/* GetWorkSubmits swagger:route GET /v1/workSubmits monitor getWorkSubmits

获取作业提交信息列表

获取作业提交信息列表

*/
type GetWorkSubmits struct {
	Context *middleware.Context
	Handler GetWorkSubmitsHandler
}

func (o *GetWorkSubmits) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetWorkSubmitsParams()
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
