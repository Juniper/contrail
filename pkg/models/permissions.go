package models

import (
	"github.com/Juniper/asf/pkg/models"
)

var defaultPermissions = &PermType{
	Owner:       "cloud-admin",
	OwnerAccess: models.PermsRWX,
	OtherAccess: models.PermsRWX,
	Group:       "cloud-admin-group",
	GroupAccess: models.PermsRWX,
}

// Merge sets undefined fields from source values
func (m *PermType) Merge(source *PermType) {
	if source == nil {
		return
	}
	if m.Owner == "" {
		m.Owner = source.Owner
	}
	if m.OwnerAccess == 0 {
		m.OwnerAccess = source.OwnerAccess
	}
	if m.OtherAccess == 0 {
		m.OtherAccess = source.OtherAccess
	}
	if m.Group == "" {
		m.Group = source.Group
	}
	if m.GroupAccess == 0 {
		m.GroupAccess = source.GroupAccess
	}
}
