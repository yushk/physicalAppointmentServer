// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UploadFileOKCode is the HTTP code returned for type UploadFileOK
const UploadFileOKCode int = 200

/*UploadFileOK A successful response.

swagger:response uploadFileOK
*/
type UploadFileOK struct {

	/*存储的文件ID
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewUploadFileOK creates UploadFileOK with default headers values
func NewUploadFileOK() *UploadFileOK {

	return &UploadFileOK{}
}

// WithPayload adds the payload to the upload file o k response
func (o *UploadFileOK) WithPayload(payload string) *UploadFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload file o k response
func (o *UploadFileOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
