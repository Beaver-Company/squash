// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DebugRequestStatus debug request status
// swagger:model DebugRequestStatus
type DebugRequestStatus struct {

	// debug attachment ref
	DebugAttachmentRef string `json:"debug_attachment_ref,omitempty"`
}

// Validate validates this debug request status
func (m *DebugRequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DebugRequestStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugRequestStatus) UnmarshalBinary(b []byte) error {
	var res DebugRequestStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
