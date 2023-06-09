// Code generated by go-swagger; DO NOT EDIT.

package oauth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	v1 "github.com/yushk/sport_backend/apiserver/v1"
)

// TokenOKCode is the HTTP code returned for type TokenOK
const TokenOKCode int = 200

/*TokenOK A successful response.

swagger:response tokenOK
*/
type TokenOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Token `json:"body,omitempty"`
}

// NewTokenOK creates TokenOK with default headers values
func NewTokenOK() *TokenOK {

	return &TokenOK{}
}

// WithPayload adds the payload to the token o k response
func (o *TokenOK) WithPayload(payload *v1.Token) *TokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the token o k response
func (o *TokenOK) SetPayload(payload *v1.Token) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
