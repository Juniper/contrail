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

// VirtualNetworkIntent
//   A struct to store attributes related to VirtualNetwork
//   needed by Intent Compiler
type VirtualNetworkIntent struct {
	Uuid string
}

// EvaluateVirtualNetwork - evaluates the VirtualNetwork
func EvaluateVirtualNetwork(obj interface{}) {
	resourceObj := obj.(VirtualNetworkIntent)
	log.Println("EvaluateVirtualNetwork Called ", resourceObj)
}

// CreateVirtualNetwork handles create request
func (service *PluginService) CreateVirtualNetwork(ctx context.Context, request *services.CreateVirtualNetworkRequest) (*services.CreateVirtualNetworkResponse, error) {
	log.Println(" CreateVirtualNetwork Entered")

	obj := request.GetVirtualNetwork()

	intentObj := VirtualNetworkIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualNetworkIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualNetworkIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VirtualNetworkIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVirtualNetwork", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VirtualNetwork")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualNetwork(ctx, request)
}

// UpdateVirtualNetwork handles update request
func (service *PluginService) UpdateVirtualNetwork(ctx context.Context, request *services.UpdateVirtualNetworkRequest) (*services.UpdateVirtualNetworkResponse, error) {
	log.Println(" UpdateVirtualNetwork ENTERED")

	obj := request.GetVirtualNetwork()

	intentObj := VirtualNetworkIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualNetworkIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualNetworkIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VirtualNetwork")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVirtualNetwork", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualNetwork(ctx, request)
}

// DeleteVirtualNetwork handles delete request
func (service *PluginService) DeleteVirtualNetwork(ctx context.Context, request *services.DeleteVirtualNetworkRequest) (*services.DeleteVirtualNetworkResponse, error) {
	log.Println(" DeleteVirtualNetwork ENTERED")

	objUUID := request.GetID()

	//intentObj := VirtualNetworkIntent {
	//VirtualNetwork: *obj,
	//}

	//EvaluateDependencies(intentObj, "VirtualNetwork")

	objMap, ok := compilationif.ObjsCache.Load("VirtualNetworkIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVirtualNetwork", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualNetwork(ctx, request)
}

// GetVirtualNetwork handles get request
func (service *PluginService) GetVirtualNetwork(ctx context.Context, request *services.GetVirtualNetworkRequest) (*services.GetVirtualNetworkResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VirtualNetwork")
	if !ok {
		return nil, errors.New("VirtualNetwork get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VirtualNetwork get failed ")
	}

	response := &services.GetVirtualNetworkResponse{
		VirtualNetwork: obj.(*models.VirtualNetwork),
	}
	return response, nil
}
