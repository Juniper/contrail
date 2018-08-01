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

// LogicalRouterIntent
//   A struct to store attributes related to LogicalRouter
//   needed by Intent Compiler
type LogicalRouterIntent struct {
	Uuid string
}

// EvaluateLogicalRouter - evaluates the LogicalRouter
func EvaluateLogicalRouter(obj interface{}) {
	resourceObj := obj.(LogicalRouterIntent)
	log.Println("EvaluateLogicalRouter Called ", resourceObj)
}

// CreateLogicalRouter handles create request
func (service *PluginService) CreateLogicalRouter(ctx context.Context, request *services.CreateLogicalRouterRequest) (*services.CreateLogicalRouterResponse, error) {
	log.Println(" CreateLogicalRouter Entered")

	obj := request.GetLogicalRouter()

	intentObj := LogicalRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LogicalRouterIntent"); !ok {
		compilationif.ObjsCache.Store("LogicalRouterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LogicalRouterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLogicalRouter", objMap.(*sync.Map))

	EvaluateDependencies(obj, "LogicalRouter")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLogicalRouter(ctx, request)
}

// UpdateLogicalRouter handles update request
func (service *PluginService) UpdateLogicalRouter(ctx context.Context, request *services.UpdateLogicalRouterRequest) (*services.UpdateLogicalRouterResponse, error) {
	log.Println(" UpdateLogicalRouter ENTERED")

	obj := request.GetLogicalRouter()

	intentObj := LogicalRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LogicalRouterIntent"); !ok {
		compilationif.ObjsCache.Store("LogicalRouterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "LogicalRouter")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLogicalRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLogicalRouter(ctx, request)
}

// DeleteLogicalRouter handles delete request
func (service *PluginService) DeleteLogicalRouter(ctx context.Context, request *services.DeleteLogicalRouterRequest) (*services.DeleteLogicalRouterResponse, error) {
	log.Println(" DeleteLogicalRouter ENTERED")

	objUUID := request.GetID()

	//intentObj := LogicalRouterIntent {
	//LogicalRouter: *obj,
	//}

	//EvaluateDependencies(intentObj, "LogicalRouter")

	objMap, ok := compilationif.ObjsCache.Load("LogicalRouterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLogicalRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLogicalRouter(ctx, request)
}

// GetLogicalRouter handles get request
func (service *PluginService) GetLogicalRouter(ctx context.Context, request *services.GetLogicalRouterRequest) (*services.GetLogicalRouterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("LogicalRouter")
	if !ok {
		return nil, errors.New("LogicalRouter get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("LogicalRouter get failed ")
	}

	response := &services.GetLogicalRouterResponse{
		LogicalRouter: obj.(*models.LogicalRouter),
	}
	return response, nil
}
