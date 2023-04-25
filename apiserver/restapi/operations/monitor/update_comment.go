// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// UpdateCommentHandlerFunc turns a function with the right signature into a update comment handler
type UpdateCommentHandlerFunc func(UpdateCommentParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCommentHandlerFunc) Handle(params UpdateCommentParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateCommentHandler interface for that can handle valid update comment params
type UpdateCommentHandler interface {
	Handle(UpdateCommentParams, *v1.Principal) middleware.Responder
}

// NewUpdateComment creates a new http.Handler for the update comment operation
func NewUpdateComment(ctx *middleware.Context, handler UpdateCommentHandler) *UpdateComment {
	return &UpdateComment{Context: ctx, Handler: handler}
}

/* UpdateComment swagger:route PUT /v1/comment/{id} monitor updateComment

编辑学生评论信息

编辑评论信息

*/
type UpdateComment struct {
	Context *middleware.Context
	Handler UpdateCommentHandler
}

func (o *UpdateComment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateCommentParams()
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