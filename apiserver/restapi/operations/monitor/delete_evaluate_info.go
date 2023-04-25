// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// DeleteEvaluateInfoHandlerFunc turns a function with the right signature into a delete evaluate info handler
type DeleteEvaluateInfoHandlerFunc func(DeleteEvaluateInfoParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteEvaluateInfoHandlerFunc) Handle(params DeleteEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteEvaluateInfoHandler interface for that can handle valid delete evaluate info params
type DeleteEvaluateInfoHandler interface {
	Handle(DeleteEvaluateInfoParams, *v1.Principal) middleware.Responder
}

// NewDeleteEvaluateInfo creates a new http.Handler for the delete evaluate info operation
func NewDeleteEvaluateInfo(ctx *middleware.Context, handler DeleteEvaluateInfoHandler) *DeleteEvaluateInfo {
	return &DeleteEvaluateInfo{Context: ctx, Handler: handler}
}

/* DeleteEvaluateInfo swagger:route DELETE /v1/evaluateInfo/{id} monitor deleteEvaluateInfo

删除评价信息

删除评价信息

*/
type DeleteEvaluateInfo struct {
	Context *middleware.Context
	Handler DeleteEvaluateInfoHandler
}

func (o *DeleteEvaluateInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteEvaluateInfoParams()
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
