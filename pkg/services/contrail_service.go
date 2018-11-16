package services

import (
	"context"

	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/models"
)

func (service *ContrailService) getAllAPIAccessLists(ctx context.Context) []*models.APIAccessList {

	noAuthCtx := auth.NoAuth(ctx)
	// Use a context with No auth for internal calls
	result, err := service.ListAPIAccessList(noAuthCtx,  &ListAPIAccessListRequest{})
	if err != nil {
		return nil
	}

	return result.APIAccessLists
}
