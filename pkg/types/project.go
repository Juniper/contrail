package types

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

// CreateProject do checks for create project.
func (sv *ContrailTypeLogicService) CreateProject(
	ctx context.Context, request *services.CreateProjectRequest,
) (response *services.CreateProjectResponse, err error) {

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {

			response, err = sv.Next().CreateProject(ctx, request)

			return sv.ensureDefaultApplicationPolicySet(ctx, request.Project)
		})

	return response, err
}

// UpdateProject do checks for update project.
func (sv *ContrailTypeLogicService) UpdateProject(
	ctx context.Context, request *services.UpdateProjectRequest,
) (response *services.UpdateProjectResponse, err error) {
	id := request.GetProject().GetUUID()

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var currentProject *models.Project
			currentProject, err = sv.getProject(ctx, id)
			if err != nil {
				return err
			}
			if err = sv.checkVxlanConfig(currentProject, request); err != nil {
				return err
			}
			//TODO: check enable_security_policy_draft

			response, err = sv.Next().UpdateProject(ctx, request)
			return err
		})

	return response, err
}

// DeleteProject do checks for delete project.
func (sv *ContrailTypeLogicService) DeleteProject(
	ctx context.Context, request *services.DeleteProjectRequest,
) (response *services.DeleteProjectResponse, err error) {
	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			//TODO: pre dbe delete
			response, err = sv.Next().DeleteProject(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getProject(ctx context.Context, id string) (*models.Project, error) {
	projectRes, err := sv.DataService.GetProject(ctx, &services.GetProjectRequest{
		ID: id,
	})
	return projectRes.GetProject(), err
}

func (sv *ContrailTypeLogicService) checkVxlanConfig(
	currentProject *models.Project, request *services.UpdateProjectRequest,
) error {
	requestedProject := request.GetProject()

	fm := request.GetFieldMask()
	isVxlanChangeRequested := common.ContainsString(fm.GetPaths(), models.ProjectPropertyIDVxlanRouting)
	if !isVxlanChangeRequested {
		return nil
	}

	willVxlanChange := currentProject.GetVxlanRouting() != requestedProject.GetVxlanRouting()
	if !willVxlanChange {
		return nil
	}

	areLogicalRoutersAlreadyConfigured := len(currentProject.GetLogicalRouters()) > 0
	if areLogicalRoutersAlreadyConfigured {
		return common.ErrorBadRequest("VxLAN Routing update for project " + currentProject.GetUUID() +
			" cannot be done when Logical Routers are configured")
	}

	return nil
}

func (sv *ContrailTypeLogicService) ensureDefaultApplicationPolicySet(ctx context.Context, project *models.Project) error {
	apsName := fmt.Sprintf("default-%s", models.KindApplicationPolicySet)

	aps := models.MakeApplicationPolicySet()
	aps.FQName = append(project.FQName, apsName)
	aps.ParentType = models.KindProject
	aps.ParentUUID = project.UUID
	aps.Name = apsName
	aps.DisplayName = apsName
	aps.AllApplications = true

	response, err := sv.DataService.CreateApplicationPolicySet(ctx, &services.CreateApplicationPolicySetRequest{
		ApplicationPolicySet: aps,
	})
	if err != nil {
		if !common.IsConflict(err) {
			return errors.Wrap(err, "unexpected error creating default application policy set for project")
		}
		// object already exists - do nothing
		return nil
	}
	// new object - create ref

	project.ApplicationPolicySetRefs = append(
		project.ApplicationPolicySetRefs,
		&models.ProjectApplicationPolicySetRef{UUID: response.ApplicationPolicySet.UUID},
	)
	return nil
}
