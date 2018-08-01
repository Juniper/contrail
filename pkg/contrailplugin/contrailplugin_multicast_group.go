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

// MulticastGroupIntent
//   A struct to store attributes related to MulticastGroup
//   needed by Intent Compiler
type MulticastGroupIntent struct {
	Uuid string
}

// EvaluateMulticastGroup - evaluates the MulticastGroup
func EvaluateMulticastGroup(obj interface{}) {
	resourceObj := obj.(MulticastGroupIntent)
	log.Println("EvaluateMulticastGroup Called ", resourceObj)
}

// CreateMulticastGroup handles create request
func (service *PluginService) CreateMulticastGroup(ctx context.Context, request *services.CreateMulticastGroupRequest) (*services.CreateMulticastGroupResponse, error) {
	log.Println(" CreateMulticastGroup Entered")

	obj := request.GetMulticastGroup()

	intentObj := MulticastGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("MulticastGroupIntent"); !ok {
		compilationif.ObjsCache.Store("MulticastGroupIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("MulticastGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateMulticastGroup", objMap.(*sync.Map))

	EvaluateDependencies(obj, "MulticastGroup")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateMulticastGroup(ctx, request)
}

// UpdateMulticastGroup handles update request
func (service *PluginService) UpdateMulticastGroup(ctx context.Context, request *services.UpdateMulticastGroupRequest) (*services.UpdateMulticastGroupResponse, error) {
	log.Println(" UpdateMulticastGroup ENTERED")

	obj := request.GetMulticastGroup()

	intentObj := MulticastGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("MulticastGroupIntent"); !ok {
		compilationif.ObjsCache.Store("MulticastGroupIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "MulticastGroup")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateMulticastGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateMulticastGroup(ctx, request)
}

// DeleteMulticastGroup handles delete request
func (service *PluginService) DeleteMulticastGroup(ctx context.Context, request *services.DeleteMulticastGroupRequest) (*services.DeleteMulticastGroupResponse, error) {
	log.Println(" DeleteMulticastGroup ENTERED")

	objUUID := request.GetID()

	//intentObj := MulticastGroupIntent {
	//MulticastGroup: *obj,
	//}

	//EvaluateDependencies(intentObj, "MulticastGroup")

	objMap, ok := compilationif.ObjsCache.Load("MulticastGroupIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteMulticastGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteMulticastGroup(ctx, request)
}

// GetMulticastGroup handles get request
func (service *PluginService) GetMulticastGroup(ctx context.Context, request *services.GetMulticastGroupRequest) (*services.GetMulticastGroupResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("MulticastGroup")
	if !ok {
		return nil, errors.New("MulticastGroup get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("MulticastGroup get failed ")
	}

	response := &services.GetMulticastGroupResponse{
		MulticastGroup: obj.(*models.MulticastGroup),
	}
	return response, nil
}
