package models

// FloatingIpPoolSubnetType

import "encoding/json"

// FloatingIpPoolSubnetType
type FloatingIpPoolSubnetType struct {
	SubnetUUID []string `json:"subnet_uuid"`
}

//  parents relation object

// String returns json representation of the object
func (model *FloatingIpPoolSubnetType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetType() *FloatingIpPoolSubnetType {
	return &FloatingIpPoolSubnetType{
		//TODO(nati): Apply default
		SubnetUUID: []string{},
	}
}

// InterfaceToFloatingIpPoolSubnetType makes FloatingIpPoolSubnetType from interface
func InterfaceToFloatingIpPoolSubnetType(iData interface{}) *FloatingIpPoolSubnetType {
	data := iData.(map[string]interface{})
	return &FloatingIpPoolSubnetType{
		SubnetUUID: data["subnet_uuid"].([]string),

		//{"Title":"","Description":"List of subnets associated with this floating ip pool.","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SubnetUUID","GoType":"string","GoPremitive":true},"GoName":"SubnetUUID","GoType":"[]string","GoPremitive":true}

	}
}

// InterfaceToFloatingIpPoolSubnetTypeSlice makes a slice of FloatingIpPoolSubnetType from interface
func InterfaceToFloatingIpPoolSubnetTypeSlice(data interface{}) []*FloatingIpPoolSubnetType {
	list := data.([]interface{})
	result := MakeFloatingIpPoolSubnetTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToFloatingIpPoolSubnetType(item))
	}
	return result
}

// MakeFloatingIpPoolSubnetTypeSlice() makes a slice of FloatingIpPoolSubnetType
func MakeFloatingIpPoolSubnetTypeSlice() []*FloatingIpPoolSubnetType {
	return []*FloatingIpPoolSubnetType{}
}
