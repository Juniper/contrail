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

// GlobalAnalyticsConfigIntent
//   A struct to store attributes related to GlobalAnalyticsConfig
//   needed by Intent Compiler
type GlobalAnalyticsConfigIntent struct {
	Uuid string
}

// EvaluateGlobalAnalyticsConfig - evaluates the GlobalAnalyticsConfig
func EvaluateGlobalAnalyticsConfig(obj interface{}) {
	resourceObj := obj.(GlobalAnalyticsConfigIntent)
	log.Println("EvaluateGlobalAnalyticsConfig Called ", resourceObj)
}

// CreateGlobalAnalyticsConfig handles create request
func (service *PluginService) CreateGlobalAnalyticsConfig(ctx context.Context, request *services.CreateGlobalAnalyticsConfigRequest) (*services.CreateGlobalAnalyticsConfigResponse, error) {
	log.Println(" CreateGlobalAnalyticsConfig Entered")

	obj := request.GetGlobalAnalyticsConfig()

	intentObj := GlobalAnalyticsConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalAnalyticsConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalAnalyticsConfigIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("GlobalAnalyticsConfigIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateGlobalAnalyticsConfig", objMap.(*sync.Map))

	EvaluateDependencies(obj, "GlobalAnalyticsConfig")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateGlobalAnalyticsConfig(ctx, request)
}

// UpdateGlobalAnalyticsConfig handles update request
func (service *PluginService) UpdateGlobalAnalyticsConfig(ctx context.Context, request *services.UpdateGlobalAnalyticsConfigRequest) (*services.UpdateGlobalAnalyticsConfigResponse, error) {
	log.Println(" UpdateGlobalAnalyticsConfig ENTERED")

	obj := request.GetGlobalAnalyticsConfig()

	intentObj := GlobalAnalyticsConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("GlobalAnalyticsConfigIntent"); !ok {
		compilationif.ObjsCache.Store("GlobalAnalyticsConfigIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "GlobalAnalyticsConfig")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateGlobalAnalyticsConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateGlobalAnalyticsConfig(ctx, request)
}

// DeleteGlobalAnalyticsConfig handles delete request
func (service *PluginService) DeleteGlobalAnalyticsConfig(ctx context.Context, request *services.DeleteGlobalAnalyticsConfigRequest) (*services.DeleteGlobalAnalyticsConfigResponse, error) {
	log.Println(" DeleteGlobalAnalyticsConfig ENTERED")

	objUUID := request.GetID()

	//intentObj := GlobalAnalyticsConfigIntent {
	//GlobalAnalyticsConfig: *obj,
	//}

	//EvaluateDependencies(intentObj, "GlobalAnalyticsConfig")

	objMap, ok := compilationif.ObjsCache.Load("GlobalAnalyticsConfigIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteGlobalAnalyticsConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteGlobalAnalyticsConfig(ctx, request)
}

// GetGlobalAnalyticsConfig handles get request
func (service *PluginService) GetGlobalAnalyticsConfig(ctx context.Context, request *services.GetGlobalAnalyticsConfigRequest) (*services.GetGlobalAnalyticsConfigResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("GlobalAnalyticsConfig")
	if !ok {
		return nil, errors.New("GlobalAnalyticsConfig get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("GlobalAnalyticsConfig get failed ")
	}

	response := &services.GetGlobalAnalyticsConfigResponse{
		GlobalAnalyticsConfig: obj.(*models.GlobalAnalyticsConfig),
	}
	return response, nil
}
