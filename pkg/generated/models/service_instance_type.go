package models

// ServiceInstanceType

import "encoding/json"

// ServiceInstanceType
type ServiceInstanceType struct {
	HaMode                   AddressMode                     `json:"ha_mode,omitempty"`
	InterfaceList            []*ServiceInstanceInterfaceType `json:"interface_list,omitempty"`
	LeftVirtualNetwork       string                          `json:"left_virtual_network,omitempty"`
	AutoPolicy               bool                            `json:"auto_policy"`
	RightVirtualNetwork      string                          `json:"right_virtual_network,omitempty"`
	RightIPAddress           IpAddressType                   `json:"right_ip_address,omitempty"`
	ManagementVirtualNetwork string                          `json:"management_virtual_network,omitempty"`
	ScaleOut                 *ServiceScaleOutType            `json:"scale_out,omitempty"`
	AvailabilityZone         string                          `json:"availability_zone,omitempty"`
	VirtualRouterID          string                          `json:"virtual_router_id,omitempty"`
	LeftIPAddress            IpAddressType                   `json:"left_ip_address,omitempty"`
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
		AvailabilityZone: "",
		VirtualRouterID:  "",
		LeftIPAddress:    MakeIpAddressType(),

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		LeftVirtualNetwork:       "",
		AutoPolicy:               false,
		RightVirtualNetwork:      "",
		RightIPAddress:           MakeIpAddressType(),
		ManagementVirtualNetwork: "",
		ScaleOut:                 MakeServiceScaleOutType(),
		HaMode:                   MakeAddressMode(),
	}
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
	return []*ServiceInstanceType{}
}
