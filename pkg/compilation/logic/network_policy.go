package logic

import (
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
)

type NetworkPolicyIntent struct {
	intent.BaseIntent
	*models.NetworkPolicy
}
