package services

import (
	"context"

	"github.com/Juniper/asf/pkg/rbac"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/models"
)

type ContrailAccessGetter struct {
	Service *ContrailService
}

func (ag ContrailAccessGetter) GetAPIAccessLists(ctx context.Context) []*rbac.APIAccessList {
	listRequest := &ListAPIAccessListRequest{
		Spec: &baseservices.ListSpec{},
	}
	if result, err := ag.Service.ListAPIAccessList(ctx, listRequest); err == nil {
		return models.APIAccessLists(result.APIAccessLists).ToRBAC()
	}
	return nil
}

type objectType interface {
	GetPerms2() *models.PermType2
}

func (ag ContrailAccessGetter) GetPermissions(ctx context.Context, typeName, uuid string) *rbac.PermType2 {
	object, err := GetObject(ctx, ag.Service.DBService, typeName, uuid)
	if err != nil {
		return nil
	}
	if specificObject, ok := object.(objectType); ok {
		return specificObject.GetPerms2().ToRBAC()
	}
	return nil
}
