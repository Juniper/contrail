package models

import (
	"strings"

	"github.com/Juniper/contrail/pkg/models"
)

func (m *VirtualMachineInterface) GetMacAddressesType() (*models.MacAddressesType, error) {

	if addrs := m.GetVirtualMachineInterfaceMacAddresses().GetMacAddress(); len(addrs) == 1 {
		newMacAddress := strings.Replace(addrs[0], "-", ":", -1)
		return &models.MacAddressesType{
			MacAddress: []string{newMacAddress},
		}, nil
	}

	macAddress, err := uuidToMac(m.GetUUID())
	if err != nil {
		return nil, err
	}

	return &models.MacAddressesType{
		MacAddress: []string{macAddress},
	}, nil
}
