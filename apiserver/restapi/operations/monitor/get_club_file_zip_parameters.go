// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetClubFileZipParams creates a new GetClubFileZipParams object
//
// There are no default values defined in the spec.
func NewGetClubFileZipParams() GetClubFileZipParams {

	return GetClubFileZipParams{}
}

// GetClubFileZipParams contains all the bound params for the get club file zip operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetClubFileZip
type GetClubFileZipParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*俱乐部名称
	  In: query
	*/
	Club *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetClubFileZipParams() beforehand.
func (o *GetClubFileZipParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qClub, qhkClub, _ := qs.GetOK("club")
	if err := o.bindClub(qClub, qhkClub, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClub binds and validates parameter Club from query.
func (o *GetClubFileZipParams) bindClub(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Club = &raw

	return nil
}
