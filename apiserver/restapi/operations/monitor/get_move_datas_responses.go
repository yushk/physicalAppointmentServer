// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetMoveDatasOKCode is the HTTP code returned for type GetMoveDatasOK
const GetMoveDatasOKCode int = 200

/*GetMoveDatasOK ok

swagger:response getMoveDatasOK
*/
type GetMoveDatasOK struct {

	/*
	  In: Body
	*/
	Payload *v1.MoveDatas `json:"body,omitempty"`
}

// NewGetMoveDatasOK creates GetMoveDatasOK with default headers values
func NewGetMoveDatasOK() *GetMoveDatasOK {

	return &GetMoveDatasOK{}
}

// WithPayload adds the payload to the get move datas o k response
func (o *GetMoveDatasOK) WithPayload(payload *v1.MoveDatas) *GetMoveDatasOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move datas o k response
func (o *GetMoveDatasOK) SetPayload(payload *v1.MoveDatas) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveDatasOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
