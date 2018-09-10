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

func (i *VirtualNetworkIntent) GetObject() basemodels.Object {
	return i.VirtualNetwork
}

func LoadVirtualNetworkIntent(
	c *intent.Cache,
	uuid string,
) (*VirtualNetworkIntent, bool) {
	i, ok := c.Load(models.KindVirtualNetwork, uuid)
	if !ok {
		return nil, false
	}
	actual, ok := i.(*VirtualNetworkIntent)
	if !ok {
		return nil, false
	}
	return actual, true
}
