package models

// AddressType

import "encoding/json"

// AddressType
type AddressType struct {
	SecurityGroup  string        `json:"security_group"`
	Subnet         *SubnetType   `json:"subnet"`
	NetworkPolicy  string        `json:"network_policy"`
	SubnetList     []*SubnetType `json:"subnet_list"`
	VirtualNetwork string        `json:"virtual_network"`
}

//  parents relation object

// String returns json representation of the object
func (model *AddressType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAddressType makes AddressType
func MakeAddressType() *AddressType {
	return &AddressType{
		//TODO(nati): Apply default
		SecurityGroup: "",
		Subnet:        MakeSubnetType(),
		NetworkPolicy: "",

		SubnetList: MakeSubnetTypeSlice(),

		VirtualNetwork: "",
	}
}

// InterfaceToAddressType makes AddressType from interface
func InterfaceToAddressType(iData interface{}) *AddressType {
	data := iData.(map[string]interface{})
	return &AddressType{
		NetworkPolicy: data["network_policy"].(string),

		//{"Title":"","Description":"Any address that belongs to virtual network which has this policy attached","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string","GoPremitive":true}

		SubnetList: InterfaceToSubnetTypeSlice(data["subnet_list"]),

		//{"Title":"","Description":"Any address that belongs to any one of subnet in this list","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType","GoPremitive":false},"GoName":"SubnetList","GoType":"[]*SubnetType","GoPremitive":true}
		VirtualNetwork: data["virtual_network"].(string),

		//{"Title":"","Description":"Any address that belongs to this virtual network ","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string","GoPremitive":true}
		SecurityGroup: data["security_group"].(string),

		//{"Title":"","Description":"Any address that belongs to interface with this security-group","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string","GoPremitive":true}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"Title":"","Description":"Any address that belongs to this subnet","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType","GoPremitive":false}

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
