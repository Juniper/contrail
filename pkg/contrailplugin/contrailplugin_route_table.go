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

// RouteTableIntent
//   A struct to store attributes related to RouteTable
//   needed by Intent Compiler
type RouteTableIntent struct {
	Uuid string
}

// EvaluateRouteTable - evaluates the RouteTable
func EvaluateRouteTable(obj interface{}) {
	resourceObj := obj.(RouteTableIntent)
	log.Println("EvaluateRouteTable Called ", resourceObj)
}

// CreateRouteTable handles create request
func (service *PluginService) CreateRouteTable(ctx context.Context, request *services.CreateRouteTableRequest) (*services.CreateRouteTableResponse, error) {
	log.Println(" CreateRouteTable Entered")

	obj := request.GetRouteTable()

	intentObj := RouteTableIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RouteTableIntent"); !ok {
		compilationif.ObjsCache.Store("RouteTableIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("RouteTableIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateRouteTable", objMap.(*sync.Map))

	EvaluateDependencies(obj, "RouteTable")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateRouteTable(ctx, request)
}

// UpdateRouteTable handles update request
func (service *PluginService) UpdateRouteTable(ctx context.Context, request *services.UpdateRouteTableRequest) (*services.UpdateRouteTableResponse, error) {
	log.Println(" UpdateRouteTable ENTERED")

	obj := request.GetRouteTable()

	intentObj := RouteTableIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RouteTableIntent"); !ok {
		compilationif.ObjsCache.Store("RouteTableIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "RouteTable")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateRouteTable", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateRouteTable(ctx, request)
}

// DeleteRouteTable handles delete request
func (service *PluginService) DeleteRouteTable(ctx context.Context, request *services.DeleteRouteTableRequest) (*services.DeleteRouteTableResponse, error) {
	log.Println(" DeleteRouteTable ENTERED")

	objUUID := request.GetID()

	//intentObj := RouteTableIntent {
	//RouteTable: *obj,
	//}

	//EvaluateDependencies(intentObj, "RouteTable")

	objMap, ok := compilationif.ObjsCache.Load("RouteTableIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteRouteTable", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteRouteTable(ctx, request)
}

// GetRouteTable handles get request
func (service *PluginService) GetRouteTable(ctx context.Context, request *services.GetRouteTableRequest) (*services.GetRouteTableResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("RouteTable")
	if !ok {
		return nil, errors.New("RouteTable get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("RouteTable get failed ")
	}

	response := &services.GetRouteTableResponse{
		RouteTable: obj.(*models.RouteTable),
	}
	return response, nil
}
