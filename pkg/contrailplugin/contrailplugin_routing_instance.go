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

// RoutingInstanceIntent
//   A struct to store attributes related to RoutingInstance
//   needed by Intent Compiler
type RoutingInstanceIntent struct {
	Uuid string
}

// EvaluateRoutingInstance - evaluates the RoutingInstance
func EvaluateRoutingInstance(obj interface{}) {
	resourceObj := obj.(RoutingInstanceIntent)
	log.Println("EvaluateRoutingInstance Called ", resourceObj)
}

// CreateRoutingInstance handles create request
func (service *PluginService) CreateRoutingInstance(ctx context.Context, request *services.CreateRoutingInstanceRequest) (*services.CreateRoutingInstanceResponse, error) {
	log.Println(" CreateRoutingInstance Entered")

	obj := request.GetRoutingInstance()

	intentObj := RoutingInstanceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RoutingInstanceIntent"); !ok {
		compilationif.ObjsCache.Store("RoutingInstanceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("RoutingInstanceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateRoutingInstance", objMap.(*sync.Map))

	EvaluateDependencies(obj, "RoutingInstance")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateRoutingInstance(ctx, request)
}

// UpdateRoutingInstance handles update request
func (service *PluginService) UpdateRoutingInstance(ctx context.Context, request *services.UpdateRoutingInstanceRequest) (*services.UpdateRoutingInstanceResponse, error) {
	log.Println(" UpdateRoutingInstance ENTERED")

	obj := request.GetRoutingInstance()

	intentObj := RoutingInstanceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RoutingInstanceIntent"); !ok {
		compilationif.ObjsCache.Store("RoutingInstanceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "RoutingInstance")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateRoutingInstance", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateRoutingInstance(ctx, request)
}

// DeleteRoutingInstance handles delete request
func (service *PluginService) DeleteRoutingInstance(ctx context.Context, request *services.DeleteRoutingInstanceRequest) (*services.DeleteRoutingInstanceResponse, error) {
	log.Println(" DeleteRoutingInstance ENTERED")

	objUUID := request.GetID()

	//intentObj := RoutingInstanceIntent {
	//RoutingInstance: *obj,
	//}

	//EvaluateDependencies(intentObj, "RoutingInstance")

	objMap, ok := compilationif.ObjsCache.Load("RoutingInstanceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteRoutingInstance", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteRoutingInstance(ctx, request)
}

// GetRoutingInstance handles get request
func (service *PluginService) GetRoutingInstance(ctx context.Context, request *services.GetRoutingInstanceRequest) (*services.GetRoutingInstanceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("RoutingInstance")
	if !ok {
		return nil, errors.New("RoutingInstance get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("RoutingInstance get failed ")
	}

	response := &services.GetRoutingInstanceResponse{
		RoutingInstance: obj.(*models.RoutingInstance),
	}
	return response, nil
}
