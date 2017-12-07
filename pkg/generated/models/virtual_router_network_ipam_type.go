package models

// VirtualRouterNetworkIpamType

import "encoding/json"

// VirtualRouterNetworkIpamType
type VirtualRouterNetworkIpamType struct {
	AllocationPools []*AllocationPoolType `json:"allocation_pools"`
	Subnet          []*SubnetType         `json:"subnet"`
}

//  parents relation object

// String returns json representation of the object
func (model *VirtualRouterNetworkIpamType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType {
	return &VirtualRouterNetworkIpamType{
		//TODO(nati): Apply default

		AllocationPools: MakeAllocationPoolTypeSlice(),

		Subnet: MakeSubnetTypeSlice(),
	}
}

// InterfaceToVirtualRouterNetworkIpamType makes VirtualRouterNetworkIpamType from interface
func InterfaceToVirtualRouterNetworkIpamType(iData interface{}) *VirtualRouterNetworkIpamType {
	data := iData.(map[string]interface{})
	return &VirtualRouterNetworkIpamType{

		Subnet: InterfaceToSubnetTypeSlice(data["subnet"]),

		//{"Title":"","Description":"List of ip prefix and length for vrouter specific subnets","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"subnet","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix_len","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType","GoPremitive":false},"GoName":"Subnet","GoType":"[]*SubnetType","GoPremitive":true}

		AllocationPools: InterfaceToAllocationPoolTypeSlice(data["allocation_pools"]),

		//{"Title":"","Description":"List of ranges of ip address for vrouter specific allocation","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"allocation_pools","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"End","GoType":"string","GoPremitive":true},"start":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Start","GoType":"string","GoPremitive":true},"vrouter_specific_pool":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VrouterSpecificPool","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllocationPoolType","CollectionType":"","Column":"","Item":null,"GoName":"AllocationPools","GoType":"AllocationPoolType","GoPremitive":false},"GoName":"AllocationPools","GoType":"[]*AllocationPoolType","GoPremitive":true}

	}
}

// InterfaceToVirtualRouterNetworkIpamTypeSlice makes a slice of VirtualRouterNetworkIpamType from interface
func InterfaceToVirtualRouterNetworkIpamTypeSlice(data interface{}) []*VirtualRouterNetworkIpamType {
	list := data.([]interface{})
	result := MakeVirtualRouterNetworkIpamTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouterNetworkIpamType(item))
	}
	return result
}

// MakeVirtualRouterNetworkIpamTypeSlice() makes a slice of VirtualRouterNetworkIpamType
func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
	return []*VirtualRouterNetworkIpamType{}
}
