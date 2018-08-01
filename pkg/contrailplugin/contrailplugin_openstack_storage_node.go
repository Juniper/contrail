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

// OpenstackStorageNodeIntent
//   A struct to store attributes related to OpenstackStorageNode
//   needed by Intent Compiler
type OpenstackStorageNodeIntent struct {
	Uuid string
}

// EvaluateOpenstackStorageNode - evaluates the OpenstackStorageNode
func EvaluateOpenstackStorageNode(obj interface{}) {
	resourceObj := obj.(OpenstackStorageNodeIntent)
	log.Println("EvaluateOpenstackStorageNode Called ", resourceObj)
}

// CreateOpenstackStorageNode handles create request
func (service *PluginService) CreateOpenstackStorageNode(ctx context.Context, request *services.CreateOpenstackStorageNodeRequest) (*services.CreateOpenstackStorageNodeResponse, error) {
	log.Println(" CreateOpenstackStorageNode Entered")

	obj := request.GetOpenstackStorageNode()

	intentObj := OpenstackStorageNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackStorageNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackStorageNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("OpenstackStorageNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateOpenstackStorageNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "OpenstackStorageNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateOpenstackStorageNode(ctx, request)
}

// UpdateOpenstackStorageNode handles update request
func (service *PluginService) UpdateOpenstackStorageNode(ctx context.Context, request *services.UpdateOpenstackStorageNodeRequest) (*services.UpdateOpenstackStorageNodeResponse, error) {
	log.Println(" UpdateOpenstackStorageNode ENTERED")

	obj := request.GetOpenstackStorageNode()

	intentObj := OpenstackStorageNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackStorageNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackStorageNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "OpenstackStorageNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateOpenstackStorageNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateOpenstackStorageNode(ctx, request)
}

// DeleteOpenstackStorageNode handles delete request
func (service *PluginService) DeleteOpenstackStorageNode(ctx context.Context, request *services.DeleteOpenstackStorageNodeRequest) (*services.DeleteOpenstackStorageNodeResponse, error) {
	log.Println(" DeleteOpenstackStorageNode ENTERED")

	objUUID := request.GetID()

	//intentObj := OpenstackStorageNodeIntent {
	//OpenstackStorageNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "OpenstackStorageNode")

	objMap, ok := compilationif.ObjsCache.Load("OpenstackStorageNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteOpenstackStorageNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteOpenstackStorageNode(ctx, request)
}

// GetOpenstackStorageNode handles get request
func (service *PluginService) GetOpenstackStorageNode(ctx context.Context, request *services.GetOpenstackStorageNodeRequest) (*services.GetOpenstackStorageNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("OpenstackStorageNode")
	if !ok {
		return nil, errors.New("OpenstackStorageNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("OpenstackStorageNode get failed ")
	}

	response := &services.GetOpenstackStorageNodeResponse{
		OpenstackStorageNode: obj.(*models.OpenstackStorageNode),
	}
	return response, nil
}
