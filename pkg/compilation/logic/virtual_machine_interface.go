package logic

import (
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
)

// VirtualNetworkIntent intent
type VirtualMachineInterfaceIntent struct {
	intent.BaseIntent
	*models.VirtualMachineInterface
}

func LoadVirtualMachineInterfaceIntent(
	c *intent.Cache,
	uuid string,
) (*VirtualMachineInterfaceIntent, bool) {
	i, ok := c.Load(models.KindVirtualMachineInterface, uuid)
	if !ok {
		return nil, false
	}
	actual, ok := i.(*VirtualMachineInterfaceIntent)
	if !ok {
		return nil, false
	}
	return actual, true
}
