package services

import (
	"context"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	//	"time"
)

func (service *ContrailService) getAllAPIAccessLists(ctx context.Context) []*models.APIAccessList {

	listSpec := &baseservices.ListSpec{}
	listRequest := &ListAPIAccessListRequest{
		Spec: listSpec,
	}

	noAuthCtx := auth.NoAuth(ctx)
	// Use a context with No auth for internal calls
	result, err := service.ListAPIAccessList(noAuthCtx, listRequest)
	if err != nil {
		return nil
	}

	return result.APIAccessLists
}
