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

// ConfigNodeIntent
//   A struct to store attributes related to ConfigNode
//   needed by Intent Compiler
type ConfigNodeIntent struct {
	Uuid string
}

// EvaluateConfigNode - evaluates the ConfigNode
func EvaluateConfigNode(obj interface{}) {
	resourceObj := obj.(ConfigNodeIntent)
	log.Println("EvaluateConfigNode Called ", resourceObj)
}

// CreateConfigNode handles create request
func (service *PluginService) CreateConfigNode(ctx context.Context, request *services.CreateConfigNodeRequest) (*services.CreateConfigNodeResponse, error) {
	log.Println(" CreateConfigNode Entered")

	obj := request.GetConfigNode()

	intentObj := ConfigNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ConfigNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ConfigNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ConfigNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateConfigNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ConfigNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateConfigNode(ctx, request)
}

// UpdateConfigNode handles update request
func (service *PluginService) UpdateConfigNode(ctx context.Context, request *services.UpdateConfigNodeRequest) (*services.UpdateConfigNodeResponse, error) {
	log.Println(" UpdateConfigNode ENTERED")

	obj := request.GetConfigNode()

	intentObj := ConfigNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ConfigNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ConfigNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ConfigNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateConfigNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateConfigNode(ctx, request)
}

// DeleteConfigNode handles delete request
func (service *PluginService) DeleteConfigNode(ctx context.Context, request *services.DeleteConfigNodeRequest) (*services.DeleteConfigNodeResponse, error) {
	log.Println(" DeleteConfigNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ConfigNodeIntent {
	//ConfigNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ConfigNode")

	objMap, ok := compilationif.ObjsCache.Load("ConfigNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteConfigNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteConfigNode(ctx, request)
}

// GetConfigNode handles get request
func (service *PluginService) GetConfigNode(ctx context.Context, request *services.GetConfigNodeRequest) (*services.GetConfigNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ConfigNode")
	if !ok {
		return nil, errors.New("ConfigNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ConfigNode get failed ")
	}

	response := &services.GetConfigNodeResponse{
		ConfigNode: obj.(*models.ConfigNode),
	}
	return response, nil
}
