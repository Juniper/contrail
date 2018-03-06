package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeAddressType makes AddressType
// nolint
func MakeAddressType() *AddressType {
	return &AddressType{
		//TODO(nati): Apply default
		SecurityGroup: "",
		Subnet:        MakeSubnetType(),
		NetworkPolicy: "",

		SubnetList: MakeSubnetTypeSlice(),

		VirtualNetwork: "",
	}
}

// MakeAddressType makes AddressType
// nolint
func InterfaceToAddressType(i interface{}) *AddressType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &AddressType{
		//TODO(nati): Apply default
		SecurityGroup: common.InterfaceToString(m["security_group"]),
		Subnet:        InterfaceToSubnetType(m["subnet"]),
		NetworkPolicy: common.InterfaceToString(m["network_policy"]),

		SubnetList: InterfaceToSubnetTypeSlice(m["subnet_list"]),

		VirtualNetwork: common.InterfaceToString(m["virtual_network"]),
	}
}

// MakeAddressTypeSlice() makes a slice of AddressType
// nolint
func MakeAddressTypeSlice() []*AddressType {
	return []*AddressType{}
}

// InterfaceToAddressTypeSlice() makes a slice of AddressType
// nolint
func InterfaceToAddressTypeSlice(i interface{}) []*AddressType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*AddressType{}
	for _, item := range list {
		result = append(result, InterfaceToAddressType(item))
	}
	return result
}
