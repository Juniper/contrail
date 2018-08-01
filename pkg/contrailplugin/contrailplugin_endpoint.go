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

// EndpointIntent
//   A struct to store attributes related to Endpoint
//   needed by Intent Compiler
type EndpointIntent struct {
	Uuid string
}

// EvaluateEndpoint - evaluates the Endpoint
func EvaluateEndpoint(obj interface{}) {
	resourceObj := obj.(EndpointIntent)
	log.Println("EvaluateEndpoint Called ", resourceObj)
}

// CreateEndpoint handles create request
func (service *PluginService) CreateEndpoint(ctx context.Context, request *services.CreateEndpointRequest) (*services.CreateEndpointResponse, error) {
	log.Println(" CreateEndpoint Entered")

	obj := request.GetEndpoint()

	intentObj := EndpointIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("EndpointIntent"); !ok {
		compilationif.ObjsCache.Store("EndpointIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("EndpointIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateEndpoint", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Endpoint")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateEndpoint(ctx, request)
}

// UpdateEndpoint handles update request
func (service *PluginService) UpdateEndpoint(ctx context.Context, request *services.UpdateEndpointRequest) (*services.UpdateEndpointResponse, error) {
	log.Println(" UpdateEndpoint ENTERED")

	obj := request.GetEndpoint()

	intentObj := EndpointIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("EndpointIntent"); !ok {
		compilationif.ObjsCache.Store("EndpointIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Endpoint")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateEndpoint", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateEndpoint(ctx, request)
}

// DeleteEndpoint handles delete request
func (service *PluginService) DeleteEndpoint(ctx context.Context, request *services.DeleteEndpointRequest) (*services.DeleteEndpointResponse, error) {
	log.Println(" DeleteEndpoint ENTERED")

	objUUID := request.GetID()

	//intentObj := EndpointIntent {
	//Endpoint: *obj,
	//}

	//EvaluateDependencies(intentObj, "Endpoint")

	objMap, ok := compilationif.ObjsCache.Load("EndpointIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteEndpoint", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteEndpoint(ctx, request)
}

// GetEndpoint handles get request
func (service *PluginService) GetEndpoint(ctx context.Context, request *services.GetEndpointRequest) (*services.GetEndpointResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Endpoint")
	if !ok {
		return nil, errors.New("Endpoint get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Endpoint get failed ")
	}

	response := &services.GetEndpointResponse{
		Endpoint: obj.(*models.Endpoint),
	}
	return response, nil
}
