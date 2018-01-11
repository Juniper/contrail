package models

// AllowedAddressPairs

import "encoding/json"

// AllowedAddressPairs
type AllowedAddressPairs struct {
	AllowedAddressPair []*AllowedAddressPair `json:"allowed_address_pair"`
}

// String returns json representation of the object
func (model *AllowedAddressPairs) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAllowedAddressPairs makes AllowedAddressPairs
func MakeAllowedAddressPairs() *AllowedAddressPairs {
	return &AllowedAddressPairs{
		//TODO(nati): Apply default

		AllowedAddressPair: MakeAllowedAddressPairSlice(),
	}
}

// MakeAllowedAddressPairsSlice() makes a slice of AllowedAddressPairs
func MakeAllowedAddressPairsSlice() []*AllowedAddressPairs {
	return []*AllowedAddressPairs{}
}
