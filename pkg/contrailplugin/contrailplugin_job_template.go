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

// JobTemplateIntent
//   A struct to store attributes related to JobTemplate
//   needed by Intent Compiler
type JobTemplateIntent struct {
	Uuid string
}

// EvaluateJobTemplate - evaluates the JobTemplate
func EvaluateJobTemplate(obj interface{}) {
	resourceObj := obj.(JobTemplateIntent)
	log.Println("EvaluateJobTemplate Called ", resourceObj)
}

// CreateJobTemplate handles create request
func (service *PluginService) CreateJobTemplate(ctx context.Context, request *services.CreateJobTemplateRequest) (*services.CreateJobTemplateResponse, error) {
	log.Println(" CreateJobTemplate Entered")

	obj := request.GetJobTemplate()

	intentObj := JobTemplateIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("JobTemplateIntent"); !ok {
		compilationif.ObjsCache.Store("JobTemplateIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("JobTemplateIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateJobTemplate", objMap.(*sync.Map))

	EvaluateDependencies(obj, "JobTemplate")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateJobTemplate(ctx, request)
}

// UpdateJobTemplate handles update request
func (service *PluginService) UpdateJobTemplate(ctx context.Context, request *services.UpdateJobTemplateRequest) (*services.UpdateJobTemplateResponse, error) {
	log.Println(" UpdateJobTemplate ENTERED")

	obj := request.GetJobTemplate()

	intentObj := JobTemplateIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("JobTemplateIntent"); !ok {
		compilationif.ObjsCache.Store("JobTemplateIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "JobTemplate")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateJobTemplate", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateJobTemplate(ctx, request)
}

// DeleteJobTemplate handles delete request
func (service *PluginService) DeleteJobTemplate(ctx context.Context, request *services.DeleteJobTemplateRequest) (*services.DeleteJobTemplateResponse, error) {
	log.Println(" DeleteJobTemplate ENTERED")

	objUUID := request.GetID()

	//intentObj := JobTemplateIntent {
	//JobTemplate: *obj,
	//}

	//EvaluateDependencies(intentObj, "JobTemplate")

	objMap, ok := compilationif.ObjsCache.Load("JobTemplateIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteJobTemplate", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteJobTemplate(ctx, request)
}

// GetJobTemplate handles get request
func (service *PluginService) GetJobTemplate(ctx context.Context, request *services.GetJobTemplateRequest) (*services.GetJobTemplateResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("JobTemplate")
	if !ok {
		return nil, errors.New("JobTemplate get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("JobTemplate get failed ")
	}

	response := &services.GetJobTemplateResponse{
		JobTemplate: obj.(*models.JobTemplate),
	}
	return response, nil
}
