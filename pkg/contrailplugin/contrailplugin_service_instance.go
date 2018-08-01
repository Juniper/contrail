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

// ServiceInstanceIntent
//   A struct to store attributes related to ServiceInstance
//   needed by Intent Compiler
type ServiceInstanceIntent struct {
	Uuid string
}

// EvaluateServiceInstance - evaluates the ServiceInstance
func EvaluateServiceInstance(obj interface{}) {
	resourceObj := obj.(ServiceInstanceIntent)
	log.Println("EvaluateServiceInstance Called ", resourceObj)
}

// CreateServiceInstance handles create request
func (service *PluginService) CreateServiceInstance(ctx context.Context, request *services.CreateServiceInstanceRequest) (*services.CreateServiceInstanceResponse, error) {
	log.Println(" CreateServiceInstance Entered")

	obj := request.GetServiceInstance()

	intentObj := ServiceInstanceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceInstanceIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceInstanceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceInstanceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceInstance", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceInstance")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceInstance(ctx, request)
}

// UpdateServiceInstance handles update request
func (service *PluginService) UpdateServiceInstance(ctx context.Context, request *services.UpdateServiceInstanceRequest) (*services.UpdateServiceInstanceResponse, error) {
	log.Println(" UpdateServiceInstance ENTERED")

	obj := request.GetServiceInstance()

	intentObj := ServiceInstanceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceInstanceIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceInstanceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceInstance")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceInstance", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceInstance(ctx, request)
}

// DeleteServiceInstance handles delete request
func (service *PluginService) DeleteServiceInstance(ctx context.Context, request *services.DeleteServiceInstanceRequest) (*services.DeleteServiceInstanceResponse, error) {
	log.Println(" DeleteServiceInstance ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceInstanceIntent {
	//ServiceInstance: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceInstance")

	objMap, ok := compilationif.ObjsCache.Load("ServiceInstanceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceInstance", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceInstance(ctx, request)
}

// GetServiceInstance handles get request
func (service *PluginService) GetServiceInstance(ctx context.Context, request *services.GetServiceInstanceRequest) (*services.GetServiceInstanceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceInstance")
	if !ok {
		return nil, errors.New("ServiceInstance get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceInstance get failed ")
	}

	response := &services.GetServiceInstanceResponse{
		ServiceInstance: obj.(*models.ServiceInstance),
	}
	return response, nil
}
