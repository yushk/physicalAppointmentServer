// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HomeWork 申报申请
//
// swagger:model HomeWork
type HomeWork struct {

	// 申请信息填报结果
	ApplyItems []*ApplyItem `json:"apply_items"`

	// 申请人id
	ApplyUserID string `json:"apply_user_id"`

	// 俱乐部名称
	ClubName string `json:"club_name"`

	// 创建时间
	Create int64 `json:"create"`

	// id
	// Read Only: true
	ID string `json:"id"`

	// 申报状态
	Status string `json:"status"`

	// 更新时间
	Update int64 `json:"update"`
}

// Validate validates this home work
func (m *HomeWork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplyItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HomeWork) validateApplyItems(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplyItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ApplyItems); i++ {
		if swag.IsZero(m.ApplyItems[i]) { // not required
			continue
		}

		if m.ApplyItems[i] != nil {
			if err := m.ApplyItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apply_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this home work based on the context it is used
func (m *HomeWork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplyItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HomeWork) contextValidateApplyItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ApplyItems); i++ {

		if m.ApplyItems[i] != nil {
			if err := m.ApplyItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apply_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HomeWork) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HomeWork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HomeWork) UnmarshalBinary(b []byte) error {
	var res HomeWork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}