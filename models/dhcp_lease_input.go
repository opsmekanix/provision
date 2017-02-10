package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DhcpLeaseInput DHCP Lease Input
// swagger:model dhcp-lease-input
type DhcpLeaseInput struct {

	// expire time
	// Required: true
	ExpireTime *strfmt.DateTime `json:"ExpireTime"`

	// Ip address
	// Required: true
	IPAddress *strfmt.IPv4 `json:"IpAddress"`

	// mac address
	// Required: true
	// Max Length: 17
	// Min Length: 17
	// Pattern: ^([0-9a-f]{2}):{5}[0-9a-f]{2}$
	MacAddress *string `json:"MacAddress"`

	// valid
	// Required: true
	Valid *bool `json:"Valid"`
}

// Validate validates this dhcp lease input
func (m *DhcpLeaseInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpireTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DhcpLeaseInput) validateExpireTime(formats strfmt.Registry) error {

	if err := validate.Required("ExpireTime", "body", m.ExpireTime); err != nil {
		return err
	}

	return nil
}

func (m *DhcpLeaseInput) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("IpAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *DhcpLeaseInput) validateMacAddress(formats strfmt.Registry) error {

	if err := validate.Required("MacAddress", "body", m.MacAddress); err != nil {
		return err
	}

	if err := validate.MinLength("MacAddress", "body", string(*m.MacAddress), 17); err != nil {
		return err
	}

	if err := validate.MaxLength("MacAddress", "body", string(*m.MacAddress), 17); err != nil {
		return err
	}

	if err := validate.Pattern("MacAddress", "body", string(*m.MacAddress), `^([0-9a-f]{2}):{5}[0-9a-f]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *DhcpLeaseInput) validateValid(formats strfmt.Registry) error {

	if err := validate.Required("Valid", "body", m.Valid); err != nil {
		return err
	}

	return nil
}
