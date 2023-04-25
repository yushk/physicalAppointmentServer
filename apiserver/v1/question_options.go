// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuestionOptions 答案选项
//
// swagger:model QuestionOptions
type QuestionOptions struct {

	// 问题描述
	Desc string `json:"desc"`

	// 唯一标识
	ID string `json:"id"`

	// 得分
	Score string `json:"score"`
}

// Validate validates this question options
func (m *QuestionOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this question options based on context it is used
func (m *QuestionOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuestionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuestionOptions) UnmarshalBinary(b []byte) error {
	var res QuestionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
