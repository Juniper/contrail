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

// AppformixNodeIntent
//   A struct to store attributes related to AppformixNode
//   needed by Intent Compiler
type AppformixNodeIntent struct {
	Uuid string
}

// EvaluateAppformixNode - evaluates the AppformixNode
func EvaluateAppformixNode(obj interface{}) {
	resourceObj := obj.(AppformixNodeIntent)
	log.Println("EvaluateAppformixNode Called ", resourceObj)
}

// CreateAppformixNode handles create request
func (service *PluginService) CreateAppformixNode(ctx context.Context, request *services.CreateAppformixNodeRequest) (*services.CreateAppformixNodeResponse, error) {
	log.Println(" CreateAppformixNode Entered")

	obj := request.GetAppformixNode()

	intentObj := AppformixNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AppformixNodeIntent"); !ok {
		compilationif.ObjsCache.Store("AppformixNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("AppformixNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAppformixNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "AppformixNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAppformixNode(ctx, request)
}

// UpdateAppformixNode handles update request
func (service *PluginService) UpdateAppformixNode(ctx context.Context, request *services.UpdateAppformixNodeRequest) (*services.UpdateAppformixNodeResponse, error) {
	log.Println(" UpdateAppformixNode ENTERED")

	obj := request.GetAppformixNode()

	intentObj := AppformixNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AppformixNodeIntent"); !ok {
		compilationif.ObjsCache.Store("AppformixNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "AppformixNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAppformixNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAppformixNode(ctx, request)
}

// DeleteAppformixNode handles delete request
func (service *PluginService) DeleteAppformixNode(ctx context.Context, request *services.DeleteAppformixNodeRequest) (*services.DeleteAppformixNodeResponse, error) {
	log.Println(" DeleteAppformixNode ENTERED")

	objUUID := request.GetID()

	//intentObj := AppformixNodeIntent {
	//AppformixNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "AppformixNode")

	objMap, ok := compilationif.ObjsCache.Load("AppformixNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAppformixNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAppformixNode(ctx, request)
}

// GetAppformixNode handles get request
func (service *PluginService) GetAppformixNode(ctx context.Context, request *services.GetAppformixNodeRequest) (*services.GetAppformixNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("AppformixNode")
	if !ok {
		return nil, errors.New("AppformixNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("AppformixNode get failed ")
	}

	response := &services.GetAppformixNodeResponse{
		AppformixNode: obj.(*models.AppformixNode),
	}
	return response, nil
}
