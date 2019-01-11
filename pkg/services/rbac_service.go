package services

import (
	"context"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

func (service *RBACService) getAllAPIAccessLists(ctx context.Context) []*models.APIAccessList {
	noAuthCtx := auth.NoAuth(ctx)
	listRequest := &ListAPIAccessListRequest{
		Spec: &baseservices.ListSpec{},
	}
	// Use a context with No auth for internal calls
	result, err := service.ReadService.ListAPIAccessList(noAuthCtx, listRequest)
	if err != nil {
		return nil
	}
	return result.APIAccessLists
}

