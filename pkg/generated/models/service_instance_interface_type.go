package models

// ServiceInstanceInterfaceType

import "encoding/json"

// ServiceInstanceInterfaceType
type ServiceInstanceInterfaceType struct {
	VirtualNetwork      string               `json:"virtual_network,omitempty"`
	IPAddress           IpAddressType        `json:"ip_address,omitempty"`
	AllowedAddressPairs *AllowedAddressPairs `json:"allowed_address_pairs,omitempty"`
	StaticRoutes        *RouteTableType      `json:"static_routes,omitempty"`
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
		StaticRoutes:        MakeRouteTableType(),
		VirtualNetwork:      "",
		IPAddress:           MakeIpAddressType(),
		AllowedAddressPairs: MakeAllowedAddressPairs(),
	}
}

// MakeServiceInstanceInterfaceTypeSlice() makes a slice of ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
	return []*ServiceInstanceInterfaceType{}
}
