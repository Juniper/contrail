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

// ServiceHealthCheckIntent
//   A struct to store attributes related to ServiceHealthCheck
//   needed by Intent Compiler
type ServiceHealthCheckIntent struct {
	Uuid string
}

// EvaluateServiceHealthCheck - evaluates the ServiceHealthCheck
func EvaluateServiceHealthCheck(obj interface{}) {
	resourceObj := obj.(ServiceHealthCheckIntent)
	log.Println("EvaluateServiceHealthCheck Called ", resourceObj)
}

// CreateServiceHealthCheck handles create request
func (service *PluginService) CreateServiceHealthCheck(ctx context.Context, request *services.CreateServiceHealthCheckRequest) (*services.CreateServiceHealthCheckResponse, error) {
	log.Println(" CreateServiceHealthCheck Entered")

	obj := request.GetServiceHealthCheck()

	intentObj := ServiceHealthCheckIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceHealthCheckIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceHealthCheckIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceHealthCheckIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceHealthCheck", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceHealthCheck")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceHealthCheck(ctx, request)
}

// UpdateServiceHealthCheck handles update request
func (service *PluginService) UpdateServiceHealthCheck(ctx context.Context, request *services.UpdateServiceHealthCheckRequest) (*services.UpdateServiceHealthCheckResponse, error) {
	log.Println(" UpdateServiceHealthCheck ENTERED")

	obj := request.GetServiceHealthCheck()

	intentObj := ServiceHealthCheckIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceHealthCheckIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceHealthCheckIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceHealthCheck")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceHealthCheck", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceHealthCheck(ctx, request)
}

// DeleteServiceHealthCheck handles delete request
func (service *PluginService) DeleteServiceHealthCheck(ctx context.Context, request *services.DeleteServiceHealthCheckRequest) (*services.DeleteServiceHealthCheckResponse, error) {
	log.Println(" DeleteServiceHealthCheck ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceHealthCheckIntent {
	//ServiceHealthCheck: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceHealthCheck")

	objMap, ok := compilationif.ObjsCache.Load("ServiceHealthCheckIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceHealthCheck", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceHealthCheck(ctx, request)
}

// GetServiceHealthCheck handles get request
func (service *PluginService) GetServiceHealthCheck(ctx context.Context, request *services.GetServiceHealthCheckRequest) (*services.GetServiceHealthCheckResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceHealthCheck")
	if !ok {
		return nil, errors.New("ServiceHealthCheck get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceHealthCheck get failed ")
	}

	response := &services.GetServiceHealthCheckResponse{
		ServiceHealthCheck: obj.(*models.ServiceHealthCheck),
	}
	return response, nil
}
