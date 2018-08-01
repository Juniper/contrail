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

// GlobalQosConfigIntent
//   A struct to store attributes related to GlobalQosConfig
//   needed by Intent Compiler
type GlobalQosConfigIntent struct {
	Uuid string
}

// EvaluateGlobalQosConfig - evaluates the GlobalQosConfig
func EvaluateGlobalQosConfig(obj interface{}) {
	resourceObj := obj.(GlobalQosConfigIntent)
	log.Println("EvaluateGlobalQosConfig Called ", resourceObj)
}

// CreateGlobalQosConfig handles create request
func (service *PluginService) CreateGlobalQosConfig(ctx context.Context, request *services.CreateGlobalQosConfigRequest) (*services.CreateGlobalQosConfigResponse, error) {
	log.Println(" CreateGlobalQosConfig Entered")

	obj := request.GetGlobalQosConfig()

	intentObj := GlobalQosConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalQosConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalQosConfigIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("GlobalQosConfigIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateGlobalQosConfig", objMap.(*sync.Map))

	EvaluateDependencies(obj, "GlobalQosConfig")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalQosConfig(ctx, request)
}

// UpdateGlobalQosConfig handles update request
func (service *PluginService) UpdateGlobalQosConfig(ctx context.Context, request *services.UpdateGlobalQosConfigRequest) (*services.UpdateGlobalQosConfigResponse, error) {
	log.Println(" UpdateGlobalQosConfig ENTERED")

	obj := request.GetGlobalQosConfig()

	intentObj := GlobalQosConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalQosConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalQosConfigIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "GlobalQosConfig")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateGlobalQosConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalQosConfig(ctx, request)
}

// DeleteGlobalQosConfig handles delete request
func (service *PluginService) DeleteGlobalQosConfig(ctx context.Context, request *services.DeleteGlobalQosConfigRequest) (*services.DeleteGlobalQosConfigResponse, error) {
	log.Println(" DeleteGlobalQosConfig ENTERED")

	objUUID := request.GetID()

	//intentObj := GlobalQosConfigIntent {
	//GlobalQosConfig: *obj,
	//}

	//EvaluateDependencies(intentObj, "GlobalQosConfig")

	objMap, ok := compilationif.ObjsCache.Load("GlobalQosConfigIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteGlobalQosConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalQosConfig(ctx, request)
}

// GetGlobalQosConfig handles get request
func (service *PluginService) GetGlobalQosConfig(ctx context.Context, request *services.GetGlobalQosConfigRequest) (*services.GetGlobalQosConfigResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("GlobalQosConfig")
	if !ok {
		return nil, errors.New("GlobalQosConfig get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("GlobalQosConfig get failed ")
	}

	response := &services.GetGlobalQosConfigResponse{
		GlobalQosConfig: obj.(*models.GlobalQosConfig),
	}
	return response, nil
}
