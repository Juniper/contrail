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

// ConfigRootIntent
//   A struct to store attributes related to ConfigRoot
//   needed by Intent Compiler
type ConfigRootIntent struct {
	Uuid string
}

// EvaluateConfigRoot - evaluates the ConfigRoot
func EvaluateConfigRoot(obj interface{}) {
	resourceObj := obj.(ConfigRootIntent)
	log.Println("EvaluateConfigRoot Called ", resourceObj)
}

// CreateConfigRoot handles create request
func (service *PluginService) CreateConfigRoot(ctx context.Context, request *services.CreateConfigRootRequest) (*services.CreateConfigRootResponse, error) {
	log.Println(" CreateConfigRoot Entered")

	obj := request.GetConfigRoot()

	intentObj := ConfigRootIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ConfigRootIntent"); !ok {
		compilationif.ObjsCache.Store("ConfigRootIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ConfigRootIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateConfigRoot", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ConfigRoot")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateConfigRoot(ctx, request)
}

// UpdateConfigRoot handles update request
func (service *PluginService) UpdateConfigRoot(ctx context.Context, request *services.UpdateConfigRootRequest) (*services.UpdateConfigRootResponse, error) {
	log.Println(" UpdateConfigRoot ENTERED")

	obj := request.GetConfigRoot()

	intentObj := ConfigRootIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ConfigRootIntent"); !ok {
		compilationif.ObjsCache.Store("ConfigRootIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ConfigRoot")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateConfigRoot", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateConfigRoot(ctx, request)
}

// DeleteConfigRoot handles delete request
func (service *PluginService) DeleteConfigRoot(ctx context.Context, request *services.DeleteConfigRootRequest) (*services.DeleteConfigRootResponse, error) {
	log.Println(" DeleteConfigRoot ENTERED")

	objUUID := request.GetID()

	//intentObj := ConfigRootIntent {
	//ConfigRoot: *obj,
	//}

	//EvaluateDependencies(intentObj, "ConfigRoot")

	objMap, ok := compilationif.ObjsCache.Load("ConfigRootIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteConfigRoot", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteConfigRoot(ctx, request)
}

// GetConfigRoot handles get request
func (service *PluginService) GetConfigRoot(ctx context.Context, request *services.GetConfigRootRequest) (*services.GetConfigRootResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ConfigRoot")
	if !ok {
		return nil, errors.New("ConfigRoot get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ConfigRoot get failed ")
	}

	response := &services.GetConfigRootResponse{
		ConfigRoot: obj.(*models.ConfigRoot),
	}
	return response, nil
}
