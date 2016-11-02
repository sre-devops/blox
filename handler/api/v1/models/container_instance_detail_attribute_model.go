package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ContainerInstanceDetailAttributeModel container instance detail attribute model
// swagger:model ContainerInstanceDetailAttributeModel
type ContainerInstanceDetailAttributeModel struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this container instance detail attribute model
func (m *ContainerInstanceDetailAttributeModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerInstanceDetailAttributeModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstanceDetailAttributeModel) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}
