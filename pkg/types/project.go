package types

import (
	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
)

const vxlanConfigField = "VxlanRouting"

// CreateProject do checks for create project.
func (service *ContrailTypeLogicService) CreateProject(ctx context.Context, request *models.CreateProjectRequest) (response *models.CreateProjectResponse, err error) {

	err = db.DoInTransaction(
		ctx, service.DB.DB(),
		func(ctx context.Context) error {

			response, err = service.Next().CreateProject(ctx, request)
			//TODO: ensure default aplication policy set

			return err
		})

	return response, err
}

// UpdateProject do checks for update project.
func (service *ContrailTypeLogicService) UpdateProject(ctx context.Context, request *models.UpdateProjectRequest) (response *models.UpdateProjectResponse, err error) {
	id := request.GetProject().GetUUID()

	err = db.DoInTransaction(
		ctx, service.DB.DB(),
		func(ctx context.Context) error {
			currentProject, inerr := service.getProject(ctx, id)
			if inerr != nil {
				return inerr
			}
			if inerr = checkVxlanConfig(currentProject, request); inerr != nil {
				return inerr
			}
			//TODO: check enable_security_policy_draft

			response, inerr = service.Next().UpdateProject(ctx, request)
			return inerr
		})

	return response, err
}

// DeleteProject do checks for delete project.
func (service *ContrailTypeLogicService) DeleteProject(ctx context.Context, request *models.DeleteProjectRequest) (response *models.DeleteProjectResponse, err error) {
	err = db.DoInTransaction(
		ctx, service.DB.DB(),
		func(ctx context.Context) error {
			//TODO: pre dbe delete
			response, err = service.Next().DeleteProject(ctx, request)
			return err
		})

	return response, err
}

func (service *ContrailTypeLogicService) getProject(ctx context.Context, id string) (*models.Project, error) {
	projectRes, err := service.DB.GetProject(ctx, &models.GetProjectRequest{
		ID: id,
	})
	return projectRes.GetProject(), err
}

func checkVxlanConfig(currentProject *models.Project, request *models.UpdateProjectRequest) error {
	requestedProject := request.GetProject()

	fm := request.GetFieldMask()
	isVxlanChangeRequested := common.ContainsString(fm.GetPaths(), vxlanConfigField)
	if !isVxlanChangeRequested {
		return nil
	}

	willVxlanChange := currentProject.GetVxlanRouting() != requestedProject.GetVxlanRouting()
	if !willVxlanChange {
		return nil
	}

	areLogicalRoutersAlreadyConfigured := len(currentProject.GetLogicalRouters()) > 0
	if areLogicalRoutersAlreadyConfigured {
		return common.ErrorBadRequest("VxLAN Routing update cannot be done when Logical Routers are configured")
	}

	return nil
}