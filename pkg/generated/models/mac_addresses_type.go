package models

// MacAddressesType

// MacAddressesType
//proteus:generate
type MacAddressesType struct {
	MacAddress []string `json:"mac_address,omitempty"`
}

// MakeMacAddressesType makes MacAddressesType
func MakeMacAddressesType() *MacAddressesType {
	return &MacAddressesType{
		//TODO(nati): Apply default
		MacAddress: []string{},
	}
}

// MakeMacAddressesTypeSlice() makes a slice of MacAddressesType
func MakeMacAddressesTypeSlice() []*MacAddressesType {
	return []*MacAddressesType{}
}
