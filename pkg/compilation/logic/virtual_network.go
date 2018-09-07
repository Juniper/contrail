package logic

import (
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
)

type VirtualNetworkIntent struct {
	intent.BaseIntent
	*models.VirtualNetwork
}
