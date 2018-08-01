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

// NodeIntent
//   A struct to store attributes related to Node
//   needed by Intent Compiler
type NodeIntent struct {
	Uuid string
}

// EvaluateNode - evaluates the Node
func EvaluateNode(obj interface{}) {
	resourceObj := obj.(NodeIntent)
	log.Println("EvaluateNode Called ", resourceObj)
}

// CreateNode handles create request
func (service *PluginService) CreateNode(ctx context.Context, request *services.CreateNodeRequest) (*services.CreateNodeResponse, error) {
	log.Println(" CreateNode Entered")

	obj := request.GetNode()

	intentObj := NodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NodeIntent"); !ok {
		compilationif.ObjsCache.Store("NodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("NodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Node")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateNode(ctx, request)
}

// UpdateNode handles update request
func (service *PluginService) UpdateNode(ctx context.Context, request *services.UpdateNodeRequest) (*services.UpdateNodeResponse, error) {
	log.Println(" UpdateNode ENTERED")

	obj := request.GetNode()

	intentObj := NodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NodeIntent"); !ok {
		compilationif.ObjsCache.Store("NodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Node")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateNode(ctx, request)
}

// DeleteNode handles delete request
func (service *PluginService) DeleteNode(ctx context.Context, request *services.DeleteNodeRequest) (*services.DeleteNodeResponse, error) {
	log.Println(" DeleteNode ENTERED")

	objUUID := request.GetID()

	//intentObj := NodeIntent {
	//Node: *obj,
	//}

	//EvaluateDependencies(intentObj, "Node")

	objMap, ok := compilationif.ObjsCache.Load("NodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteNode(ctx, request)
}

// GetNode handles get request
func (service *PluginService) GetNode(ctx context.Context, request *services.GetNodeRequest) (*services.GetNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Node")
	if !ok {
		return nil, errors.New("Node get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Node get failed ")
	}

	response := &services.GetNodeResponse{
		Node: obj.(*models.Node),
	}
	return response, nil
}
