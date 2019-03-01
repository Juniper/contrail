package models

import (
	"strings"
)

// GetMacAddressesType returns MAC address of a VMI in MacAddressType format.
// If MAC address is unknown, then new one is initialized based on object's UUID.
func (m *VirtualMachineInterface) GetMacAddressesType() (*MacAddressesType, error) {

	if addrs := m.GetVirtualMachineInterfaceMacAddresses().GetMacAddress(); len(addrs) == 1 {
		newMacAddress := strings.Replace(addrs[0], "-", ":", -1)
		return &MacAddressesType{
			MacAddress: []string{newMacAddress},
		}, nil
	}

	macAddress, err := uuidToMac(m.GetUUID())

	return &MacAddressesType{
		MacAddress: []string{macAddress},
	}, err
}

// GetRouterID returns UUID of VMI's LR.
func (m *VirtualMachineInterface) GetRouterID() string {
	if len(m.GetLogicalRouterBackRefs()) == 0 {
		return ""
	}
	return m.GetLogicalRouterBackRefs()[0].GetUUID()
}
