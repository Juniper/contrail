package types

import (
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"golang.org/x/net/context"
	google_protobuf1 "github.com/gogo/protobuf/types"
	"github.com/Juniper/contrail/pkg/common"
)

//CreateProject do checks for create project.
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

//UpdateProject do checks for update project.
func (service *ContrailTypeLogicService) UpdateProject(ctx context.Context, request *models.UpdateProjectRequest) (response *models.UpdateProjectResponse, err error) {
	err = db.DoInTransaction(
		ctx, service.DB.DB(),
		func(ctx context.Context) error {
			service.checkVxlanConfig(ctx, request)
			//TODO: check enable_security_policy_draft

			response, err = service.Next().UpdateProject(ctx, request)
			return err
		})

	return response, err
}

//DeleteProject do checks for delete project.
func (service *ContrailTypeLogicService) DeleteProject(ctx context.Context, request *models.DeleteProjectRequest) (response *models.DeleteProjectResponse, err error) {
	err = db.DoInTransaction(
		ctx, service.DB.DB(),
		func(ctx context.Context) error {
			//TODO: pre dbe delete
			response, err = service.Next().DeleteProject(ctx, request)
			//TODO: post dbe delete
			return err
		})

	return response, err
}


func (service *ContrailTypeLogicService) getProject(ctx context.Context, id string) (*models.Project, error) {
	var projectRes *models.GetProjectResponse
	projectRes, err := service.DB.GetProject(ctx, &models.GetProjectRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return projectRes.GetProject(), nil
}

func (service *ContrailTypeLogicService) checkVxlanConfig(ctx context.Context, request *models.UpdateProjectRequest) (ok bool, err error) {
	requestProject := request.GetProject()
	id := requestProject.GetUUID()
	currentProject, err := service.getProject(ctx, id)
	if err != nil {
		return false, err
	}
	fieldMask := request.GetFieldMask()
	if ok := isVxlanInFiledMask(fieldMask); !ok {
		return true, nil
	}
	if currentProject.GetVxlanRouting() == requestProject.GetVxlanRouting() {
		return true, nil
	}
	if len(currentProject.GetLogicalRouters()) != 0 {
		return false, common.ErrorBadRequest("VxLAN Routing update cannot be done when Logical Routers are configured")
	}
	return true, nil
}

func (service *ContrailTypeLogicService) ensureDefaultApplicationPolicySet(ctx context.Context, project_uuid string, poject_fq_name []string) error {
	//TODO: implement
	return nil
}

func isVxlanInFiledMask(fm google_protobuf1.FieldMask) bool {
	for _, m := range fm.GetPaths() {
		if m == "VxlanRouting" {
			return true
		}
	}
	return false
}