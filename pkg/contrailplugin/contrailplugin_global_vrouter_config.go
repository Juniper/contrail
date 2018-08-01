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

// GlobalVrouterConfigIntent
//   A struct to store attributes related to GlobalVrouterConfig
//   needed by Intent Compiler
type GlobalVrouterConfigIntent struct {
	Uuid string
}

// EvaluateGlobalVrouterConfig - evaluates the GlobalVrouterConfig
func EvaluateGlobalVrouterConfig(obj interface{}) {
	resourceObj := obj.(GlobalVrouterConfigIntent)
	log.Println("EvaluateGlobalVrouterConfig Called ", resourceObj)
}

// CreateGlobalVrouterConfig handles create request
func (service *PluginService) CreateGlobalVrouterConfig(ctx context.Context, request *services.CreateGlobalVrouterConfigRequest) (*services.CreateGlobalVrouterConfigResponse, error) {
	log.Println(" CreateGlobalVrouterConfig Entered")

	obj := request.GetGlobalVrouterConfig()

	intentObj := GlobalVrouterConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalVrouterConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalVrouterConfigIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("GlobalVrouterConfigIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateGlobalVrouterConfig", objMap.(*sync.Map))

	EvaluateDependencies(obj, "GlobalVrouterConfig")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalVrouterConfig(ctx, request)
}

// UpdateGlobalVrouterConfig handles update request
func (service *PluginService) UpdateGlobalVrouterConfig(ctx context.Context, request *services.UpdateGlobalVrouterConfigRequest) (*services.UpdateGlobalVrouterConfigResponse, error) {
	log.Println(" UpdateGlobalVrouterConfig ENTERED")

	obj := request.GetGlobalVrouterConfig()

	intentObj := GlobalVrouterConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalVrouterConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalVrouterConfigIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "GlobalVrouterConfig")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateGlobalVrouterConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalVrouterConfig(ctx, request)
}

// DeleteGlobalVrouterConfig handles delete request
func (service *PluginService) DeleteGlobalVrouterConfig(ctx context.Context, request *services.DeleteGlobalVrouterConfigRequest) (*services.DeleteGlobalVrouterConfigResponse, error) {
	log.Println(" DeleteGlobalVrouterConfig ENTERED")

	objUUID := request.GetID()

	//intentObj := GlobalVrouterConfigIntent {
	//GlobalVrouterConfig: *obj,
	//}

	//EvaluateDependencies(intentObj, "GlobalVrouterConfig")

	objMap, ok := compilationif.ObjsCache.Load("GlobalVrouterConfigIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteGlobalVrouterConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalVrouterConfig(ctx, request)
}

// GetGlobalVrouterConfig handles get request
func (service *PluginService) GetGlobalVrouterConfig(ctx context.Context, request *services.GetGlobalVrouterConfigRequest) (*services.GetGlobalVrouterConfigResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("GlobalVrouterConfig")
	if !ok {
		return nil, errors.New("GlobalVrouterConfig get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("GlobalVrouterConfig get failed ")
	}

	response := &services.GetGlobalVrouterConfigResponse{
		GlobalVrouterConfig: obj.(*models.GlobalVrouterConfig),
	}
	return response, nil
}
