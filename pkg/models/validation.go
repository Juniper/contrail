package models

import (
	fmt "fmt"
)

//NewTypeValidatorWithFormat creates new TypeValidator with format validators
func NewTypeValidatorWithFormat() (*TypeValidator, error) {
	tv := &TypeValidator{}
	tv.SchemaValidator = SchemaValidator{}

	tv.SchemaValidator.validators = map[string]func(string) error{}

	// Register all format validators

	tv.SchemaValidator.addFormatValidator("hostname", func(hostname string) error {
		// Validate hostname
		// TODO
		return nil
	})

	tv.SchemaValidator.addFormatValidator("date-time", func(hostname string) error {
		// Validate date-time
		// TODO
		return nil
	})

	tv.SchemaValidator.addFormatValidator("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", func(hostname string) error {
		// Validate ^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$
		// TODO
		return nil
	})

	// TODO check all required format validators registered.

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
