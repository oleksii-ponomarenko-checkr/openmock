// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ExpectAMQP expect a m q p
// swagger:model ExpectAMQP
type ExpectAMQP struct {

	// TODO
	Exchange string `json:"exchange,omitempty"`

	// TODO
	Queue string `json:"queue,omitempty"`

	// TODO
	RoutingKey string `json:"routing_key,omitempty"`
}

// Validate validates this expect a m q p
func (m *ExpectAMQP) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExpectAMQP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpectAMQP) UnmarshalBinary(b []byte) error {
	var res ExpectAMQP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}