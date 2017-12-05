package models

// AddressType

import "encoding/json"

type AddressType struct {
	NetworkPolicy  string        `json:"network_policy"`
	SubnetList     []*SubnetType `json:"subnet_list"`
	VirtualNetwork string        `json:"virtual_network"`
	SecurityGroup  string        `json:"security_group"`
	Subnet         *SubnetType   `json:"subnet"`
}

func (model *AddressType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

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

func InterfaceToAddressType(iData interface{}) *AddressType {
	data := iData.(map[string]interface{})
	return &AddressType{
		SecurityGroup: data["security_group"].(string),

		//{"Title":"","Description":"Any address that belongs to interface with this security-group","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroup","GoType":"string"}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"Title":"","Description":"Any address that belongs to this subnet","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"}
		NetworkPolicy: data["network_policy"].(string),

		//{"Title":"","Description":"Any address that belongs to virtual network which has this policy attached","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NetworkPolicy","GoType":"string"}

		SubnetList: InterfaceToSubnetTypeSlice(data["subnet_list"]),

		//{"Title":"","Description":"Any address that belongs to any one of subnet in this list","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"SubnetList","GoType":"SubnetType"},"GoName":"SubnetList","GoType":"[]*SubnetType"}
		VirtualNetwork: data["virtual_network"].(string),

		//{"Title":"","Description":"Any address that belongs to this virtual network ","SQL":"","Default":null,"Operation":"","Presence":"exclusive","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string"}

	}
}

func InterfaceToAddressTypeSlice(data interface{}) []*AddressType {
	list := data.([]interface{})
	result := MakeAddressTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToAddressType(item))
	}
	return result
}

func MakeAddressTypeSlice() []*AddressType {
	return []*AddressType{}
}
