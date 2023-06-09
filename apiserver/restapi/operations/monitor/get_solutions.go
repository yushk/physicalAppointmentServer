// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetSolutionsHandlerFunc turns a function with the right signature into a get solutions handler
type GetSolutionsHandlerFunc func(GetSolutionsParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSolutionsHandlerFunc) Handle(params GetSolutionsParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetSolutionsHandler interface for that can handle valid get solutions params
type GetSolutionsHandler interface {
	Handle(GetSolutionsParams, *v1.Principal) middleware.Responder
}

// NewGetSolutions creates a new http.Handler for the get solutions operation
func NewGetSolutions(ctx *middleware.Context, handler GetSolutionsHandler) *GetSolutions {
	return &GetSolutions{Context: ctx, Handler: handler}
}

/* GetSolutions swagger:route GET /v1/solutions monitor getSolutions

获取学生提交作业信息列表

获取学生提交作业信息列表

*/
type GetSolutions struct {
	Context *middleware.Context
	Handler GetSolutionsHandler
}

func (o *GetSolutions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSolutionsParams()
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
