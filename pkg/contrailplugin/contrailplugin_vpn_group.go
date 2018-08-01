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

// VPNGroupIntent
//   A struct to store attributes related to VPNGroup
//   needed by Intent Compiler
type VPNGroupIntent struct {
	Uuid string
}

// EvaluateVPNGroup - evaluates the VPNGroup
func EvaluateVPNGroup(obj interface{}) {
	resourceObj := obj.(VPNGroupIntent)
	log.Println("EvaluateVPNGroup Called ", resourceObj)
}

// CreateVPNGroup handles create request
func (service *PluginService) CreateVPNGroup(ctx context.Context, request *services.CreateVPNGroupRequest) (*services.CreateVPNGroupResponse, error) {
	log.Println(" CreateVPNGroup Entered")

	obj := request.GetVPNGroup()

	intentObj := VPNGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VPNGroupIntent"); !ok {
		compilationif.ObjsCache.Store("VPNGroupIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VPNGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVPNGroup", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VPNGroup")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVPNGroup(ctx, request)
}

// UpdateVPNGroup handles update request
func (service *PluginService) UpdateVPNGroup(ctx context.Context, request *services.UpdateVPNGroupRequest) (*services.UpdateVPNGroupResponse, error) {
	log.Println(" UpdateVPNGroup ENTERED")

	obj := request.GetVPNGroup()

	intentObj := VPNGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VPNGroupIntent"); !ok {
		compilationif.ObjsCache.Store("VPNGroupIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VPNGroup")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVPNGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVPNGroup(ctx, request)
}

// DeleteVPNGroup handles delete request
func (service *PluginService) DeleteVPNGroup(ctx context.Context, request *services.DeleteVPNGroupRequest) (*services.DeleteVPNGroupResponse, error) {
	log.Println(" DeleteVPNGroup ENTERED")

	objUUID := request.GetID()

	//intentObj := VPNGroupIntent {
	//VPNGroup: *obj,
	//}

	//EvaluateDependencies(intentObj, "VPNGroup")

	objMap, ok := compilationif.ObjsCache.Load("VPNGroupIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVPNGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVPNGroup(ctx, request)
}

// GetVPNGroup handles get request
func (service *PluginService) GetVPNGroup(ctx context.Context, request *services.GetVPNGroupRequest) (*services.GetVPNGroupResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VPNGroup")
	if !ok {
		return nil, errors.New("VPNGroup get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VPNGroup get failed ")
	}

	response := &services.GetVPNGroupResponse{
		VPNGroup: obj.(*models.VPNGroup),
	}
	return response, nil
}
