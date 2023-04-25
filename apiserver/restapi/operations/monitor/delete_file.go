// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// DeleteFileHandlerFunc turns a function with the right signature into a delete file handler
type DeleteFileHandlerFunc func(DeleteFileParams, *v1.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFileHandlerFunc) Handle(params DeleteFileParams, principal *v1.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteFileHandler interface for that can handle valid delete file params
type DeleteFileHandler interface {
	Handle(DeleteFileParams, *v1.Principal) middleware.Responder
}

// NewDeleteFile creates a new http.Handler for the delete file operation
func NewDeleteFile(ctx *middleware.Context, handler DeleteFileHandler) *DeleteFile {
	return &DeleteFile{Context: ctx, Handler: handler}
}

/* DeleteFile swagger:route DELETE /v1/file/delete/{id} monitor deleteFile

删除文件

删除文件

*/
type DeleteFile struct {
	Context *middleware.Context
	Handler DeleteFileHandler
}

func (o *DeleteFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteFileParams()
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
