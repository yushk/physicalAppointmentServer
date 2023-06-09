// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// CreateCaseDataOKCode is the HTTP code returned for type CreateCaseDataOK
const CreateCaseDataOKCode int = 200

/*CreateCaseDataOK ok

swagger:response createCaseDataOK
*/
type CreateCaseDataOK struct {

	/*
	  In: Body
	*/
	Payload *v1.CaseData `json:"body,omitempty"`
}

// NewCreateCaseDataOK creates CreateCaseDataOK with default headers values
func NewCreateCaseDataOK() *CreateCaseDataOK {

	return &CreateCaseDataOK{}
}

// WithPayload adds the payload to the create case data o k response
func (o *CreateCaseDataOK) WithPayload(payload *v1.CaseData) *CreateCaseDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create case data o k response
func (o *CreateCaseDataOK) SetPayload(payload *v1.CaseData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCaseDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
