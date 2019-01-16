package services

import (
	"context"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
)

// isVisibleObject verifies that the object is visible to a user without administrator rights
func isVisibleObject(ctx context.Context, idPerms *models.IdPermsType) error {
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
	if !idPerms.GetUserVisible() && !auth.IsAdmin() {
		return errutil.ErrorBadRequestf("this object is not visible by users: %s", auth.UserID())
	}
	return nil
}

func getStoredIDPerms(ctx context.Context, rs ReadService, typeName, uuid string) (*models.IdPermsType, error) {
	base, err := GetObject(ctx, rs, typeName, uuid)
	if err != nil {
		return nil, err
	}
	type baseIDPermser interface {
		GetIDPerms() *models.IdPermsType
	}
	m, ok := base.(baseIDPermser)
	if !ok {
		return nil, errutil.ErrorInternalf("method IDPerms() not found")
	}
	return m.GetIDPerms(), nil
}
