// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// DeleteCourseHandlerFunc turns a function with the right signature into a delete course handler
type DeleteCourseHandlerFunc func(DeleteCourseParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCourseHandlerFunc) Handle(params DeleteCourseParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteCourseHandler interface for that can handle valid delete course params
type DeleteCourseHandler interface {
	Handle(DeleteCourseParams, *v1.Principal) middleware.Responder
}

// NewDeleteCourse creates a new http.Handler for the delete course operation
func NewDeleteCourse(ctx *middleware.Context, handler DeleteCourseHandler) *DeleteCourse {
	return &DeleteCourse{Context: ctx, Handler: handler}
}

/* DeleteCourse swagger:route DELETE /v1/course/{id} monitor deleteCourse

删除课程信息

删除课程信息

*/
type DeleteCourse struct {
	Context *middleware.Context
	Handler DeleteCourseHandler
}

func (o *DeleteCourse) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteCourseParams()
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