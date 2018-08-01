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

// NetworkDeviceConfigIntent
//   A struct to store attributes related to NetworkDeviceConfig
//   needed by Intent Compiler
type NetworkDeviceConfigIntent struct {
	Uuid string
}

// EvaluateNetworkDeviceConfig - evaluates the NetworkDeviceConfig
func EvaluateNetworkDeviceConfig(obj interface{}) {
	resourceObj := obj.(NetworkDeviceConfigIntent)
	log.Println("EvaluateNetworkDeviceConfig Called ", resourceObj)
}

// CreateNetworkDeviceConfig handles create request
func (service *PluginService) CreateNetworkDeviceConfig(ctx context.Context, request *services.CreateNetworkDeviceConfigRequest) (*services.CreateNetworkDeviceConfigResponse, error) {
	log.Println(" CreateNetworkDeviceConfig Entered")

	obj := request.GetNetworkDeviceConfig()

	intentObj := NetworkDeviceConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NetworkDeviceConfigIntent"); !ok {
		compilationif.ObjsCache.Store("NetworkDeviceConfigIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("NetworkDeviceConfigIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateNetworkDeviceConfig", objMap.(*sync.Map))

	EvaluateDependencies(obj, "NetworkDeviceConfig")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkDeviceConfig(ctx, request)
}

// UpdateNetworkDeviceConfig handles update request
func (service *PluginService) UpdateNetworkDeviceConfig(ctx context.Context, request *services.UpdateNetworkDeviceConfigRequest) (*services.UpdateNetworkDeviceConfigResponse, error) {
	log.Println(" UpdateNetworkDeviceConfig ENTERED")

	obj := request.GetNetworkDeviceConfig()

	intentObj := NetworkDeviceConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NetworkDeviceConfigIntent"); !ok {
		compilationif.ObjsCache.Store("NetworkDeviceConfigIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "NetworkDeviceConfig")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateNetworkDeviceConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkDeviceConfig(ctx, request)
}

// DeleteNetworkDeviceConfig handles delete request
func (service *PluginService) DeleteNetworkDeviceConfig(ctx context.Context, request *services.DeleteNetworkDeviceConfigRequest) (*services.DeleteNetworkDeviceConfigResponse, error) {
	log.Println(" DeleteNetworkDeviceConfig ENTERED")

	objUUID := request.GetID()

	//intentObj := NetworkDeviceConfigIntent {
	//NetworkDeviceConfig: *obj,
	//}

	//EvaluateDependencies(intentObj, "NetworkDeviceConfig")

	objMap, ok := compilationif.ObjsCache.Load("NetworkDeviceConfigIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteNetworkDeviceConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkDeviceConfig(ctx, request)
}

// GetNetworkDeviceConfig handles get request
func (service *PluginService) GetNetworkDeviceConfig(ctx context.Context, request *services.GetNetworkDeviceConfigRequest) (*services.GetNetworkDeviceConfigResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("NetworkDeviceConfig")
	if !ok {
		return nil, errors.New("NetworkDeviceConfig get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("NetworkDeviceConfig get failed ")
	}

	response := &services.GetNetworkDeviceConfigResponse{
		NetworkDeviceConfig: obj.(*models.NetworkDeviceConfig),
	}
	return response, nil
}
