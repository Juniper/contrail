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

// ServiceConnectionModuleIntent
//   A struct to store attributes related to ServiceConnectionModule
//   needed by Intent Compiler
type ServiceConnectionModuleIntent struct {
	Uuid string
}

// EvaluateServiceConnectionModule - evaluates the ServiceConnectionModule
func EvaluateServiceConnectionModule(obj interface{}) {
	resourceObj := obj.(ServiceConnectionModuleIntent)
	log.Println("EvaluateServiceConnectionModule Called ", resourceObj)
}

// CreateServiceConnectionModule handles create request
func (service *PluginService) CreateServiceConnectionModule(ctx context.Context, request *services.CreateServiceConnectionModuleRequest) (*services.CreateServiceConnectionModuleResponse, error) {
	log.Println(" CreateServiceConnectionModule Entered")

	obj := request.GetServiceConnectionModule()

	intentObj := ServiceConnectionModuleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceConnectionModuleIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceConnectionModuleIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceConnectionModuleIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceConnectionModule", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceConnectionModule")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceConnectionModule(ctx, request)
}

// UpdateServiceConnectionModule handles update request
func (service *PluginService) UpdateServiceConnectionModule(ctx context.Context, request *services.UpdateServiceConnectionModuleRequest) (*services.UpdateServiceConnectionModuleResponse, error) {
	log.Println(" UpdateServiceConnectionModule ENTERED")

	obj := request.GetServiceConnectionModule()

	intentObj := ServiceConnectionModuleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceConnectionModuleIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceConnectionModuleIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceConnectionModule")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceConnectionModule", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceConnectionModule(ctx, request)
}

// DeleteServiceConnectionModule handles delete request
func (service *PluginService) DeleteServiceConnectionModule(ctx context.Context, request *services.DeleteServiceConnectionModuleRequest) (*services.DeleteServiceConnectionModuleResponse, error) {
	log.Println(" DeleteServiceConnectionModule ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceConnectionModuleIntent {
	//ServiceConnectionModule: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceConnectionModule")

	objMap, ok := compilationif.ObjsCache.Load("ServiceConnectionModuleIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceConnectionModule", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceConnectionModule(ctx, request)
}

// GetServiceConnectionModule handles get request
func (service *PluginService) GetServiceConnectionModule(ctx context.Context, request *services.GetServiceConnectionModuleRequest) (*services.GetServiceConnectionModuleResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceConnectionModule")
	if !ok {
		return nil, errors.New("ServiceConnectionModule get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceConnectionModule get failed ")
	}

	response := &services.GetServiceConnectionModuleResponse{
		ServiceConnectionModule: obj.(*models.ServiceConnectionModule),
	}
	return response, nil
}
