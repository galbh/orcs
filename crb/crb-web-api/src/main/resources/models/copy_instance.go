package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// CopyInstance copyURL
// swagger:model CopyInstance
type CopyInstance struct {

	// copy URL
	CopyURL string `json:"copyURL,omitempty"`
}

// Validate validates this copy instance
func (m *CopyInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
