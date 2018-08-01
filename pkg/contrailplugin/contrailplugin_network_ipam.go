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

// NetworkIpamIntent
//   A struct to store attributes related to NetworkIpam
//   needed by Intent Compiler
type NetworkIpamIntent struct {
	Uuid string
}

// EvaluateNetworkIpam - evaluates the NetworkIpam
func EvaluateNetworkIpam(obj interface{}) {
	resourceObj := obj.(NetworkIpamIntent)
	log.Println("EvaluateNetworkIpam Called ", resourceObj)
}

// CreateNetworkIpam handles create request
func (service *PluginService) CreateNetworkIpam(ctx context.Context, request *services.CreateNetworkIpamRequest) (*services.CreateNetworkIpamResponse, error) {
	log.Println(" CreateNetworkIpam Entered")

	obj := request.GetNetworkIpam()

	intentObj := NetworkIpamIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NetworkIpamIntent"); !ok {
		compilationif.ObjsCache.Store("NetworkIpamIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("NetworkIpamIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateNetworkIpam", objMap.(*sync.Map))

	EvaluateDependencies(obj, "NetworkIpam")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkIpam(ctx, request)
}

// UpdateNetworkIpam handles update request
func (service *PluginService) UpdateNetworkIpam(ctx context.Context, request *services.UpdateNetworkIpamRequest) (*services.UpdateNetworkIpamResponse, error) {
	log.Println(" UpdateNetworkIpam ENTERED")

	obj := request.GetNetworkIpam()

	intentObj := NetworkIpamIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NetworkIpamIntent"); !ok {
		compilationif.ObjsCache.Store("NetworkIpamIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "NetworkIpam")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateNetworkIpam", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkIpam(ctx, request)
}

// DeleteNetworkIpam handles delete request
func (service *PluginService) DeleteNetworkIpam(ctx context.Context, request *services.DeleteNetworkIpamRequest) (*services.DeleteNetworkIpamResponse, error) {
	log.Println(" DeleteNetworkIpam ENTERED")

	objUUID := request.GetID()

	//intentObj := NetworkIpamIntent {
	//NetworkIpam: *obj,
	//}

	//EvaluateDependencies(intentObj, "NetworkIpam")

	objMap, ok := compilationif.ObjsCache.Load("NetworkIpamIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteNetworkIpam", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkIpam(ctx, request)
}

// GetNetworkIpam handles get request
func (service *PluginService) GetNetworkIpam(ctx context.Context, request *services.GetNetworkIpamRequest) (*services.GetNetworkIpamResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("NetworkIpam")
	if !ok {
		return nil, errors.New("NetworkIpam get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("NetworkIpam get failed ")
	}

	response := &services.GetNetworkIpamResponse{
		NetworkIpam: obj.(*models.NetworkIpam),
	}
	return response, nil
}
