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

// RouteAggregateIntent
//   A struct to store attributes related to RouteAggregate
//   needed by Intent Compiler
type RouteAggregateIntent struct {
	Uuid string
}

// EvaluateRouteAggregate - evaluates the RouteAggregate
func EvaluateRouteAggregate(obj interface{}) {
	resourceObj := obj.(RouteAggregateIntent)
	log.Println("EvaluateRouteAggregate Called ", resourceObj)
}

// CreateRouteAggregate handles create request
func (service *PluginService) CreateRouteAggregate(ctx context.Context, request *services.CreateRouteAggregateRequest) (*services.CreateRouteAggregateResponse, error) {
	log.Println(" CreateRouteAggregate Entered")

	obj := request.GetRouteAggregate()

	intentObj := RouteAggregateIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RouteAggregateIntent"); !ok {
		compilationif.ObjsCache.Store("RouteAggregateIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("RouteAggregateIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateRouteAggregate", objMap.(*sync.Map))

	EvaluateDependencies(obj, "RouteAggregate")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateRouteAggregate(ctx, request)
}

// UpdateRouteAggregate handles update request
func (service *PluginService) UpdateRouteAggregate(ctx context.Context, request *services.UpdateRouteAggregateRequest) (*services.UpdateRouteAggregateResponse, error) {
	log.Println(" UpdateRouteAggregate ENTERED")

	obj := request.GetRouteAggregate()

	intentObj := RouteAggregateIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RouteAggregateIntent"); !ok {
		compilationif.ObjsCache.Store("RouteAggregateIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "RouteAggregate")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateRouteAggregate", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteAggregate(ctx, request)
}

// DeleteRouteAggregate handles delete request
func (service *PluginService) DeleteRouteAggregate(ctx context.Context, request *services.DeleteRouteAggregateRequest) (*services.DeleteRouteAggregateResponse, error) {
	log.Println(" DeleteRouteAggregate ENTERED")

	objUUID := request.GetID()

	//intentObj := RouteAggregateIntent {
	//RouteAggregate: *obj,
	//}

	//EvaluateDependencies(intentObj, "RouteAggregate")

	objMap, ok := compilationif.ObjsCache.Load("RouteAggregateIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteRouteAggregate", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteAggregate(ctx, request)
}

// GetRouteAggregate handles get request
func (service *PluginService) GetRouteAggregate(ctx context.Context, request *services.GetRouteAggregateRequest) (*services.GetRouteAggregateResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("RouteAggregate")
	if !ok {
		return nil, errors.New("RouteAggregate get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("RouteAggregate get failed ")
	}

	response := &services.GetRouteAggregateResponse{
		RouteAggregate: obj.(*models.RouteAggregate),
	}
	return response, nil
}
