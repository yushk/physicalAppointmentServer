// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// CreateCommentOKCode is the HTTP code returned for type CreateCommentOK
const CreateCommentOKCode int = 200

/*CreateCommentOK ok

swagger:response createCommentOK
*/
type CreateCommentOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Comment `json:"body,omitempty"`
}

// NewCreateCommentOK creates CreateCommentOK with default headers values
func NewCreateCommentOK() *CreateCommentOK {

	return &CreateCommentOK{}
}

// WithPayload adds the payload to the create comment o k response
func (o *CreateCommentOK) WithPayload(payload *v1.Comment) *CreateCommentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create comment o k response
func (o *CreateCommentOK) SetPayload(payload *v1.Comment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCommentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
