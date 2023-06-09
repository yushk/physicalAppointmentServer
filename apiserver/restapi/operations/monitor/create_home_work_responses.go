// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// CreateHomeWorkOKCode is the HTTP code returned for type CreateHomeWorkOK
const CreateHomeWorkOKCode int = 200

/*CreateHomeWorkOK ok

swagger:response createHomeWorkOK
*/
type CreateHomeWorkOK struct {

	/*
	  In: Body
	*/
	Payload *v1.HomeWork `json:"body,omitempty"`
}

// NewCreateHomeWorkOK creates CreateHomeWorkOK with default headers values
func NewCreateHomeWorkOK() *CreateHomeWorkOK {

	return &CreateHomeWorkOK{}
}

// WithPayload adds the payload to the create home work o k response
func (o *CreateHomeWorkOK) WithPayload(payload *v1.HomeWork) *CreateHomeWorkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create home work o k response
func (o *CreateHomeWorkOK) SetPayload(payload *v1.HomeWork) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHomeWorkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
