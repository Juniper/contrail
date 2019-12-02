package logic

import (
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
)

// NetworkPolicyIntent intent
type NetworkPolicyIntent struct {
	intent.BaseIntent
	*models.NetworkPolicy
}

// GetObject returns embedded resource object
func (i *NetworkPolicyIntent) GetObject() basemodels.Object {
	return i.NetworkPolicy
}
