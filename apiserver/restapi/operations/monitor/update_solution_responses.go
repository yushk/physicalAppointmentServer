// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// UpdateSolutionOKCode is the HTTP code returned for type UpdateSolutionOK
const UpdateSolutionOKCode int = 200

/*UpdateSolutionOK A successful response.

swagger:response updateSolutionOK
*/
type UpdateSolutionOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Solution `json:"body,omitempty"`
}

// NewUpdateSolutionOK creates UpdateSolutionOK with default headers values
func NewUpdateSolutionOK() *UpdateSolutionOK {

	return &UpdateSolutionOK{}
}

// WithPayload adds the payload to the update solution o k response
func (o *UpdateSolutionOK) WithPayload(payload *v1.Solution) *UpdateSolutionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update solution o k response
func (o *UpdateSolutionOK) SetPayload(payload *v1.Solution) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSolutionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
