// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetMeasureDetailsOKCode is the HTTP code returned for type GetMeasureDetailsOK
const GetMeasureDetailsOKCode int = 200

/*GetMeasureDetailsOK ok

swagger:response getMeasureDetailsOK
*/
type GetMeasureDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *v1.MeasureDetails `json:"body,omitempty"`
}

// NewGetMeasureDetailsOK creates GetMeasureDetailsOK with default headers values
func NewGetMeasureDetailsOK() *GetMeasureDetailsOK {

	return &GetMeasureDetailsOK{}
}

// WithPayload adds the payload to the get measure details o k response
func (o *GetMeasureDetailsOK) WithPayload(payload *v1.MeasureDetails) *GetMeasureDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get measure details o k response
func (o *GetMeasureDetailsOK) SetPayload(payload *v1.MeasureDetails) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeasureDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}