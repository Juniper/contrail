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

// RouteTargetIntent
//   A struct to store attributes related to RouteTarget
//   needed by Intent Compiler
type RouteTargetIntent struct {
	Uuid string
}

// EvaluateRouteTarget - evaluates the RouteTarget
func EvaluateRouteTarget(obj interface{}) {
	resourceObj := obj.(RouteTargetIntent)
	log.Println("EvaluateRouteTarget Called ", resourceObj)
}

// CreateRouteTarget handles create request
func (service *PluginService) CreateRouteTarget(ctx context.Context, request *services.CreateRouteTargetRequest) (*services.CreateRouteTargetResponse, error) {
	log.Println(" CreateRouteTarget Entered")

	obj := request.GetRouteTarget()

	intentObj := RouteTargetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RouteTargetIntent"); !ok {
		compilationif.ObjsCache.Store("RouteTargetIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("RouteTargetIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateRouteTarget", objMap.(*sync.Map))

	EvaluateDependencies(obj, "RouteTarget")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateRouteTarget(ctx, request)
}

// UpdateRouteTarget handles update request
func (service *PluginService) UpdateRouteTarget(ctx context.Context, request *services.UpdateRouteTargetRequest) (*services.UpdateRouteTargetResponse, error) {
	log.Println(" UpdateRouteTarget ENTERED")

	obj := request.GetRouteTarget()

	intentObj := RouteTargetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RouteTargetIntent"); !ok {
		compilationif.ObjsCache.Store("RouteTargetIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "RouteTarget")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateRouteTarget", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteTarget(ctx, request)
}

// DeleteRouteTarget handles delete request
func (service *PluginService) DeleteRouteTarget(ctx context.Context, request *services.DeleteRouteTargetRequest) (*services.DeleteRouteTargetResponse, error) {
	log.Println(" DeleteRouteTarget ENTERED")

	objUUID := request.GetID()

	//intentObj := RouteTargetIntent {
	//RouteTarget: *obj,
	//}

	//EvaluateDependencies(intentObj, "RouteTarget")

	objMap, ok := compilationif.ObjsCache.Load("RouteTargetIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteRouteTarget", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteTarget(ctx, request)
}

// GetRouteTarget handles get request
func (service *PluginService) GetRouteTarget(ctx context.Context, request *services.GetRouteTargetRequest) (*services.GetRouteTargetResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("RouteTarget")
	if !ok {
		return nil, errors.New("RouteTarget get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("RouteTarget get failed ")
	}

	response := &services.GetRouteTargetResponse{
		RouteTarget: obj.(*models.RouteTarget),
	}
	return response, nil
}
