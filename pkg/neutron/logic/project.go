package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

func getProject(ctx context.Context, rp RequestParameters) (*models.Project, error) {
	projectID, err := neutronIDToVncUUID(rp.RequestContext.TenantID)
	if err != nil {
		return nil, err
	}

	projectResponse, err := rp.ReadService.GetProject(
		ctx,
		&services.GetProjectRequest{
			ID: projectID,
		},
	)
	if err != nil {
		return nil, err

	}

	return projectResponse.GetProject(), nil
}
