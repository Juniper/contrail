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

// VirtualMachineInterfaceIntent
//   A struct to store attributes related to VirtualMachineInterface
//   needed by Intent Compiler
type VirtualMachineInterfaceIntent struct {
	Uuid string
}

// EvaluateVirtualMachineInterface - evaluates the VirtualMachineInterface
func EvaluateVirtualMachineInterface(obj interface{}) {
	resourceObj := obj.(VirtualMachineInterfaceIntent)
	log.Println("EvaluateVirtualMachineInterface Called ", resourceObj)
}

// CreateVirtualMachineInterface handles create request
func (service *PluginService) CreateVirtualMachineInterface(ctx context.Context, request *services.CreateVirtualMachineInterfaceRequest) (*services.CreateVirtualMachineInterfaceResponse, error) {
	log.Println(" CreateVirtualMachineInterface Entered")

	obj := request.GetVirtualMachineInterface()

	intentObj := VirtualMachineInterfaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualMachineInterfaceIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualMachineInterfaceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VirtualMachineInterfaceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVirtualMachineInterface", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VirtualMachineInterface")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualMachineInterface(ctx, request)
}

// UpdateVirtualMachineInterface handles update request
func (service *PluginService) UpdateVirtualMachineInterface(ctx context.Context, request *services.UpdateVirtualMachineInterfaceRequest) (*services.UpdateVirtualMachineInterfaceResponse, error) {
	log.Println(" UpdateVirtualMachineInterface ENTERED")

	obj := request.GetVirtualMachineInterface()

	intentObj := VirtualMachineInterfaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualMachineInterfaceIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualMachineInterfaceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VirtualMachineInterface")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVirtualMachineInterface", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualMachineInterface(ctx, request)
}

// DeleteVirtualMachineInterface handles delete request
func (service *PluginService) DeleteVirtualMachineInterface(ctx context.Context, request *services.DeleteVirtualMachineInterfaceRequest) (*services.DeleteVirtualMachineInterfaceResponse, error) {
	log.Println(" DeleteVirtualMachineInterface ENTERED")

	objUUID := request.GetID()

	//intentObj := VirtualMachineInterfaceIntent {
	//VirtualMachineInterface: *obj,
	//}

	//EvaluateDependencies(intentObj, "VirtualMachineInterface")

	objMap, ok := compilationif.ObjsCache.Load("VirtualMachineInterfaceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVirtualMachineInterface", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualMachineInterface(ctx, request)
}

// GetVirtualMachineInterface handles get request
func (service *PluginService) GetVirtualMachineInterface(ctx context.Context, request *services.GetVirtualMachineInterfaceRequest) (*services.GetVirtualMachineInterfaceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VirtualMachineInterface")
	if !ok {
		return nil, errors.New("VirtualMachineInterface get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VirtualMachineInterface get failed ")
	}

	response := &services.GetVirtualMachineInterfaceResponse{
		VirtualMachineInterface: obj.(*models.VirtualMachineInterface),
	}
	return response, nil
}
