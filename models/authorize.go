// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Authorize authorize
//
// swagger:model authorize
type Authorize struct {

	// data
	Data string `json:"data,omitempty"`
}

// Validate validates this authorize
func (m *Authorize) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this authorize based on context it is used
func (m *Authorize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Authorize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Authorize) UnmarshalBinary(b []byte) error {
	var res Authorize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
