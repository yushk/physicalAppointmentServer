// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// CreateCaseDataHandlerFunc turns a function with the right signature into a create case data handler
type CreateCaseDataHandlerFunc func(CreateCaseDataParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCaseDataHandlerFunc) Handle(params CreateCaseDataParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateCaseDataHandler interface for that can handle valid create case data params
type CreateCaseDataHandler interface {
	Handle(CreateCaseDataParams, *v1.Principal) middleware.Responder
}

// NewCreateCaseData creates a new http.Handler for the create case data operation
func NewCreateCaseData(ctx *middleware.Context, handler CreateCaseDataHandler) *CreateCaseData {
	return &CreateCaseData{Context: ctx, Handler: handler}
}

/* CreateCaseData swagger:route POST /v1/caseData monitor createCaseData

创建病例数据信息

创建病例数据信息

*/
type CreateCaseData struct {
	Context *middleware.Context
	Handler CreateCaseDataHandler
}

func (o *CreateCaseData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateCaseDataParams()
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
