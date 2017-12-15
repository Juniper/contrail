package models

// AddressType

import "encoding/json"

// AddressType
type AddressType struct {
	Subnet         *SubnetType   `json:"subnet"`
	NetworkPolicy  string        `json:"network_policy"`
	SubnetList     []*SubnetType `json:"subnet_list"`
	VirtualNetwork string        `json:"virtual_network"`
	SecurityGroup  string        `json:"security_group"`
}

// String returns json representation of the object
func (model *AddressType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAddressType makes AddressType
func MakeAddressType() *AddressType {
	return &AddressType{
		//TODO(nati): Apply default
		Subnet:        MakeSubnetType(),
		NetworkPolicy: "",

		SubnetList: MakeSubnetTypeSlice(),

		VirtualNetwork: "",
		SecurityGroup:  "",
	}
}

// InterfaceToAddressType makes AddressType from interface
func InterfaceToAddressType(iData interface{}) *AddressType {
	data := iData.(map[string]interface{})
	return &AddressType{
		SecurityGroup: data["security_group"].(string),

		//{"description":"Any address that belongs to interface with this security-group","type":"string"}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"description":"Any address that belongs to this subnet","type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}
		NetworkPolicy: data["network_policy"].(string),

		//{"description":"Any address that belongs to virtual network which has this policy attached","type":"string"}

		SubnetList: InterfaceToSubnetTypeSlice(data["subnet_list"]),

		//{"description":"Any address that belongs to any one of subnet in this list","type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}}
		VirtualNetwork: data["virtual_network"].(string),

		//{"description":"Any address that belongs to this virtual network ","type":"string"}

	}
}

// InterfaceToAddressTypeSlice makes a slice of AddressType from interface
func InterfaceToAddressTypeSlice(data interface{}) []*AddressType {
	list := data.([]interface{})
	result := MakeAddressTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAddressType(item))
	}
	return result
}

// MakeAddressTypeSlice() makes a slice of AddressType
func MakeAddressTypeSlice() []*AddressType {
	return []*AddressType{}
}
