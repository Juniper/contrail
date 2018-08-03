package models

import (
	fmt "fmt"
	"regexp"
	"strings"
	"time"

	"net"

	errors "github.com/pkg/errors"
)

//NewTypeValidatorWithFormat creates new TypeValidator with format validators
func NewTypeValidatorWithFormat() (*TypeValidator, error) {
	tv := &TypeValidator{}
	tv.SchemaValidator = SchemaValidator{}

	tv.SchemaValidator.validators = map[string]func(string) error{}

	// Initialize TypeValidator

	// Create regex used while validating CommunityAttributes
	tv.communityAttributeRegexStr = "^[0-9]+:[0-9]+$"
	r, err := regexp.Compile(tv.communityAttributeRegexStr)
	if err != nil {
		return nil, err
	}
	tv.communityAttributeRegex = r

	// Register all format validators
	err = tv.addHostnameFormatValidator()
	if err != nil {
		return nil, err
	}

	err = tv.addIPv4FormatValidator()
	if err != nil {
		return nil, err
	}

	err = tv.addMacAddressFormatValidator()
	if err != nil {
		return nil, err
	}

	err = tv.addDateTimeFormatValidator()
	if err != nil {
		return nil, err
	}

	err = tv.addServiceInterfaceTypeFormatValidator()
	if err != nil {
		return nil, err
	}

	return tv, nil
}

func (tv *TypeValidator) addHostnameFormatValidator() error {
	validator := "hostname"
	// regex from https://www.regextester.com/23s
	regexStr := `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*` +
		`([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}

	tv.SchemaValidator.addFormatValidator(validator, func(value string) error {
		// Validate hostname

		if len(value) > 255 {
			return errors.Errorf("Invalid format. Hostname too long.")
		}

		if value[len(value)-1] == '.' {
			value = value[:len(value)-1]
		}

		if !regex.MatchString(value) {
			return errors.Errorf("Invalid hostname format.")
		}

		return nil
	})
	return nil
}

func (tv *TypeValidator) addIPv4FormatValidator() error {
	validator := "ipv4"
	ipv4Format := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	regex, err := regexp.Compile(ipv4Format)
	if err != nil {
		return err
	}

	tv.SchemaValidator.addFormatValidator(validator, func(value string) error {
		ip := net.ParseIP(value)
		if ip == nil {
			return errors.Errorf("\"%s\" is an invalid IPv4 value", value)
		}
		if !regex.MatchString(value) {
			return errors.Errorf("Invalid IPv4 format. It should match \"%s\"", ipv4Format)
		}
		return nil
	})
	return nil
}

func (tv *TypeValidator) addMacAddressFormatValidator() error {
	validator := "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
	// TODO: We should extend schema for using regexes directly from it or rename such formats (e.g. mac-address)
	regex, err := regexp.Compile(validator)
	if err != nil {
		return err
	}

	tv.SchemaValidator.addFormatValidator(validator, func(value string) error {
		// Validate ^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$

		if !regex.MatchString(value) {
			return errors.Errorf("Invalid format. It should match \"%s\"", validator)
		}
		return nil
	})
	return nil
}

func (tv *TypeValidator) addDateTimeFormatValidator() error {
	validator := "date-time"

	tv.SchemaValidator.addFormatValidator(validator, func(value string) error {
		// Validate date-time
		dateTimeFormat := "2006-01-02T15:04:05"
		_, err := time.Parse(dateTimeFormat, value)
		if err != nil {
			return errors.Wrapf(err, "Invalid format. Expected: %s", dateTimeFormat)
		}
		return nil
	})
	return nil
}

func (tv *TypeValidator) addServiceInterfaceTypeFormatValidator() error {
	validator := "service_interface_type_format"
	regexStr := "^other[0-9]*$"
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}

	tv.SchemaValidator.addFormatValidator(validator, func(value string) error {
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

		if !regex.MatchString(value) {
			return errors.Errorf("ServiceInterfaceType value (%s) must be either one of "+
				"[%s] or match \"%s\"", value, strings.Join(mapKeys(restrictions), ", "), regexStr)
		}
		return nil
	})

	return nil
}

//SchemaValidator implementing basic checks based on information in schema
type SchemaValidator struct {
	validators map[string]func(string) error
}

//TypeValidator embedding SchemaValidator validator. It enables defining custom validation for each type
type TypeValidator struct {
	SchemaValidator
	communityAttributeRegex    *regexp.Regexp
	communityAttributeRegexStr string
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
