package models

// ServiceInstanceType

import "encoding/json"

// ServiceInstanceType
type ServiceInstanceType struct {
	RightIPAddress           IpAddressType                   `json:"right_ip_address"`
	AvailabilityZone         string                          `json:"availability_zone"`
	ManagementVirtualNetwork string                          `json:"management_virtual_network"`
	HaMode                   AddressMode                     `json:"ha_mode"`
	InterfaceList            []*ServiceInstanceInterfaceType `json:"interface_list"`
	LeftIPAddress            IpAddressType                   `json:"left_ip_address"`
	AutoPolicy               bool                            `json:"auto_policy"`
	RightVirtualNetwork      string                          `json:"right_virtual_network"`
	ScaleOut                 *ServiceScaleOutType            `json:"scale_out"`
	VirtualRouterID          string                          `json:"virtual_router_id"`
	LeftVirtualNetwork       string                          `json:"left_virtual_network"`
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
		RightVirtualNetwork:      "",
		ScaleOut:                 MakeServiceScaleOutType(),
		VirtualRouterID:          "",
		LeftVirtualNetwork:       "",
		AutoPolicy:               false,
		RightIPAddress:           MakeIpAddressType(),
		AvailabilityZone:         "",
		ManagementVirtualNetwork: "",
		HaMode: MakeAddressMode(),

		InterfaceList: MakeServiceInstanceInterfaceTypeSlice(),

		LeftIPAddress: MakeIpAddressType(),
	}
}

// MakeServiceInstanceTypeSlice() makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
	return []*ServiceInstanceType{}
}
