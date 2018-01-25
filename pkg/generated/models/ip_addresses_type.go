package models

// IpAddressesType

// IpAddressesType
//proteus:generate
type IpAddressesType struct {
	IPAddress IpAddressType `json:"ip_address,omitempty"`
}

// MakeIpAddressesType makes IpAddressesType
func MakeIpAddressesType() *IpAddressesType {
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: MakeIpAddressType(),
	}
}

// MakeIpAddressesTypeSlice() makes a slice of IpAddressesType
func MakeIpAddressesTypeSlice() []*IpAddressesType {
	return []*IpAddressesType{}
}
