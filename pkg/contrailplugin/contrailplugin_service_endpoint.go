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

// ServiceEndpointIntent
//   A struct to store attributes related to ServiceEndpoint
//   needed by Intent Compiler
type ServiceEndpointIntent struct {
	Uuid string
}

// EvaluateServiceEndpoint - evaluates the ServiceEndpoint
func EvaluateServiceEndpoint(obj interface{}) {
	resourceObj := obj.(ServiceEndpointIntent)
	log.Println("EvaluateServiceEndpoint Called ", resourceObj)
}

// CreateServiceEndpoint handles create request
func (service *PluginService) CreateServiceEndpoint(ctx context.Context, request *services.CreateServiceEndpointRequest) (*services.CreateServiceEndpointResponse, error) {
	log.Println(" CreateServiceEndpoint Entered")

	obj := request.GetServiceEndpoint()

	intentObj := ServiceEndpointIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceEndpointIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceEndpointIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceEndpointIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceEndpoint", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceEndpoint")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceEndpoint(ctx, request)
}

// UpdateServiceEndpoint handles update request
func (service *PluginService) UpdateServiceEndpoint(ctx context.Context, request *services.UpdateServiceEndpointRequest) (*services.UpdateServiceEndpointResponse, error) {
	log.Println(" UpdateServiceEndpoint ENTERED")

	obj := request.GetServiceEndpoint()

	intentObj := ServiceEndpointIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceEndpointIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceEndpointIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceEndpoint")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceEndpoint", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceEndpoint(ctx, request)
}

// DeleteServiceEndpoint handles delete request
func (service *PluginService) DeleteServiceEndpoint(ctx context.Context, request *services.DeleteServiceEndpointRequest) (*services.DeleteServiceEndpointResponse, error) {
	log.Println(" DeleteServiceEndpoint ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceEndpointIntent {
	//ServiceEndpoint: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceEndpoint")

	objMap, ok := compilationif.ObjsCache.Load("ServiceEndpointIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceEndpoint", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceEndpoint(ctx, request)
}

// GetServiceEndpoint handles get request
func (service *PluginService) GetServiceEndpoint(ctx context.Context, request *services.GetServiceEndpointRequest) (*services.GetServiceEndpointResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceEndpoint")
	if !ok {
		return nil, errors.New("ServiceEndpoint get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceEndpoint get failed ")
	}

	response := &services.GetServiceEndpointResponse{
		ServiceEndpoint: obj.(*models.ServiceEndpoint),
	}
	return response, nil
}
