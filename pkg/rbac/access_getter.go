package rbac

import (
	"context"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/rbac"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"
)

// ContrailAccessGetter allows getting APIAccessLists and resources' Perms2.
type ContrailAccessGetter struct {
	readService services.ReadService
	log         *logrus.Entry
}

// NewContrailAccessGetter creates new ContrailAccessGetter.
func NewContrailAccessGetter(rs services.ReadService) *ContrailAccessGetter {
	return &ContrailAccessGetter{
		readService: rs,
		log:         logutil.NewLogger("rbac-access-getter"),
	}
}

// GetAPIAccessLists returns APIAccessLists in RBAC terms.
func (ag ContrailAccessGetter) GetAPIAccessLists(ctx context.Context) []*rbac.APIAccessList {
	result, err := ag.readService.ListAPIAccessList(
		ctx,
		&services.ListAPIAccessListRequest{
			Spec: &baseservices.ListSpec{},
		},
	)
	if err != nil {
		// TODO(dfurman): return error instead
		ag.log.WithError(err).Debug("Failed to retrieve API access lists from read service")
		return nil
	}
	return models.APIAccessLists(result.APIAccessLists).ToRBAC()
}

type hasPerms2 interface {
	GetPerms2() *models.PermType2
}

// GetPermissions returns PermType2 in RBAC terms.
func (ag ContrailAccessGetter) GetPermissions(ctx context.Context, typeName, uuid string) *rbac.PermType2 {
	object, err := services.GetObject(ctx, ag.readService, typeName, uuid)
	if err != nil {
		// TODO(dfurman): return error instead
		ag.log.WithError(err).WithFields(logrus.Fields{
			"type-name": typeName,
			"uuid":      uuid,
		}).Debug("Failed to retrieve object from read service")
		return nil
	}

	objectWithPerms2, ok := object.(hasPerms2)
	if !ok {
		// TODO(dfurman): return error instead
		ag.log.WithFields(logrus.Fields{
			"type-name": typeName,
			"uuid":      uuid,
		}).Debug("Object does not contain Perms2 needed for RBAC")
		return nil
	}
	return objectWithPerms2.GetPerms2().ToRBAC()
}
