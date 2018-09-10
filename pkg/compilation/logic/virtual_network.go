package logic

import (
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// VirtualNetworkIntent intent
type VirtualNetworkIntent struct {
	intent.BaseIntent
	*models.VirtualNetwork
}

// GetObject returns embedded resource object
func (i *VirtualNetworkIntent) GetObject() basemodels.Object {
	return i.VirtualNetwork
}

// LoadVirtualNetworkIntent loads VirtualNetworkIntent from cache
func LoadVirtualNetworkIntent(
	c *intent.Cache,
	uuid string,
) *VirtualNetworkIntent {
	i := c.Load(models.KindVirtualNetwork, intent.ByUUID(uuid))
	actual, _ := i.(*VirtualNetworkIntent)
	return actual
}
