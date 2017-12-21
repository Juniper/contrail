package models

// ServiceInstanceType

import "encoding/json"

// ServiceInstanceType
type ServiceInstanceType struct {
	AutoPolicy               bool                            `json:"auto_policy"`
	RightVirtualNetwork      string                          `json:"right_virtual_network"`
	AvailabilityZone         string                          `json:"availability_zone"`
	HaMode                   AddressMode                     `json:"ha_mode"`
	VirtualRouterID          string                          `json:"virtual_router_id"`
	LeftIPAddress            IpAddressType                   `json:"left_ip_address"`
	LeftVirtualNetwork       string                          `json:"left_virtual_network"`
	RightIPAddress           IpAddressType                   `json:"right_ip_address"`
	ManagementVirtualNetwork string                          `json:"management_virtual_network"`
	ScaleOut                 *ServiceScaleOutType            `json:"scale_out"`
	InterfaceList            []*ServiceInstanceInterfaceType `json:"interface_list"`
}

// String returns json representation of the object
func (model *ServiceInstanceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceInstanceType makes ServiceInstanceType
func MakeServiceInstanceType() *ServiceInstanceType {
	return &ServiceInstanceType{
		//TODO(nati): Apply default
		RightIPAddress:           MakeIpAddressType(),
		ManagementVirtualNetwork: "",
		ScaleOut:                 MakeServiceScaleOutType(),

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		LeftIPAddress:       MakeIpAddressType(),
		LeftVirtualNetwork:  "",
		RightVirtualNetwork: "",
		AvailabilityZone:    "",
		HaMode:              MakeAddressMode(),
		VirtualRouterID:     "",
		AutoPolicy:          false,
	}
}

// InterfaceToServiceInstanceType makes ServiceInstanceType from interface
func InterfaceToServiceInstanceType(iData interface{}) *ServiceInstanceType {
	data := iData.(map[string]interface{})
	return &ServiceInstanceType{
		AutoPolicy: data["auto_policy"].(bool),

		//{"description":"Set when system creates internal service chains, example SNAT with router external flag in logical router","type":"boolean"}
		RightVirtualNetwork: data["right_virtual_network"].(string),

		//{"description":"Deprecated","type":"string"}
		AvailabilityZone: data["availability_zone"].(string),

		//{"description":"Availability zone used to spawn VM(s) for this service instance, used in version 1 (V1) only","type":"string"}
		HaMode: InterfaceToAddressMode(data["ha_mode"]),

		//{"description":"When scale-out is greater than one, decides if active-active or active-backup, used in version 1 (V1) only","type":"string","enum":["active-active","active-standby"]}
		VirtualRouterID: data["virtual_router_id"].(string),

		//{"description":"UUID of a virtual-router on which this service instance need to spawn. Used to spawn services on CPE device when Nova is not present","type":"string"}
		LeftIPAddress: InterfaceToIpAddressType(data["left_ip_address"]),

		//{"description":"Deprecated","type":"string"}
		LeftVirtualNetwork: data["left_virtual_network"].(string),

		//{"description":"Deprecated","type":"string"}
		RightIPAddress: InterfaceToIpAddressType(data["right_ip_address"]),

		//{"description":"Deprecated","type":"string"}
		ManagementVirtualNetwork: data["management_virtual_network"].(string),

		//{"description":"Deprecated","type":"string"}
		ScaleOut: InterfaceToServiceScaleOutType(data["scale_out"]),

		//{"description":"Number of virtual machines in this service instance, used in version 1 (V1) only","type":"object","properties":{"auto_scale":{"type":"boolean"},"max_instances":{"type":"integer"}}}

		InterfaceList: InterfaceToServiceInstanceInterfaceTypeSlice(data["interface_list"]),

		//{"description":"List of service instance interface properties. Ordered list as per service template","type":"array","item":{"type":"object","properties":{"allowed_address_pairs":{"type":"object","properties":{"allowed_address_pair":{"type":"array","item":{"type":"object","properties":{"address_mode":{"type":"string","enum":["active-active","active-standby"]},"ip":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"mac":{"type":"string"}}}}}},"ip_address":{"type":"string"},"static_routes":{"type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}},"virtual_network":{"type":"string"}}}}

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
