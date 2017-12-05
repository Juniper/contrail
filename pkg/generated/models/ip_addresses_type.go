package models

// IpAddressesType

import "encoding/json"

type IpAddressesType struct {
	IPAddress IpAddressType `json:"ip_address"`
}

func (model *IpAddressesType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeIpAddressesType() *IpAddressesType {
	return &IpAddressesType{
		//TODO(nati): Apply default
		IPAddress: MakeIpAddressType(),
	}
}

func InterfaceToIpAddressesType(iData interface{}) *IpAddressesType {
	data := iData.(map[string]interface{})
	return &IpAddressesType{
		IPAddress: InterfaceToIpAddressType(data["ip_address"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IPAddress","GoType":"IpAddressType"}

	}
}

func InterfaceToIpAddressesTypeSlice(data interface{}) []*IpAddressesType {
	list := data.([]interface{})
	result := MakeIpAddressesTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpAddressesType(item))
	}
	return result
}

func MakeIpAddressesTypeSlice() []*IpAddressesType {
	return []*IpAddressesType{}
}
