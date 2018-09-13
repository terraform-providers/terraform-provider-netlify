// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountType account type
// swagger:model accountType
type AccountType struct {

	// capabilities
	Capabilities interface{} `json:"capabilities,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// monthly dollar price
	MonthlyDollarPrice int64 `json:"monthly_dollar_price,omitempty"`

	// monthly seats addon dollar price
	MonthlySeatsAddonDollarPrice int64 `json:"monthly_seats_addon_dollar_price,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// yearly dollar price
	YearlyDollarPrice int64 `json:"yearly_dollar_price,omitempty"`

	// yearly seats addon dollar price
	YearlySeatsAddonDollarPrice int64 `json:"yearly_seats_addon_dollar_price,omitempty"`
}

// Validate validates this account type
func (m *AccountType) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AccountType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountType) UnmarshalBinary(b []byte) error {
	var res AccountType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
