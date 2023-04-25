// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetPersonalsOKCode is the HTTP code returned for type GetPersonalsOK
const GetPersonalsOKCode int = 200

/*GetPersonalsOK ok

swagger:response getPersonalsOK
*/
type GetPersonalsOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Personals `json:"body,omitempty"`
}

// NewGetPersonalsOK creates GetPersonalsOK with default headers values
func NewGetPersonalsOK() *GetPersonalsOK {

	return &GetPersonalsOK{}
}

// WithPayload adds the payload to the get personals o k response
func (o *GetPersonalsOK) WithPayload(payload *v1.Personals) *GetPersonalsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get personals o k response
func (o *GetPersonalsOK) SetPayload(payload *v1.Personals) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPersonalsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
