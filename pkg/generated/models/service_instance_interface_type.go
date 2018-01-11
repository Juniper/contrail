package models

// ServiceInstanceInterfaceType

import "encoding/json"

// ServiceInstanceInterfaceType
type ServiceInstanceInterfaceType struct {
	StaticRoutes        *RouteTableType      `json:"static_routes"`
	VirtualNetwork      string               `json:"virtual_network"`
	IPAddress           IpAddressType        `json:"ip_address"`
	AllowedAddressPairs *AllowedAddressPairs `json:"allowed_address_pairs"`
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
		AllowedAddressPairs: MakeAllowedAddressPairs(),
		StaticRoutes:        MakeRouteTableType(),
		VirtualNetwork:      "",
		IPAddress:           MakeIpAddressType(),
	}
}

// MakeServiceInstanceInterfaceTypeSlice() makes a slice of ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
	return []*ServiceInstanceInterfaceType{}
}
