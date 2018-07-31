package models

import "fmt"

// CheckMacAddress creates and assigns mac address to vmi instance when not provided
func (vmi *VirtualMachineInterface) CheckMacAddress() {
	if len(vmi.GetVirtualMachineInterfaceMacAddresses().GetMacAddress()) == 0 {
		uuid := vmi.GetUUID()
		macAddress := fmt.Sprintf("02:%s:%s:%s:%s:%s", uuid[0:2], uuid[2:4], uuid[4:6], uuid[6:8], uuid[9:11])
		vmi.VirtualMachineInterfaceMacAddresses = &MacAddressesType{
			MacAddress: []string{macAddress},
		}
	}
}
