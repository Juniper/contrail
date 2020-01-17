package services

import (
	"context"

	"github.com/Juniper/asf/pkg/rbac"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
)

// ContrailAccessGetter is used to get APIAccessLists and resources' Perms2
type ContrailAccessGetter struct {
	ReadService ReadService
}

// GetAPIAccessLists returns APIAccessLists in rbac terms
func (ag ContrailAccessGetter) GetAPIAccessLists(ctx context.Context) []*rbac.APIAccessList {
	listRequest := &ListAPIAccessListRequest{
		Spec: &baseservices.ListSpec{},
	}
	if result, err := ag.ReadService.ListAPIAccessList(ctx, listRequest); err == nil {
		return models.APIAccessLists(result.APIAccessLists).ToRBAC()
	}
	return nil
}

type hasPerms2 interface {
	GetPerms2() *models.PermType2
}

// GetPermissions returns PermType2 in rbac terms
func (ag ContrailAccessGetter) GetPermissions(ctx context.Context, typeName, uuid string) *rbac.PermType2 {
	object, err := GetObject(ctx, ag.ReadService, typeName, uuid)
	if err != nil {
		return nil
	}
	if objectWithPerms2, ok := object.(hasPerms2); ok {
		return objectWithPerms2.GetPerms2().ToRBAC()
	}
	return nil
}
