package services

import (
	"context"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	//	"time"
)

func (service *ContrailService) listAPIAccessGetRules(ctx context.Context) []*models.APIAccessList {

	listSpec := &baseservices.ListSpec{}
	noAuthCtx := auth.NoAuth(ctx)

	listRequest := &ListAPIAccessListRequest{
		Spec: listSpec,
	}
	// Use a context with No auth for internal calls
	result, err := service.ListAPIAccessList(noAuthCtx, listRequest)
	if err != nil {
		return nil
	}

	apiAccessRules := result.APIAccessLists

	return apiAccessRules

}
