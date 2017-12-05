package models

// FloatingIpPoolSubnetType

import "encoding/json"

type FloatingIpPoolSubnetType struct {
	SubnetUUID []string `json:"subnet_uuid"`
}

func (model *FloatingIpPoolSubnetType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeFloatingIpPoolSubnetType() *FloatingIpPoolSubnetType {
	return &FloatingIpPoolSubnetType{
		//TODO(nati): Apply default
		SubnetUUID: []string{},
	}
}

func InterfaceToFloatingIpPoolSubnetType(iData interface{}) *FloatingIpPoolSubnetType {
	data := iData.(map[string]interface{})
	return &FloatingIpPoolSubnetType{
		SubnetUUID: data["subnet_uuid"].([]string),

		//{"Title":"","Description":"List of subnets associated with this floating ip pool.","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SubnetUUID","GoType":"string"},"GoName":"SubnetUUID","GoType":"[]string"}

	}
}

func InterfaceToFloatingIpPoolSubnetTypeSlice(data interface{}) []*FloatingIpPoolSubnetType {
	list := data.([]interface{})
	result := MakeFloatingIpPoolSubnetTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFloatingIpPoolSubnetType(item))
	}
	return result
}

func MakeFloatingIpPoolSubnetTypeSlice() []*FloatingIpPoolSubnetType {
	return []*FloatingIpPoolSubnetType{}
}
