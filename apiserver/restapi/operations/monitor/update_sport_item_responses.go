// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// UpdateSportItemOKCode is the HTTP code returned for type UpdateSportItemOK
const UpdateSportItemOKCode int = 200

/*UpdateSportItemOK A successful response.

swagger:response updateSportItemOK
*/
type UpdateSportItemOK struct {

	/*
	  In: Body
	*/
	Payload *v1.SportItem `json:"body,omitempty"`
}

// NewUpdateSportItemOK creates UpdateSportItemOK with default headers values
func NewUpdateSportItemOK() *UpdateSportItemOK {

	return &UpdateSportItemOK{}
}

// WithPayload adds the payload to the update sport item o k response
func (o *UpdateSportItemOK) WithPayload(payload *v1.SportItem) *UpdateSportItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update sport item o k response
func (o *UpdateSportItemOK) SetPayload(payload *v1.SportItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSportItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
