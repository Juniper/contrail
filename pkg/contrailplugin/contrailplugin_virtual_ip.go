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

// VirtualIPIntent
//   A struct to store attributes related to VirtualIP
//   needed by Intent Compiler
type VirtualIPIntent struct {
	Uuid string
}

// EvaluateVirtualIP - evaluates the VirtualIP
func EvaluateVirtualIP(obj interface{}) {
	resourceObj := obj.(VirtualIPIntent)
	log.Println("EvaluateVirtualIP Called ", resourceObj)
}

// CreateVirtualIP handles create request
func (service *PluginService) CreateVirtualIP(ctx context.Context, request *services.CreateVirtualIPRequest) (*services.CreateVirtualIPResponse, error) {
	log.Println(" CreateVirtualIP Entered")

	obj := request.GetVirtualIP()

	intentObj := VirtualIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualIPIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualIPIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VirtualIPIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVirtualIP", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VirtualIP")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualIP(ctx, request)
}

// UpdateVirtualIP handles update request
func (service *PluginService) UpdateVirtualIP(ctx context.Context, request *services.UpdateVirtualIPRequest) (*services.UpdateVirtualIPResponse, error) {
	log.Println(" UpdateVirtualIP ENTERED")

	obj := request.GetVirtualIP()

	intentObj := VirtualIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualIPIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualIPIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VirtualIP")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVirtualIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualIP(ctx, request)
}

// DeleteVirtualIP handles delete request
func (service *PluginService) DeleteVirtualIP(ctx context.Context, request *services.DeleteVirtualIPRequest) (*services.DeleteVirtualIPResponse, error) {
	log.Println(" DeleteVirtualIP ENTERED")

	objUUID := request.GetID()

	//intentObj := VirtualIPIntent {
	//VirtualIP: *obj,
	//}

	//EvaluateDependencies(intentObj, "VirtualIP")

	objMap, ok := compilationif.ObjsCache.Load("VirtualIPIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVirtualIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualIP(ctx, request)
}

// GetVirtualIP handles get request
func (service *PluginService) GetVirtualIP(ctx context.Context, request *services.GetVirtualIPRequest) (*services.GetVirtualIPResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VirtualIP")
	if !ok {
		return nil, errors.New("VirtualIP get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VirtualIP get failed ")
	}

	response := &services.GetVirtualIPResponse{
		VirtualIP: obj.(*models.VirtualIP),
	}
	return response, nil
}
