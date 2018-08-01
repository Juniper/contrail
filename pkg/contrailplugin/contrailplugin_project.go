// nolint
package contrailplugin

import (
	"context"
	"errors"
	"sync"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
)

// ProjectIntent
//   A struct to store attributes related to Project
//   needed by Intent Compiler
type ProjectIntent struct {
	Uuid string
}

// EvaluateProject - evaluates the Project
func EvaluateProject(obj interface{}) {
	resourceObj := obj.(ProjectIntent)
	log.Println("EvaluateProject Called ", resourceObj)
}

// CreateProject handles create request
func (service *PluginService) CreateProject(ctx context.Context, request *services.CreateProjectRequest) (*services.CreateProjectResponse, error) {
	log.Println(" CreateProject Entered")

	obj := request.GetProject()

	intentObj := ProjectIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ProjectIntent"); !ok {
		compilationif.ObjsCache.Store("ProjectIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ProjectIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateProject", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Project")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateProject(ctx, request)
}

// UpdateProject handles update request
func (service *PluginService) UpdateProject(ctx context.Context, request *services.UpdateProjectRequest) (*services.UpdateProjectResponse, error) {
	log.Println(" UpdateProject ENTERED")

	obj := request.GetProject()

	intentObj := ProjectIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ProjectIntent"); !ok {
		compilationif.ObjsCache.Store("ProjectIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Project")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateProject", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateProject(ctx, request)
}

// DeleteProject handles delete request
func (service *PluginService) DeleteProject(ctx context.Context, request *services.DeleteProjectRequest) (*services.DeleteProjectResponse, error) {
	log.Println(" DeleteProject ENTERED")

	objUUID := request.GetID()

	//intentObj := ProjectIntent {
	//Project: *obj,
	//}

	//EvaluateDependencies(intentObj, "Project")

	objMap, ok := compilationif.ObjsCache.Load("ProjectIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteProject", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteProject(ctx, request)
}

// GetProject handles get request
func (service *PluginService) GetProject(ctx context.Context, request *services.GetProjectRequest) (*services.GetProjectResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Project")
	if !ok {
		return nil, errors.New("Project get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Project get failed ")
	}

	response := &services.GetProjectResponse{
		Project: obj.(*models.Project),
	}
	return response, nil
}
