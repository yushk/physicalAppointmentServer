// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetMoveDataHandlerFunc turns a function with the right signature into a get move data handler
type GetMoveDataHandlerFunc func(GetMoveDataParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMoveDataHandlerFunc) Handle(params GetMoveDataParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetMoveDataHandler interface for that can handle valid get move data params
type GetMoveDataHandler interface {
	Handle(GetMoveDataParams, *v1.Principal) middleware.Responder
}

// NewGetMoveData creates a new http.Handler for the get move data operation
func NewGetMoveData(ctx *middleware.Context, handler GetMoveDataHandler) *GetMoveData {
	return &GetMoveData{Context: ctx, Handler: handler}
}

/* GetMoveData swagger:route GET /v1/moveData/{id} monitor getMoveData

获取运动数据信息

获取运动数据信息

*/
type GetMoveData struct {
	Context *middleware.Context
	Handler GetMoveDataHandler
}

func (o *GetMoveData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMoveDataParams()
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
