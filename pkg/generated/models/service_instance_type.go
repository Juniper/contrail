package models

// ServiceInstanceType

import "encoding/json"

// ServiceInstanceType
type ServiceInstanceType struct {
	HaMode                   AddressMode                     `json:"ha_mode,omitempty"`
	VirtualRouterID          string                          `json:"virtual_router_id,omitempty"`
	InterfaceList            []*ServiceInstanceInterfaceType `json:"interface_list,omitempty"`
	LeftIPAddress            IpAddressType                   `json:"left_ip_address,omitempty"`
	RightVirtualNetwork      string                          `json:"right_virtual_network,omitempty"`
	RightIPAddress           IpAddressType                   `json:"right_ip_address,omitempty"`
	AvailabilityZone         string                          `json:"availability_zone,omitempty"`
	ScaleOut                 *ServiceScaleOutType            `json:"scale_out,omitempty"`
	LeftVirtualNetwork       string                          `json:"left_virtual_network,omitempty"`
	ManagementVirtualNetwork string                          `json:"management_virtual_network,omitempty"`
	AutoPolicy               bool                            `json:"auto_policy"`
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
		AutoPolicy:               false,
		VirtualRouterID:          "",

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		LeftIPAddress:       MakeIpAddressType(),
		RightVirtualNetwork: "",
		RightIPAddress:      MakeIpAddressType(),
		AvailabilityZone:    "",
		ScaleOut:            MakeServiceScaleOutType(),
		HaMode:              MakeAddressMode(),
		LeftVirtualNetwork:  "",
	}
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
	return []*ServiceInstanceType{}
}
