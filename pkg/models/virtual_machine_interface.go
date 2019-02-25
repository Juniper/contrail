package models

import (
	"strings"
)

func (m *VirtualMachineInterface) GetMacAddressesType() (*MacAddressesType, error) {

	if addrs := m.GetVirtualMachineInterfaceMacAddresses().GetMacAddress(); len(addrs) == 1 {
		newMacAddress := strings.Replace(addrs[0], "-", ":", -1)
		return &MacAddressesType{
			MacAddress: []string{newMacAddress},
		}, nil
	}

	macAddress, err := uuidToMac(m.GetUUID())
	if err != nil {
		return nil, err
	}

	return &MacAddressesType{
		MacAddress: []string{macAddress},
	}, nil
}
