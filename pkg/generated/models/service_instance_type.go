package models

// ServiceInstanceType

import "encoding/json"

// ServiceInstanceType
type ServiceInstanceType struct {
	RightIPAddress           IpAddressType                   `json:"right_ip_address,omitempty"`
	ManagementVirtualNetwork string                          `json:"management_virtual_network,omitempty"`
	VirtualRouterID          string                          `json:"virtual_router_id,omitempty"`
	LeftIPAddress            IpAddressType                   `json:"left_ip_address,omitempty"`
	LeftVirtualNetwork       string                          `json:"left_virtual_network,omitempty"`
	AutoPolicy               bool                            `json:"auto_policy"`
	RightVirtualNetwork      string                          `json:"right_virtual_network,omitempty"`
	AvailabilityZone         string                          `json:"availability_zone,omitempty"`
	ScaleOut                 *ServiceScaleOutType            `json:"scale_out,omitempty"`
	HaMode                   AddressMode                     `json:"ha_mode,omitempty"`
	InterfaceList            []*ServiceInstanceInterfaceType `json:"interface_list,omitempty"`
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
		ManagementVirtualNetwork: "",
		VirtualRouterID:          "",
		LeftIPAddress:            MakeIpAddressType(),
		LeftVirtualNetwork:       "",
		AutoPolicy:               false,
		RightIPAddress:           MakeIpAddressType(),
		AvailabilityZone:         "",
		ScaleOut:                 MakeServiceScaleOutType(),
		HaMode:                   MakeAddressMode(),

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		RightVirtualNetwork: "",
	}
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
	return []*ServiceInstanceType{}
}
