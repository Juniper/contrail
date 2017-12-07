package models

// ServiceInstanceType

import "encoding/json"

// ServiceInstanceType
type ServiceInstanceType struct {
	AvailabilityZone         string                          `json:"availability_zone"`
	ScaleOut                 *ServiceScaleOutType            `json:"scale_out"`
	VirtualRouterID          string                          `json:"virtual_router_id"`
	InterfaceList            []*ServiceInstanceInterfaceType `json:"interface_list"`
	LeftVirtualNetwork       string                          `json:"left_virtual_network"`
	AutoPolicy               bool                            `json:"auto_policy"`
	RightVirtualNetwork      string                          `json:"right_virtual_network"`
	RightIPAddress           IpAddressType                   `json:"right_ip_address"`
	ManagementVirtualNetwork string                          `json:"management_virtual_network"`
	HaMode                   AddressMode                     `json:"ha_mode"`
	LeftIPAddress            IpAddressType                   `json:"left_ip_address"`
}

//  parents relation object

// String returns json representation of the object
func (model *ServiceInstanceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceInstanceType makes ServiceInstanceType
func MakeServiceInstanceType() *ServiceInstanceType {
	return &ServiceInstanceType{
		//TODO(nati): Apply default
		HaMode:                   MakeAddressMode(),
		LeftIPAddress:            MakeIpAddressType(),
		RightVirtualNetwork:      "",
		RightIPAddress:           MakeIpAddressType(),
		ManagementVirtualNetwork: "",

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		LeftVirtualNetwork: "",
		AutoPolicy:         false,
		AvailabilityZone:   "",
		ScaleOut:           MakeServiceScaleOutType(),
		VirtualRouterID:    "",
	}
}

// InterfaceToServiceInstanceType makes ServiceInstanceType from interface
func InterfaceToServiceInstanceType(iData interface{}) *ServiceInstanceType {
	data := iData.(map[string]interface{})
	return &ServiceInstanceType{
		RightVirtualNetwork: data["right_virtual_network"].(string),

		//{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"RightVirtualNetwork","GoType":"string","GoPremitive":true}
		RightIPAddress: InterfaceToIpAddressType(data["right_ip_address"]),

		//{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"RightIPAddress","GoType":"IpAddressType","GoPremitive":false}
		ManagementVirtualNetwork: data["management_virtual_network"].(string),

		//{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ManagementVirtualNetwork","GoType":"string","GoPremitive":true}
		HaMode: InterfaceToAddressMode(data["ha_mode"]),

		//{"Title":"","Description":"When scale-out is greater than one, decides if active-active or active-backup, used in version 1 (V1) only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"","Item":null,"GoName":"HaMode","GoType":"AddressMode","GoPremitive":false}
		LeftIPAddress: InterfaceToIpAddressType(data["left_ip_address"]),

		//{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"LeftIPAddress","GoType":"IpAddressType","GoPremitive":false}
		AvailabilityZone: data["availability_zone"].(string),

		//{"Title":"","Description":"Availability zone used to spawn VM(s) for this service instance, used in version 1 (V1) only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AvailabilityZone","GoType":"string","GoPremitive":true}
		ScaleOut: InterfaceToServiceScaleOutType(data["scale_out"]),

		//{"Title":"","Description":"Number of virtual machines in this service instance, used in version 1 (V1) only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"auto_scale":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AutoScale","GoType":"bool","GoPremitive":true},"max_instances":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MaxInstances","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceScaleOutType","CollectionType":"","Column":"","Item":null,"GoName":"ScaleOut","GoType":"ServiceScaleOutType","GoPremitive":false}
		VirtualRouterID: data["virtual_router_id"].(string),

		//{"Title":"","Description":"UUID of a virtual-router on which this service instance need to spawn. Used to spawn services on CPE device when Nova is not present","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualRouterID","GoType":"string","GoPremitive":true}

		InterfaceList: InterfaceToServiceInstanceInterfaceTypeSlice(data["interface_list"]),

		//{"Title":"","Description":"List of service instance interface properties. Ordered list as per service template","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"allowed_address_pairs":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"allowed_address_pair":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"address_mode":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["active-active","active-standby"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AddressMode","CollectionType":"","Column":"","Item":null,"GoName":"AddressMode","GoType":"AddressMode","GoPremitive":false},"ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"ip_prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefix","GoType":"string","GoPremitive":true},"ip_prefix_len":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"IPPrefixLen","GoType":"int","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/SubnetType","CollectionType":"","Column":"","Item":null,"GoName":"IP","GoType":"SubnetType","GoPremitive":false},"mac":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Mac","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPair","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPair","GoType":"AllowedAddressPair","GoPremitive":false},"GoName":"AllowedAddressPair","GoType":"[]*AllowedAddressPair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/AllowedAddressPairs","CollectionType":"","Column":"","Item":null,"GoName":"AllowedAddressPairs","GoType":"AllowedAddressPairs","GoPremitive":false},"ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IPAddress","GoType":"IpAddressType","GoPremitive":false},"static_routes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"route":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attributes":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"community_attribute":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttribute","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttribute","GoType":"CommunityAttribute","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/CommunityAttributes","CollectionType":"","Column":"","Item":null,"GoName":"CommunityAttributes","GoType":"CommunityAttributes","GoPremitive":false},"next_hop":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"NextHop","GoType":"string","GoPremitive":true},"next_hop_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":["service-instance","ip-address"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteNextHopType","CollectionType":"","Column":"","Item":null,"GoName":"NextHopType","GoType":"RouteNextHopType","GoPremitive":false},"prefix":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Prefix","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteType","CollectionType":"","Column":"","Item":null,"GoName":"Route","GoType":"RouteType","GoPremitive":false},"GoName":"Route","GoType":"[]*RouteType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/RouteTableType","CollectionType":"","Column":"","Item":null,"GoName":"StaticRoutes","GoType":"RouteTableType","GoPremitive":false},"virtual_network":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetwork","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInstanceInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"InterfaceList","GoType":"ServiceInstanceInterfaceType","GoPremitive":false},"GoName":"InterfaceList","GoType":"[]*ServiceInstanceInterfaceType","GoPremitive":true}
		LeftVirtualNetwork: data["left_virtual_network"].(string),

		//{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"LeftVirtualNetwork","GoType":"string","GoPremitive":true}
		AutoPolicy: data["auto_policy"].(bool),

		//{"Title":"","Description":"Set when system creates internal service chains, example SNAT with router external flag in logical router","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AutoPolicy","GoType":"bool","GoPremitive":true}

	}
}

// InterfaceToServiceInstanceTypeSlice makes a slice of ServiceInstanceType from interface
func InterfaceToServiceInstanceTypeSlice(data interface{}) []*ServiceInstanceType {
	list := data.([]interface{})
	result := MakeServiceInstanceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInstanceType(item))
	}
	return result
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
	return []*ServiceInstanceType{}
}
