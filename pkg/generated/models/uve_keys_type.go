package models

// UveKeysType

import "encoding/json"

// UveKeysType
type UveKeysType struct {
	UveKey []string `json:"uve_key,omitempty"`
}

// String returns json representation of the object
func (model *UveKeysType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeUveKeysType makes UveKeysType
func MakeUveKeysType() *UveKeysType {
	return &UveKeysType{
		//TODO(nati): Apply default
		UveKey: []string{},
	}
}

// MakeUveKeysTypeSlice() makes a slice of UveKeysType
func MakeUveKeysTypeSlice() []*UveKeysType {
	return []*UveKeysType{}
}
