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

// PortGroupIntent
//   A struct to store attributes related to PortGroup
//   needed by Intent Compiler
type PortGroupIntent struct {
	Uuid string
}

// EvaluatePortGroup - evaluates the PortGroup
func EvaluatePortGroup(obj interface{}) {
	resourceObj := obj.(PortGroupIntent)
	log.Println("EvaluatePortGroup Called ", resourceObj)
}

// CreatePortGroup handles create request
func (service *PluginService) CreatePortGroup(ctx context.Context, request *services.CreatePortGroupRequest) (*services.CreatePortGroupResponse, error) {
	log.Println(" CreatePortGroup Entered")

	obj := request.GetPortGroup()

	intentObj := PortGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PortGroupIntent"); !ok {
		compilationif.ObjsCache.Store("PortGroupIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("PortGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreatePortGroup", objMap.(*sync.Map))

	EvaluateDependencies(obj, "PortGroup")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreatePortGroup(ctx, request)
}

// UpdatePortGroup handles update request
func (service *PluginService) UpdatePortGroup(ctx context.Context, request *services.UpdatePortGroupRequest) (*services.UpdatePortGroupResponse, error) {
	log.Println(" UpdatePortGroup ENTERED")

	obj := request.GetPortGroup()

	intentObj := PortGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PortGroupIntent"); !ok {
		compilationif.ObjsCache.Store("PortGroupIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "PortGroup")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdatePortGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdatePortGroup(ctx, request)
}

// DeletePortGroup handles delete request
func (service *PluginService) DeletePortGroup(ctx context.Context, request *services.DeletePortGroupRequest) (*services.DeletePortGroupResponse, error) {
	log.Println(" DeletePortGroup ENTERED")

	objUUID := request.GetID()

	//intentObj := PortGroupIntent {
	//PortGroup: *obj,
	//}

	//EvaluateDependencies(intentObj, "PortGroup")

	objMap, ok := compilationif.ObjsCache.Load("PortGroupIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeletePortGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeletePortGroup(ctx, request)
}

// GetPortGroup handles get request
func (service *PluginService) GetPortGroup(ctx context.Context, request *services.GetPortGroupRequest) (*services.GetPortGroupResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("PortGroup")
	if !ok {
		return nil, errors.New("PortGroup get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("PortGroup get failed ")
	}

	response := &services.GetPortGroupResponse{
		PortGroup: obj.(*models.PortGroup),
	}
	return response, nil
}
