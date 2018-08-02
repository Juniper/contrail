package basemodels

import (
	fmt "fmt"
	"regexp"
	"strings"
	"time"

	"net"

	errors "github.com/pkg/errors"
)

//NewBaseValidatorWithFormat creates new BaseValidator with format validators
func NewBaseValidatorWithFormat() (*BaseValidator, error) {
	tv := &BaseValidator{}

	tv.validators = map[string]func(string) error{}

	// Register all format validators
	err := tv.addHostnameFormatValidator()
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

	err = tv.addBase64FormatValidator()
	if err != nil {
		return nil, err
	}

	return tv, nil
}

func (tv *BaseValidator) addBase64FormatValidator() error {
	validator := "base64"
	regexStr := "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}

	tv.AddFormatValidator(validator, func(value string) error {

		if !regex.MatchString(value) {
			return errors.Errorf("Invalid base64 format")
		}
		return nil
	})
	return nil
}

func (tv *BaseValidator) addHostnameFormatValidator() error {
	validator := "hostname"
	// regex from https://www.regextester.com/23s
	regexStr := `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*` +
		`([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}

	tv.AddFormatValidator(validator, func(value string) error {
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

func (tv *BaseValidator) addIPv4FormatValidator() error {
	validator := "ipv4"
	ipv4Format := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	regex, err := regexp.Compile(ipv4Format)
	if err != nil {
		return err
	}

	tv.AddFormatValidator(validator, func(value string) error {
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

func (tv *BaseValidator) addMacAddressFormatValidator() error {
	validator := "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
	// TODO: We should extend schema for using regexes directly from it or rename such formats (e.g. mac-address)
	regex, err := regexp.Compile(validator)
	if err != nil {
		return err
	}

	tv.AddFormatValidator(validator, func(value string) error {
		// Validate ^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$

		if !regex.MatchString(value) {
			return errors.Errorf("Invalid format. It should match \"%s\"", validator)
		}
		return nil
	})
	return nil
}

func (tv *BaseValidator) addDateTimeFormatValidator() error {
	validator := "date-time"

	tv.AddFormatValidator(validator, func(value string) error {
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

func (tv *BaseValidator) addServiceInterfaceTypeFormatValidator() error {
	validator := "service_interface_type_format"
	regexStr := "^other[0-9]*$"
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}

	tv.AddFormatValidator(validator, func(value string) error {
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

//BaseValidator embedding SchemaValidator validator. It enables defining custom validation for each type
type BaseValidator struct {
	validators map[string]func(string) error
}

//AddFormatValidator adds format validator.
func (tv *BaseValidator) AddFormatValidator(format string, validator func(string) error) {
	_, present := tv.validators[format]
	if !present {
		tv.validators[format] = validator
	}
}

//GetFormatValidator  get a format validator.
func (tv *BaseValidator) GetFormatValidator(format string) (func(string) error, error) {
	validator, present := tv.validators[format]
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
