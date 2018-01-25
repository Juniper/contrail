package models

// AllowedAddressPair

// AllowedAddressPair
//proteus:generate
type AllowedAddressPair struct {
	IP          *SubnetType `json:"ip,omitempty"`
	Mac         string      `json:"mac,omitempty"`
	AddressMode AddressMode `json:"address_mode,omitempty"`
}

// MakeAllowedAddressPair makes AllowedAddressPair
func MakeAllowedAddressPair() *AllowedAddressPair {
	return &AllowedAddressPair{
		//TODO(nati): Apply default
		IP:          MakeSubnetType(),
		Mac:         "",
		AddressMode: MakeAddressMode(),
	}
}

// MakeAllowedAddressPairSlice() makes a slice of AllowedAddressPair
func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
	return []*AllowedAddressPair{}
}
