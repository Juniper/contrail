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

// VirtualDNSIntent
//   A struct to store attributes related to VirtualDNS
//   needed by Intent Compiler
type VirtualDNSIntent struct {
	Uuid string
}

// EvaluateVirtualDNS - evaluates the VirtualDNS
func EvaluateVirtualDNS(obj interface{}) {
	resourceObj := obj.(VirtualDNSIntent)
	log.Println("EvaluateVirtualDNS Called ", resourceObj)
}

// CreateVirtualDNS handles create request
func (service *PluginService) CreateVirtualDNS(ctx context.Context, request *services.CreateVirtualDNSRequest) (*services.CreateVirtualDNSResponse, error) {
	log.Println(" CreateVirtualDNS Entered")

	obj := request.GetVirtualDNS()

	intentObj := VirtualDNSIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualDNSIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualDNSIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VirtualDNSIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVirtualDNS", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VirtualDNS")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualDNS(ctx, request)
}

// UpdateVirtualDNS handles update request
func (service *PluginService) UpdateVirtualDNS(ctx context.Context, request *services.UpdateVirtualDNSRequest) (*services.UpdateVirtualDNSResponse, error) {
	log.Println(" UpdateVirtualDNS ENTERED")

	obj := request.GetVirtualDNS()

	intentObj := VirtualDNSIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualDNSIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualDNSIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VirtualDNS")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVirtualDNS", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualDNS(ctx, request)
}

// DeleteVirtualDNS handles delete request
func (service *PluginService) DeleteVirtualDNS(ctx context.Context, request *services.DeleteVirtualDNSRequest) (*services.DeleteVirtualDNSResponse, error) {
	log.Println(" DeleteVirtualDNS ENTERED")

	objUUID := request.GetID()

	//intentObj := VirtualDNSIntent {
	//VirtualDNS: *obj,
	//}

	//EvaluateDependencies(intentObj, "VirtualDNS")

	objMap, ok := compilationif.ObjsCache.Load("VirtualDNSIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVirtualDNS", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualDNS(ctx, request)
}

// GetVirtualDNS handles get request
func (service *PluginService) GetVirtualDNS(ctx context.Context, request *services.GetVirtualDNSRequest) (*services.GetVirtualDNSResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VirtualDNS")
	if !ok {
		return nil, errors.New("VirtualDNS get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VirtualDNS get failed ")
	}

	response := &services.GetVirtualDNSResponse{
		VirtualDNS: obj.(*models.VirtualDNS),
	}
	return response, nil
}
