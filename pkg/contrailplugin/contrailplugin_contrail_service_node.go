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

// ContrailServiceNodeIntent
//   A struct to store attributes related to ContrailServiceNode
//   needed by Intent Compiler
type ContrailServiceNodeIntent struct {
	Uuid string
}

// EvaluateContrailServiceNode - evaluates the ContrailServiceNode
func EvaluateContrailServiceNode(obj interface{}) {
	resourceObj := obj.(ContrailServiceNodeIntent)
	log.Println("EvaluateContrailServiceNode Called ", resourceObj)
}

// CreateContrailServiceNode handles create request
func (service *PluginService) CreateContrailServiceNode(ctx context.Context, request *services.CreateContrailServiceNodeRequest) (*services.CreateContrailServiceNodeResponse, error) {
	log.Println(" CreateContrailServiceNode Entered")

	obj := request.GetContrailServiceNode()

	intentObj := ContrailServiceNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailServiceNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailServiceNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailServiceNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailServiceNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailServiceNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailServiceNode(ctx, request)
}

// UpdateContrailServiceNode handles update request
func (service *PluginService) UpdateContrailServiceNode(ctx context.Context, request *services.UpdateContrailServiceNodeRequest) (*services.UpdateContrailServiceNodeResponse, error) {
	log.Println(" UpdateContrailServiceNode ENTERED")

	obj := request.GetContrailServiceNode()

	intentObj := ContrailServiceNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailServiceNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailServiceNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailServiceNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailServiceNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailServiceNode(ctx, request)
}

// DeleteContrailServiceNode handles delete request
func (service *PluginService) DeleteContrailServiceNode(ctx context.Context, request *services.DeleteContrailServiceNodeRequest) (*services.DeleteContrailServiceNodeResponse, error) {
	log.Println(" DeleteContrailServiceNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailServiceNodeIntent {
	//ContrailServiceNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailServiceNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailServiceNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailServiceNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailServiceNode(ctx, request)
}

// GetContrailServiceNode handles get request
func (service *PluginService) GetContrailServiceNode(ctx context.Context, request *services.GetContrailServiceNodeRequest) (*services.GetContrailServiceNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailServiceNode")
	if !ok {
		return nil, errors.New("ContrailServiceNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailServiceNode get failed ")
	}

	response := &services.GetContrailServiceNodeResponse{
		ContrailServiceNode: obj.(*models.ContrailServiceNode),
	}
	return response, nil
}
