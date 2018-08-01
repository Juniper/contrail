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

// VirtualRouterIntent
//   A struct to store attributes related to VirtualRouter
//   needed by Intent Compiler
type VirtualRouterIntent struct {
	Uuid string
}

// EvaluateVirtualRouter - evaluates the VirtualRouter
func EvaluateVirtualRouter(obj interface{}) {
	resourceObj := obj.(VirtualRouterIntent)
	log.Println("EvaluateVirtualRouter Called ", resourceObj)
}

// CreateVirtualRouter handles create request
func (service *PluginService) CreateVirtualRouter(ctx context.Context, request *services.CreateVirtualRouterRequest) (*services.CreateVirtualRouterResponse, error) {
	log.Println(" CreateVirtualRouter Entered")

	obj := request.GetVirtualRouter()

	intentObj := VirtualRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualRouterIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualRouterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VirtualRouterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVirtualRouter", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VirtualRouter")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualRouter(ctx, request)
}

// UpdateVirtualRouter handles update request
func (service *PluginService) UpdateVirtualRouter(ctx context.Context, request *services.UpdateVirtualRouterRequest) (*services.UpdateVirtualRouterResponse, error) {
	log.Println(" UpdateVirtualRouter ENTERED")

	obj := request.GetVirtualRouter()

	intentObj := VirtualRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualRouterIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualRouterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VirtualRouter")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVirtualRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualRouter(ctx, request)
}

// DeleteVirtualRouter handles delete request
func (service *PluginService) DeleteVirtualRouter(ctx context.Context, request *services.DeleteVirtualRouterRequest) (*services.DeleteVirtualRouterResponse, error) {
	log.Println(" DeleteVirtualRouter ENTERED")

	objUUID := request.GetID()

	//intentObj := VirtualRouterIntent {
	//VirtualRouter: *obj,
	//}

	//EvaluateDependencies(intentObj, "VirtualRouter")

	objMap, ok := compilationif.ObjsCache.Load("VirtualRouterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVirtualRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualRouter(ctx, request)
}

// GetVirtualRouter handles get request
func (service *PluginService) GetVirtualRouter(ctx context.Context, request *services.GetVirtualRouterRequest) (*services.GetVirtualRouterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VirtualRouter")
	if !ok {
		return nil, errors.New("VirtualRouter get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VirtualRouter get failed ")
	}

	response := &services.GetVirtualRouterResponse{
		VirtualRouter: obj.(*models.VirtualRouter),
	}
	return response, nil
}
