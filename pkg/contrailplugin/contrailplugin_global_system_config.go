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

// GlobalSystemConfigIntent
//   A struct to store attributes related to GlobalSystemConfig
//   needed by Intent Compiler
type GlobalSystemConfigIntent struct {
	Uuid string
}

// EvaluateGlobalSystemConfig - evaluates the GlobalSystemConfig
func EvaluateGlobalSystemConfig(obj interface{}) {
	resourceObj := obj.(GlobalSystemConfigIntent)
	log.Println("EvaluateGlobalSystemConfig Called ", resourceObj)
}

// CreateGlobalSystemConfig handles create request
func (service *PluginService) CreateGlobalSystemConfig(ctx context.Context, request *services.CreateGlobalSystemConfigRequest) (*services.CreateGlobalSystemConfigResponse, error) {
	log.Println(" CreateGlobalSystemConfig Entered")

	obj := request.GetGlobalSystemConfig()

	intentObj := GlobalSystemConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalSystemConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalSystemConfigIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("GlobalSystemConfigIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateGlobalSystemConfig", objMap.(*sync.Map))

	EvaluateDependencies(obj, "GlobalSystemConfig")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalSystemConfig(ctx, request)
}

// UpdateGlobalSystemConfig handles update request
func (service *PluginService) UpdateGlobalSystemConfig(ctx context.Context, request *services.UpdateGlobalSystemConfigRequest) (*services.UpdateGlobalSystemConfigResponse, error) {
	log.Println(" UpdateGlobalSystemConfig ENTERED")

	obj := request.GetGlobalSystemConfig()

	intentObj := GlobalSystemConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalSystemConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalSystemConfigIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "GlobalSystemConfig")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateGlobalSystemConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalSystemConfig(ctx, request)
}

// DeleteGlobalSystemConfig handles delete request
func (service *PluginService) DeleteGlobalSystemConfig(ctx context.Context, request *services.DeleteGlobalSystemConfigRequest) (*services.DeleteGlobalSystemConfigResponse, error) {
	log.Println(" DeleteGlobalSystemConfig ENTERED")

	objUUID := request.GetID()

	//intentObj := GlobalSystemConfigIntent {
	//GlobalSystemConfig: *obj,
	//}

	//EvaluateDependencies(intentObj, "GlobalSystemConfig")

	objMap, ok := compilationif.ObjsCache.Load("GlobalSystemConfigIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteGlobalSystemConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalSystemConfig(ctx, request)
}

// GetGlobalSystemConfig handles get request
func (service *PluginService) GetGlobalSystemConfig(ctx context.Context, request *services.GetGlobalSystemConfigRequest) (*services.GetGlobalSystemConfigResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("GlobalSystemConfig")
	if !ok {
		return nil, errors.New("GlobalSystemConfig get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("GlobalSystemConfig get failed ")
	}

	response := &services.GetGlobalSystemConfigResponse{
		GlobalSystemConfig: obj.(*models.GlobalSystemConfig),
	}
	return response, nil
}
