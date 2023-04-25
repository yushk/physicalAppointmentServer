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
	"github.com/go-openapi/swag"
)

// NewGetSolutionsParams creates a new GetSolutionsParams object
//
// There are no default values defined in the spec.
func NewGetSolutionsParams() GetSolutionsParams {

	return GetSolutionsParams{}
}

// GetSolutionsParams contains all the bound params for the get solutions operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSolutions
type GetSolutionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*限制条目数
	  In: query
	*/
	Limit *int64
	/*筛选条件
	  In: query
	*/
	Query *string
	/*翻过条目数
	  In: query
	*/
	Skip *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSolutionsParams() beforehand.
func (o *GetSolutionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qQuery, qhkQuery, _ := qs.GetOK("query")
	if err := o.bindQuery(qQuery, qhkQuery, route.Formats); err != nil {
		res = append(res, err)
	}

	qSkip, qhkSkip, _ := qs.GetOK("skip")
	if err := o.bindSkip(qSkip, qhkSkip, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetSolutionsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindQuery binds and validates parameter Query from query.
func (o *GetSolutionsParams) bindQuery(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Query = &raw

	return nil
}

// bindSkip binds and validates parameter Skip from query.
func (o *GetSolutionsParams) bindSkip(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("skip", "query", "int64", raw)
	}
	o.Skip = &value

	return nil
}
