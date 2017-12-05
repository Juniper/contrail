package models

// VirtualRouterNetworkIpamType

import "encoding/json"

type VirtualRouterNetworkIpamType struct {
	AllocationPools []*AllocationPoolType `json:"allocation_pools"`
	Subnet          []*SubnetType         `json:"subnet"`
}

func (model *VirtualRouterNetworkIpamType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeVirtualRouterNetworkIpamType() *VirtualRouterNetworkIpamType {
	return &VirtualRouterNetworkIpamType{
		//TODO(nati): Apply default

		Subnet: MakeSubnetTypeSlice(),

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

func InterfaceToVirtualRouterNetworkIpamType(iData interface{}) *VirtualRouterNetworkIpamType {
	data := iData.(map[string]interface{})
	return &VirtualRouterNetworkIpamType{

		AllocationPools: InterfaceToAllocationPoolTypeSlice(data["allocation_pools"]),

		//{"Title":"","Description":"List of ranges of ip address for vrouter specific allocation","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"allocation_pools","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"end":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"End","GoType":"string"},"start":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Start","GoType":"string"},"vrouter_specific_pool":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VrouterSpecificPool","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllocationPoolType","CollectionType":"","Column":"","Item":null,"GoName":"AllocationPools","GoType":"AllocationPoolType"},"GoName":"AllocationPools","GoType":"[]*AllocationPoolType"}

		Subnet: InterfaceToSubnetTypeSlice(data["subnet"]),

		//{"Title":"","Description":"List of ip prefix and length for vrouter specific subnets","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"subnet","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix","Item":null,"GoName":"IPPrefix","GoType":"string"},"ip_prefix_len":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_prefix_len","Item":null,"GoName":"IPPrefixLen","GoType":"int"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"Subnet","GoType":"SubnetType"},"GoName":"Subnet","GoType":"[]*SubnetType"}

	}
}

func InterfaceToVirtualRouterNetworkIpamTypeSlice(data interface{}) []*VirtualRouterNetworkIpamType {
	list := data.([]interface{})
	result := MakeVirtualRouterNetworkIpamTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToVirtualRouterNetworkIpamType(item))
	}
	return result
}

func MakeVirtualRouterNetworkIpamTypeSlice() []*VirtualRouterNetworkIpamType {
	return []*VirtualRouterNetworkIpamType{}
}
