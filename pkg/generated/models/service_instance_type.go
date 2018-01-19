package models

// ServiceInstanceType

import "encoding/json"

// ServiceInstanceType
type ServiceInstanceType struct {
	AvailabilityZone         string                          `json:"availability_zone,omitempty"`
	ManagementVirtualNetwork string                          `json:"management_virtual_network,omitempty"`
	ScaleOut                 *ServiceScaleOutType            `json:"scale_out,omitempty"`
	AutoPolicy               bool                            `json:"auto_policy"`
	RightVirtualNetwork      string                          `json:"right_virtual_network,omitempty"`
	RightIPAddress           IpAddressType                   `json:"right_ip_address,omitempty"`
	HaMode                   AddressMode                     `json:"ha_mode,omitempty"`
	VirtualRouterID          string                          `json:"virtual_router_id,omitempty"`
	InterfaceList            []*ServiceInstanceInterfaceType `json:"interface_list,omitempty"`
	LeftIPAddress            IpAddressType                   `json:"left_ip_address,omitempty"`
	LeftVirtualNetwork       string                          `json:"left_virtual_network,omitempty"`
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
		RightVirtualNetwork: "",
		RightIPAddress:      MakeIpAddressType(),
		HaMode:              MakeAddressMode(),
		VirtualRouterID:     "",

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		LeftIPAddress:            MakeIpAddressType(),
		LeftVirtualNetwork:       "",
		AvailabilityZone:         "",
		ManagementVirtualNetwork: "",
		ScaleOut:                 MakeServiceScaleOutType(),
		AutoPolicy:               false,
	}
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
	return []*ServiceInstanceType{}
}
