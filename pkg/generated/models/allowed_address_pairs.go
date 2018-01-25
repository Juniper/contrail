package models

// AllowedAddressPairs

// AllowedAddressPairs
//proteus:generate
type AllowedAddressPairs struct {
	AllowedAddressPair []*AllowedAddressPair `json:"allowed_address_pair,omitempty"`
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
