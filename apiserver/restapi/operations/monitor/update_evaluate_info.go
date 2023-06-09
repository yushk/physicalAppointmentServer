// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// UpdateEvaluateInfoHandlerFunc turns a function with the right signature into a update evaluate info handler
type UpdateEvaluateInfoHandlerFunc func(UpdateEvaluateInfoParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateEvaluateInfoHandlerFunc) Handle(params UpdateEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateEvaluateInfoHandler interface for that can handle valid update evaluate info params
type UpdateEvaluateInfoHandler interface {
	Handle(UpdateEvaluateInfoParams, *v1.Principal) middleware.Responder
}

// NewUpdateEvaluateInfo creates a new http.Handler for the update evaluate info operation
func NewUpdateEvaluateInfo(ctx *middleware.Context, handler UpdateEvaluateInfoHandler) *UpdateEvaluateInfo {
	return &UpdateEvaluateInfo{Context: ctx, Handler: handler}
}

/* UpdateEvaluateInfo swagger:route PUT /v1/evaluateInfo/{id} monitor updateEvaluateInfo

编辑评价信息

编辑评价信息

*/
type UpdateEvaluateInfo struct {
	Context *middleware.Context
	Handler UpdateEvaluateInfoHandler
}

func (o *UpdateEvaluateInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateEvaluateInfoParams()
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
