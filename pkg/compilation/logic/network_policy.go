package logic

import (
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
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
