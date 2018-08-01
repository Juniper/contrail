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

// OpenstackComputeNodeIntent
//   A struct to store attributes related to OpenstackComputeNode
//   needed by Intent Compiler
type OpenstackComputeNodeIntent struct {
	Uuid string
}

// EvaluateOpenstackComputeNode - evaluates the OpenstackComputeNode
func EvaluateOpenstackComputeNode(obj interface{}) {
	resourceObj := obj.(OpenstackComputeNodeIntent)
	log.Println("EvaluateOpenstackComputeNode Called ", resourceObj)
}

// CreateOpenstackComputeNode handles create request
func (service *PluginService) CreateOpenstackComputeNode(ctx context.Context, request *services.CreateOpenstackComputeNodeRequest) (*services.CreateOpenstackComputeNodeResponse, error) {
	log.Println(" CreateOpenstackComputeNode Entered")

	obj := request.GetOpenstackComputeNode()

	intentObj := OpenstackComputeNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackComputeNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackComputeNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("OpenstackComputeNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateOpenstackComputeNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "OpenstackComputeNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateOpenstackComputeNode(ctx, request)
}

// UpdateOpenstackComputeNode handles update request
func (service *PluginService) UpdateOpenstackComputeNode(ctx context.Context, request *services.UpdateOpenstackComputeNodeRequest) (*services.UpdateOpenstackComputeNodeResponse, error) {
	log.Println(" UpdateOpenstackComputeNode ENTERED")

	obj := request.GetOpenstackComputeNode()

	intentObj := OpenstackComputeNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackComputeNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackComputeNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "OpenstackComputeNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateOpenstackComputeNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateOpenstackComputeNode(ctx, request)
}

// DeleteOpenstackComputeNode handles delete request
func (service *PluginService) DeleteOpenstackComputeNode(ctx context.Context, request *services.DeleteOpenstackComputeNodeRequest) (*services.DeleteOpenstackComputeNodeResponse, error) {
	log.Println(" DeleteOpenstackComputeNode ENTERED")

	objUUID := request.GetID()

	//intentObj := OpenstackComputeNodeIntent {
	//OpenstackComputeNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "OpenstackComputeNode")

	objMap, ok := compilationif.ObjsCache.Load("OpenstackComputeNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteOpenstackComputeNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteOpenstackComputeNode(ctx, request)
}

// GetOpenstackComputeNode handles get request
func (service *PluginService) GetOpenstackComputeNode(ctx context.Context, request *services.GetOpenstackComputeNodeRequest) (*services.GetOpenstackComputeNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("OpenstackComputeNode")
	if !ok {
		return nil, errors.New("OpenstackComputeNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("OpenstackComputeNode get failed ")
	}

	response := &services.GetOpenstackComputeNodeResponse{
		OpenstackComputeNode: obj.(*models.OpenstackComputeNode),
	}
	return response, nil
}
