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

// VirtualMachineIntent
//   A struct to store attributes related to VirtualMachine
//   needed by Intent Compiler
type VirtualMachineIntent struct {
	Uuid string
}

// EvaluateVirtualMachine - evaluates the VirtualMachine
func EvaluateVirtualMachine(obj interface{}) {
	resourceObj := obj.(VirtualMachineIntent)
	log.Println("EvaluateVirtualMachine Called ", resourceObj)
}

// CreateVirtualMachine handles create request
func (service *PluginService) CreateVirtualMachine(ctx context.Context, request *services.CreateVirtualMachineRequest) (*services.CreateVirtualMachineResponse, error) {
	log.Println(" CreateVirtualMachine Entered")

	obj := request.GetVirtualMachine()

	intentObj := VirtualMachineIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualMachineIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualMachineIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VirtualMachineIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVirtualMachine", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VirtualMachine")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualMachine(ctx, request)
}

// UpdateVirtualMachine handles update request
func (service *PluginService) UpdateVirtualMachine(ctx context.Context, request *services.UpdateVirtualMachineRequest) (*services.UpdateVirtualMachineResponse, error) {
	log.Println(" UpdateVirtualMachine ENTERED")

	obj := request.GetVirtualMachine()

	intentObj := VirtualMachineIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualMachineIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualMachineIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VirtualMachine")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVirtualMachine", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualMachine(ctx, request)
}

// DeleteVirtualMachine handles delete request
func (service *PluginService) DeleteVirtualMachine(ctx context.Context, request *services.DeleteVirtualMachineRequest) (*services.DeleteVirtualMachineResponse, error) {
	log.Println(" DeleteVirtualMachine ENTERED")

	objUUID := request.GetID()

	//intentObj := VirtualMachineIntent {
	//VirtualMachine: *obj,
	//}

	//EvaluateDependencies(intentObj, "VirtualMachine")

	objMap, ok := compilationif.ObjsCache.Load("VirtualMachineIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVirtualMachine", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualMachine(ctx, request)
}

// GetVirtualMachine handles get request
func (service *PluginService) GetVirtualMachine(ctx context.Context, request *services.GetVirtualMachineRequest) (*services.GetVirtualMachineResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VirtualMachine")
	if !ok {
		return nil, errors.New("VirtualMachine get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VirtualMachine get failed ")
	}

	response := &services.GetVirtualMachineResponse{
		VirtualMachine: obj.(*models.VirtualMachine),
	}
	return response, nil
}
