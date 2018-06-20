package models

import (
	fmt "fmt"
	"regexp"
	"strconv"
	"strings"

	"net"

	errors "github.com/pkg/errors"
)

//NewTypeValidatorWithFormat creates new TypeValidator with format validators
func NewTypeValidatorWithFormat() (*TypeValidator, error) {
	tv := &TypeValidator{}
	tv.SchemaValidator = SchemaValidator{}

	tv.SchemaValidator.validators = map[string]func(string) error{}

	// Register all format validators

	tv.SchemaValidator.addFormatValidator("hostname", func(value string) error {
		// Validate hostname
		// TODO
		return nil
	})

	tv.SchemaValidator.addFormatValidator("date-time", func(value string) error {
		// Validate date-time
		// TODO
		return nil
	})

	tv.SchemaValidator.addFormatValidator("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", func(value string) error {
		// Validate ^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$
		// TODO
		return nil
	})

	tv.SchemaValidator.addFormatValidator("service_interface_type_format", func(value string) error {
		// Validate service_interface_type_format
		restrictions := map[string]struct{}{

			"management": {},
			"left":       {},
			"right":      {},
		}

		_, present := restrictions[value]

		if present {
			return nil
		}

		regexStr := "other[0-9]*"

		r, _ := regexp.Compile(regexStr)

		if !r.MatchString(value) {
			return errors.Errorf("ServiceInterfaceType value (%s) must be either one of "+
				"[%s] or match \"%s\"", value, strings.Join(mapKeys(restrictions), ", "), regexStr)
		}
		return nil
	})

	return tv, nil
}

//SchemaValidator implementing basic checks based on information in schema
type SchemaValidator struct {
	validators map[string]func(string) error
}

//TypeValidator embedding SchemaValidator validator. It enables defining custom validation for each type
type TypeValidator struct {
	SchemaValidator
}

func (tv *TypeValidator) ValidateAllowedAddressPair(obj *AllowedAddressPair) error {
	err := tv.SchemaValidator.ValidateAllowedAddressPair(obj)

	if err != nil {
		return err
	}

	if obj.AddressMode != "active-standby" {
		return nil
	}

	ip := net.ParseIP(obj.IP.IPPrefix)

	if ip.To4() != nil {
		if obj.IP.IPPrefixLen < 120 {
			return errors.Errorf("IPv6 Prefix length lesser than 120 is not acceptable")
		}
	} else {
		if obj.IP.IPPrefixLen < 24 {
			return errors.Errorf("IPv4 Prefix length lesser than 24 is not acceptable")
		}
	}
	return nil
}

func (tv *TypeValidator) ValidateCommunityAttributes(obj *CommunityAttributes) error {

	restrictions := map[string]struct{}{

		"no-export":           {},
		"accept-own":          {},
		"no-advertise":        {},
		"no-export-subconfed": {},
		"no-reoriginate":      {},
	}

	regexStr := "[0-9]+:[0-9]+"

	r, _ := regexp.Compile(regexStr)

	for _, value := range obj.CommunityAttribute {

		_, present := restrictions[value]

		if present {
			continue
		}

		if !r.MatchString(value) {
			return errors.Errorf("CommunityAttribute value (%s) must be either one of "+
				"[%s] or match \"%s\"", value, strings.Join(mapKeys(restrictions), ", "), regexStr)
		}
		asn := strings.Split(value, ":")

		i, err := strconv.Atoi(asn[0])

		if err != nil {
			return errors.Wrapf(err, "Error while parsing CommunityAttribute.")
		}

		if i > 65535 {
			return errors.Errorf("Out of range ASN value %s. ASN values cannot exceed 65535.", i)

		}
	}

	return nil
}

func (sv *SchemaValidator) addFormatValidator(format string, validator func(string) error) {
	_, present := sv.validators[format]
	if !present {
		sv.validators[format] = validator
	}
}

func (sv *SchemaValidator) getFormatValidator(format string) (func(string) error, error) {
	validator, present := sv.validators[format]
	if !present {
		return nil, fmt.Errorf("%s format validator not found", format)
	}
	return validator, nil
}

// Returns array of map keys
func mapKeys(m map[string]struct{}) (keys []string) {
	for s := range m {
		keys = append(keys, s)
	}
	return keys
}
