// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Product product
// swagger:model Product
type Product struct {

	// available
	Available []string `json:"available"`

	// included
	Included []string `json:"included"`

	// name
	Name string `json:"name,omitempty"`

	// plans
	Plans []*Plan `json:"plans"`

	// pretty name
	PrettyName string `json:"prettyName,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIncluded(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlans(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(m.Available) { // not required
		return nil
	}

	return nil
}

func (m *Product) validateIncluded(formats strfmt.Registry) error {

	if swag.IsZero(m.Included) { // not required
		return nil
	}

	return nil
}

func (m *Product) validatePlans(formats strfmt.Registry) error {

	if swag.IsZero(m.Plans) { // not required
		return nil
	}

	for i := 0; i < len(m.Plans); i++ {

		if swag.IsZero(m.Plans[i]) { // not required
			continue
		}

		if m.Plans[i] != nil {

			if err := m.Plans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plans" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
