// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StoredBook stored book
// swagger:model StoredBook
type StoredBook struct {

	// ID
	// Required: true
	ID *int64 `json:"ID"`

	// book
	Book *Book `json:"book,omitempty"`
}

// Validate validates this stored book
func (m *StoredBook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoredBook) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *StoredBook) validateBook(formats strfmt.Registry) error {

	if swag.IsZero(m.Book) { // not required
		return nil
	}

	if m.Book != nil {
		if err := m.Book.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("book")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoredBook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoredBook) UnmarshalBinary(b []byte) error {
	var res StoredBook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
