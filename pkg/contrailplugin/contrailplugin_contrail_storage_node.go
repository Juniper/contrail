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

// ContrailStorageNodeIntent
//   A struct to store attributes related to ContrailStorageNode
//   needed by Intent Compiler
type ContrailStorageNodeIntent struct {
	Uuid string
}

// EvaluateContrailStorageNode - evaluates the ContrailStorageNode
func EvaluateContrailStorageNode(obj interface{}) {
	resourceObj := obj.(ContrailStorageNodeIntent)
	log.Println("EvaluateContrailStorageNode Called ", resourceObj)
}

// CreateContrailStorageNode handles create request
func (service *PluginService) CreateContrailStorageNode(ctx context.Context, request *services.CreateContrailStorageNodeRequest) (*services.CreateContrailStorageNodeResponse, error) {
	log.Println(" CreateContrailStorageNode Entered")

	obj := request.GetContrailStorageNode()

	intentObj := ContrailStorageNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailStorageNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailStorageNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailStorageNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailStorageNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailStorageNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailStorageNode(ctx, request)
}

// UpdateContrailStorageNode handles update request
func (service *PluginService) UpdateContrailStorageNode(ctx context.Context, request *services.UpdateContrailStorageNodeRequest) (*services.UpdateContrailStorageNodeResponse, error) {
	log.Println(" UpdateContrailStorageNode ENTERED")

	obj := request.GetContrailStorageNode()

	intentObj := ContrailStorageNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailStorageNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailStorageNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailStorageNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailStorageNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailStorageNode(ctx, request)
}

// DeleteContrailStorageNode handles delete request
func (service *PluginService) DeleteContrailStorageNode(ctx context.Context, request *services.DeleteContrailStorageNodeRequest) (*services.DeleteContrailStorageNodeResponse, error) {
	log.Println(" DeleteContrailStorageNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailStorageNodeIntent {
	//ContrailStorageNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailStorageNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailStorageNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailStorageNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailStorageNode(ctx, request)
}

// GetContrailStorageNode handles get request
func (service *PluginService) GetContrailStorageNode(ctx context.Context, request *services.GetContrailStorageNodeRequest) (*services.GetContrailStorageNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailStorageNode")
	if !ok {
		return nil, errors.New("ContrailStorageNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailStorageNode get failed ")
	}

	response := &services.GetContrailStorageNodeResponse{
		ContrailStorageNode: obj.(*models.ContrailStorageNode),
	}
	return response, nil
}
