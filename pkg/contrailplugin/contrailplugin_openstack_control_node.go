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

// OpenstackControlNodeIntent
//   A struct to store attributes related to OpenstackControlNode
//   needed by Intent Compiler
type OpenstackControlNodeIntent struct {
	Uuid string
}

// EvaluateOpenstackControlNode - evaluates the OpenstackControlNode
func EvaluateOpenstackControlNode(obj interface{}) {
	resourceObj := obj.(OpenstackControlNodeIntent)
	log.Println("EvaluateOpenstackControlNode Called ", resourceObj)
}

// CreateOpenstackControlNode handles create request
func (service *PluginService) CreateOpenstackControlNode(ctx context.Context, request *services.CreateOpenstackControlNodeRequest) (*services.CreateOpenstackControlNodeResponse, error) {
	log.Println(" CreateOpenstackControlNode Entered")

	obj := request.GetOpenstackControlNode()

	intentObj := OpenstackControlNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackControlNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackControlNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("OpenstackControlNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateOpenstackControlNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "OpenstackControlNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateOpenstackControlNode(ctx, request)
}

// UpdateOpenstackControlNode handles update request
func (service *PluginService) UpdateOpenstackControlNode(ctx context.Context, request *services.UpdateOpenstackControlNodeRequest) (*services.UpdateOpenstackControlNodeResponse, error) {
	log.Println(" UpdateOpenstackControlNode ENTERED")

	obj := request.GetOpenstackControlNode()

	intentObj := OpenstackControlNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackControlNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackControlNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "OpenstackControlNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateOpenstackControlNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateOpenstackControlNode(ctx, request)
}

// DeleteOpenstackControlNode handles delete request
func (service *PluginService) DeleteOpenstackControlNode(ctx context.Context, request *services.DeleteOpenstackControlNodeRequest) (*services.DeleteOpenstackControlNodeResponse, error) {
	log.Println(" DeleteOpenstackControlNode ENTERED")

	objUUID := request.GetID()

	//intentObj := OpenstackControlNodeIntent {
	//OpenstackControlNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "OpenstackControlNode")

	objMap, ok := compilationif.ObjsCache.Load("OpenstackControlNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteOpenstackControlNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteOpenstackControlNode(ctx, request)
}

// GetOpenstackControlNode handles get request
func (service *PluginService) GetOpenstackControlNode(ctx context.Context, request *services.GetOpenstackControlNodeRequest) (*services.GetOpenstackControlNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("OpenstackControlNode")
	if !ok {
		return nil, errors.New("OpenstackControlNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("OpenstackControlNode get failed ")
	}

	response := &services.GetOpenstackControlNodeResponse{
		OpenstackControlNode: obj.(*models.OpenstackControlNode),
	}
	return response, nil
}
