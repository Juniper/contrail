package models

// IpAddressesType

import "encoding/json"

// IpAddressesType
type IpAddressesType struct {
	IPAddress IpAddressType `json:"ip_address"`
}

// String returns json representation of the object
func (model *IpAddressesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpAddressesType makes IpAddressesType
func MakeIpAddressesType() *IpAddressesType {
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: MakeIpAddressType(),
	}
}

// InterfaceToIpAddressesType makes IpAddressesType from interface
func InterfaceToIpAddressesType(iData interface{}) *IpAddressesType {
	data := iData.(map[string]interface{})
	return &IpAddressesType{
		IPAddress: InterfaceToIpAddressType(data["ip_address"]),

		//{"type":"string"}

	}
}

// InterfaceToIpAddressesTypeSlice makes a slice of IpAddressesType from interface
func InterfaceToIpAddressesTypeSlice(data interface{}) []*IpAddressesType {
	list := data.([]interface{})
	result := MakeIpAddressesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpAddressesType(item))
	}
	return result
}

// MakeIpAddressesTypeSlice() makes a slice of IpAddressesType
func MakeIpAddressesTypeSlice() []*IpAddressesType {
	return []*IpAddressesType{}
}
