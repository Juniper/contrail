package types

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateProject creates a project and ensures a default application policy set for it.
func (sv *ContrailTypeLogicService) CreateProject(
	ctx context.Context, request *services.CreateProjectRequest,
) (response *services.CreateProjectResponse, err error) {

	err = sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			response, err = sv.BaseService.CreateProject(ctx, request)
			if err != nil {
				return err
			}

			return sv.ensureDefaultApplicationPolicySet(ctx, request.Project)
		})

	return response, err
}

// UpdateProject validates the request and updates the project.
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

			response, err = sv.BaseService.UpdateProject(ctx, request)
			return err
		})

	return response, err
}

// DeleteProject deletes the project with its default application policy set.
func (sv *ContrailTypeLogicService) DeleteProject(
	ctx context.Context, request *services.DeleteProjectRequest,
) (*services.DeleteProjectResponse, error) {
	var response *services.DeleteProjectResponse

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			//TODO: pre dbe delete
			project, err := sv.getProject(ctx, request.GetID())
			if err != nil {
				return err
			}

			if err = sv.deleteDefaultApplicationPolicySet(ctx, project); err != nil {
				return err
			}

			response, err = sv.BaseService.DeleteProject(ctx, request)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) getProject(ctx context.Context, id string) (*models.Project, error) {
	projectRes, err := sv.ReadService.GetProject(ctx, &services.GetProjectRequest{
		ID: id,
	})
	return projectRes.GetProject(), err
}

func (sv *ContrailTypeLogicService) checkVxlanConfig(
	currentProject *models.Project, request *services.UpdateProjectRequest,
) error {
	requestedProject := request.GetProject()

	fm := request.GetFieldMask()
	isVxlanChangeRequested := basemodels.FieldMaskContains(&fm, models.ProjectFieldVxlanRouting)
	if !isVxlanChangeRequested {
		return nil
	}

	willVxlanChange := currentProject.GetVxlanRouting() != requestedProject.GetVxlanRouting()
	if !willVxlanChange {
		return nil
	}

	areLogicalRoutersAlreadyConfigured := len(currentProject.GetLogicalRouters()) > 0
	if areLogicalRoutersAlreadyConfigured {
		return errutil.ErrorBadRequest("VxLAN Routing update for project " + currentProject.GetUUID() +
			" cannot be done when Logical Routers are configured")
	}

	return nil
}

func (sv *ContrailTypeLogicService) ensureDefaultApplicationPolicySet(
	ctx context.Context, project *models.Project,
) error {
	apsName := basemodels.DefaultNameForKind(models.KindApplicationPolicySet)

	response, err := sv.WriteService.CreateApplicationPolicySet(
		ctx, &services.CreateApplicationPolicySetRequest{
			ApplicationPolicySet: &models.ApplicationPolicySet{
				FQName:          basemodels.ChildFQName(project.GetFQName(), apsName),
				ParentType:      project.Kind(),
				ParentUUID:      project.GetUUID(),
				Name:            apsName,
				DisplayName:     apsName,
				AllApplications: true,
			},
		})
	if errutil.IsConflict(err) {
		return nil // object already exists - do nothing
	} else if err != nil {
		return errors.Wrap(err, "failed to create default application policy set for project")
	}

	project.ApplicationPolicySetRefs = append(
		project.ApplicationPolicySetRefs,
		&models.ProjectApplicationPolicySetRef{
			UUID: response.GetApplicationPolicySet().GetUUID(),
			To:   response.GetApplicationPolicySet().GetFQName(),
		},
	)
	_, err = sv.WriteService.UpdateProject(
		ctx, &services.UpdateProjectRequest{
			Project:   project,
			FieldMask: types.FieldMask{Paths: []string{models.ProjectFieldApplicationPolicySetRefs}},
		},
	)

	return err
}

func (sv *ContrailTypeLogicService) deleteDefaultApplicationPolicySet(
	ctx context.Context, project *models.Project,
) error {
	// delete aps refs to make default application policy set deletion possible
	project.ApplicationPolicySetRefs = project.ApplicationPolicySetRefs[:0]
	_, err := sv.WriteService.UpdateProject(
		ctx, &services.UpdateProjectRequest{
			Project:   project,
			FieldMask: types.FieldMask{Paths: []string{models.ProjectFieldApplicationPolicySetRefs}},
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to delete application policy set refs")
	}

	defaultAPSName := basemodels.DefaultNameForKind(models.KindApplicationPolicySet)

	for _, aps := range project.GetApplicationPolicySets() {
		if aps.GetName() == defaultAPSName {
			_, err := sv.WriteService.DeleteApplicationPolicySet(
				ctx, &services.DeleteApplicationPolicySetRequest{ID: aps.UUID},
			)
			if err != nil {
				return errors.Wrap(err, "failed to delete child application policy set")
			}
			return nil
		}
	}
	return nil
}
