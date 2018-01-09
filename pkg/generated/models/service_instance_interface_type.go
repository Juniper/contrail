package models

// ServiceInstanceInterfaceType

import "encoding/json"

// ServiceInstanceInterfaceType
type ServiceInstanceInterfaceType struct {
	VirtualNetwork      string               `json:"virtual_network"`
	IPAddress           IpAddressType        `json:"ip_address"`
	AllowedAddressPairs *AllowedAddressPairs `json:"allowed_address_pairs"`
	StaticRoutes        *RouteTableType      `json:"static_routes"`
}

// String returns json representation of the object
func (model *ServiceInstanceInterfaceType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceInstanceInterfaceType makes ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceType() *ServiceInstanceInterfaceType {
	return &ServiceInstanceInterfaceType{
		//TODO(nati): Apply default
		VirtualNetwork:      "",
		IPAddress:           MakeIpAddressType(),
		AllowedAddressPairs: MakeAllowedAddressPairs(),
		StaticRoutes:        MakeRouteTableType(),
	}
}

// InterfaceToServiceInstanceInterfaceType makes ServiceInstanceInterfaceType from interface
func InterfaceToServiceInstanceInterfaceType(iData interface{}) *ServiceInstanceInterfaceType {
	data := iData.(map[string]interface{})
	return &ServiceInstanceInterfaceType{
		IPAddress: InterfaceToIpAddressType(data["ip_address"]),

		//{"description":"Shared ip for this interface (Only V1)","type":"string"}
		AllowedAddressPairs: InterfaceToAllowedAddressPairs(data["allowed_address_pairs"]),

		//{"description":"Allowed address pairs, list of (IP address, MAC) for this interface","type":"object","properties":{"allowed_address_pair":{"type":"array","item":{"type":"object","properties":{"address_mode":{"type":"string","enum":["active-active","active-standby"]},"ip":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"mac":{"type":"string"}}}}}}
		StaticRoutes: InterfaceToRouteTableType(data["static_routes"]),

		//{"description":"Static routes for this interface (Only V1)","type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}}
		VirtualNetwork: data["virtual_network"].(string),

		//{"description":"Interface belongs to this virtual network.","type":"string"}

	}
}

// InterfaceToServiceInstanceInterfaceTypeSlice makes a slice of ServiceInstanceInterfaceType from interface
func InterfaceToServiceInstanceInterfaceTypeSlice(data interface{}) []*ServiceInstanceInterfaceType {
	list := data.([]interface{})
	result := MakeServiceInstanceInterfaceTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInstanceInterfaceType(item))
	}
	return result
}

// MakeServiceInstanceInterfaceTypeSlice() makes a slice of ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
	return []*ServiceInstanceInterfaceType{}
}
