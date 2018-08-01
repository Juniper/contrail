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

// ServiceTemplateIntent
//   A struct to store attributes related to ServiceTemplate
//   needed by Intent Compiler
type ServiceTemplateIntent struct {
	Uuid string
}

// EvaluateServiceTemplate - evaluates the ServiceTemplate
func EvaluateServiceTemplate(obj interface{}) {
	resourceObj := obj.(ServiceTemplateIntent)
	log.Println("EvaluateServiceTemplate Called ", resourceObj)
}

// CreateServiceTemplate handles create request
func (service *PluginService) CreateServiceTemplate(ctx context.Context, request *services.CreateServiceTemplateRequest) (*services.CreateServiceTemplateResponse, error) {
	log.Println(" CreateServiceTemplate Entered")

	obj := request.GetServiceTemplate()

	intentObj := ServiceTemplateIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceTemplateIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceTemplateIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceTemplateIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceTemplate", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceTemplate")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceTemplate(ctx, request)
}

// UpdateServiceTemplate handles update request
func (service *PluginService) UpdateServiceTemplate(ctx context.Context, request *services.UpdateServiceTemplateRequest) (*services.UpdateServiceTemplateResponse, error) {
	log.Println(" UpdateServiceTemplate ENTERED")

	obj := request.GetServiceTemplate()

	intentObj := ServiceTemplateIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceTemplateIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceTemplateIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceTemplate")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceTemplate", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceTemplate(ctx, request)
}

// DeleteServiceTemplate handles delete request
func (service *PluginService) DeleteServiceTemplate(ctx context.Context, request *services.DeleteServiceTemplateRequest) (*services.DeleteServiceTemplateResponse, error) {
	log.Println(" DeleteServiceTemplate ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceTemplateIntent {
	//ServiceTemplate: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceTemplate")

	objMap, ok := compilationif.ObjsCache.Load("ServiceTemplateIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceTemplate", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceTemplate(ctx, request)
}

// GetServiceTemplate handles get request
func (service *PluginService) GetServiceTemplate(ctx context.Context, request *services.GetServiceTemplateRequest) (*services.GetServiceTemplateResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceTemplate")
	if !ok {
		return nil, errors.New("ServiceTemplate get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceTemplate get failed ")
	}

	response := &services.GetServiceTemplateResponse{
		ServiceTemplate: obj.(*models.ServiceTemplate),
	}
	return response, nil
}
