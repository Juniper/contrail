package services

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// IsVisibleObject verifies that the object is visible to a user without administrator rights
func IsVisibleObject(ctx context.Context, idPerms *models.IdPermsType) error {
	if ctx == nil {
		return nil
	}
	auth := auth.GetAuthCTX(ctx)
	if auth == nil {
		return errutil.ErrorBadRequestf("user unauthenticated: %s", auth.UserID())
	}
	if idPerms == nil {
		// by default UserVisible == true
		return nil
	}
	if !idPerms.UserVisible && !auth.IsAdmin() {
		return errutil.ErrorBadRequestf("this object is not visible by users: %s", auth.UserID())
	}
	return nil
}

var defaultIDPerms = &models.IdPermsType{
	Enable: true,
	Permissions: &models.PermType{
		Owner:       "cloud-admin",
		OwnerAccess: basemodels.PermsRWX,
		OtherAccess: basemodels.PermsRWX,
		Group:       "cloud-admin-group",
		GroupAccess: basemodels.PermsRWX,
	},
}

// CreateIDPerms sets default permissions for a new object
func CreateIDPerms(idPerms *models.IdPermsType, uuid string) error {
	if idPerms == nil {
		return nil
	}
	if idPerms.UUID == nil {
		idPerms.UUID = models.NewUUIDType(uuid)
	}

	if idPerms.Permissions == nil {
		idPerms.Permissions = new(models.PermType)
	}
	return EnsureIDPerms(idPerms, defaultIDPerms)
}

// EnsureIDPerms sets default permissions for an existing object
// nolint: gocyclo
func EnsureIDPerms(idPerms *models.IdPermsType, storedIDPerms *models.IdPermsType) error {
	if idPerms == nil {
		return nil
	}
	if storedIDPerms == nil {
		return nil
	}
	if !idPerms.Enable {
		idPerms.Enable = storedIDPerms.Enable
	}
	if idPerms.Description == "" {
		idPerms.Description = storedIDPerms.Description
	}
	if idPerms.Created == "" {
		idPerms.Created = storedIDPerms.Created
	}
	if idPerms.Creator == "" {
		idPerms.Creator = storedIDPerms.Creator
	}
	if idPerms.LastModified == "" {
		idPerms.LastModified = storedIDPerms.LastModified
	}

	if storedIDPerms.Permissions == nil {
		return nil
	}
	if idPerms.Permissions == nil {
		idPerms.Permissions = new(models.PermType)
	}
	if idPerms.Permissions.Owner == "" {
		idPerms.Permissions.Owner = storedIDPerms.Permissions.Owner
	}
	if idPerms.Permissions.OwnerAccess == 0 {
		idPerms.Permissions.OwnerAccess = storedIDPerms.Permissions.OwnerAccess
	}
	if idPerms.Permissions.OtherAccess == 0 {
		idPerms.Permissions.OtherAccess = storedIDPerms.Permissions.OtherAccess
	}
	if idPerms.Permissions.Group == "" {
		idPerms.Permissions.Group = storedIDPerms.Permissions.Group
	}
	if idPerms.Permissions.GroupAccess == 0 {
		idPerms.Permissions.GroupAccess = storedIDPerms.Permissions.GroupAccess
	}

	return nil
}

// IsPermsUUIDMatch warns when idPerm.UUID does not match uuid
func IsPermsUUIDMatch(idPerm *models.IdPermsType, uuid string) error {
	if idPerm == nil {
		return nil
	}
	if idPerm.UUID == nil {
		return nil
	}
	uuidType := models.NewUUIDType(uuid)
	if uuidType.UUIDLslong != idPerm.UUID.UUIDLslong || uuidType.UUIDMslong != idPerm.UUID.UUIDMslong {
		logrus.Warn("UUID mismatch")
	}
	return nil
}
