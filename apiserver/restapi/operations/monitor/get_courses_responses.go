// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// GetCoursesOKCode is the HTTP code returned for type GetCoursesOK
const GetCoursesOKCode int = 200

/*GetCoursesOK ok

swagger:response getCoursesOK
*/
type GetCoursesOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Courses `json:"body,omitempty"`
}

// NewGetCoursesOK creates GetCoursesOK with default headers values
func NewGetCoursesOK() *GetCoursesOK {

	return &GetCoursesOK{}
}

// WithPayload adds the payload to the get courses o k response
func (o *GetCoursesOK) WithPayload(payload *v1.Courses) *GetCoursesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get courses o k response
func (o *GetCoursesOK) SetPayload(payload *v1.Courses) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCoursesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
