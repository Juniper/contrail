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

// PhysicalRouterIntent
//   A struct to store attributes related to PhysicalRouter
//   needed by Intent Compiler
type PhysicalRouterIntent struct {
	Uuid string
}

// EvaluatePhysicalRouter - evaluates the PhysicalRouter
func EvaluatePhysicalRouter(obj interface{}) {
	resourceObj := obj.(PhysicalRouterIntent)
	log.Println("EvaluatePhysicalRouter Called ", resourceObj)
}

// CreatePhysicalRouter handles create request
func (service *PluginService) CreatePhysicalRouter(ctx context.Context, request *services.CreatePhysicalRouterRequest) (*services.CreatePhysicalRouterResponse, error) {
	log.Println(" CreatePhysicalRouter Entered")

	obj := request.GetPhysicalRouter()

	intentObj := PhysicalRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PhysicalRouterIntent"); !ok {
		compilationif.ObjsCache.Store("PhysicalRouterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("PhysicalRouterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreatePhysicalRouter", objMap.(*sync.Map))

	EvaluateDependencies(obj, "PhysicalRouter")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreatePhysicalRouter(ctx, request)
}

// UpdatePhysicalRouter handles update request
func (service *PluginService) UpdatePhysicalRouter(ctx context.Context, request *services.UpdatePhysicalRouterRequest) (*services.UpdatePhysicalRouterResponse, error) {
	log.Println(" UpdatePhysicalRouter ENTERED")

	obj := request.GetPhysicalRouter()

	intentObj := PhysicalRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PhysicalRouterIntent"); !ok {
		compilationif.ObjsCache.Store("PhysicalRouterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "PhysicalRouter")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdatePhysicalRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdatePhysicalRouter(ctx, request)
}

// DeletePhysicalRouter handles delete request
func (service *PluginService) DeletePhysicalRouter(ctx context.Context, request *services.DeletePhysicalRouterRequest) (*services.DeletePhysicalRouterResponse, error) {
	log.Println(" DeletePhysicalRouter ENTERED")

	objUUID := request.GetID()

	//intentObj := PhysicalRouterIntent {
	//PhysicalRouter: *obj,
	//}

	//EvaluateDependencies(intentObj, "PhysicalRouter")

	objMap, ok := compilationif.ObjsCache.Load("PhysicalRouterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeletePhysicalRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeletePhysicalRouter(ctx, request)
}

// GetPhysicalRouter handles get request
func (service *PluginService) GetPhysicalRouter(ctx context.Context, request *services.GetPhysicalRouterRequest) (*services.GetPhysicalRouterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("PhysicalRouter")
	if !ok {
		return nil, errors.New("PhysicalRouter get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("PhysicalRouter get failed ")
	}

	response := &services.GetPhysicalRouterResponse{
		PhysicalRouter: obj.(*models.PhysicalRouter),
	}
	return response, nil
}
