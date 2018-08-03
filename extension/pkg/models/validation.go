package models

import (
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

//NewTypeValidatorWithFormat creates new TypeValidator with format validators
func NewTypeValidatorWithFormat() (*TypeValidator, error) {
	base, err := basemodels.NewBaseValidatorWithFormat()
	if err != nil {
		return nil, err
	}
	tv := &TypeValidator{
		SchemaValidator: SchemaValidator{
			BaseValidator: base,
		},
	}
	return tv, nil
}

//SchemaValidator implementing basic checks based on information in schema
type SchemaValidator struct {
	*basemodels.BaseValidator
}

//TypeValidator embedding SchemaValidator validator. It enables defining custom validation for each type
type TypeValidator struct {
	SchemaValidator
}
