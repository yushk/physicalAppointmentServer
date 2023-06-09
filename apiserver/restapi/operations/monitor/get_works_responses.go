// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetWorksOKCode is the HTTP code returned for type GetWorksOK
const GetWorksOKCode int = 200

/*GetWorksOK ok

swagger:response getWorksOK
*/
type GetWorksOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Works `json:"body,omitempty"`
}

// NewGetWorksOK creates GetWorksOK with default headers values
func NewGetWorksOK() *GetWorksOK {

	return &GetWorksOK{}
}

// WithPayload adds the payload to the get works o k response
func (o *GetWorksOK) WithPayload(payload *v1.Works) *GetWorksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get works o k response
func (o *GetWorksOK) SetPayload(payload *v1.Works) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWorksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
