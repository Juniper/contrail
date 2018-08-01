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

// OpenstackNetworkNodeIntent
//   A struct to store attributes related to OpenstackNetworkNode
//   needed by Intent Compiler
type OpenstackNetworkNodeIntent struct {
	Uuid string
}

// EvaluateOpenstackNetworkNode - evaluates the OpenstackNetworkNode
func EvaluateOpenstackNetworkNode(obj interface{}) {
	resourceObj := obj.(OpenstackNetworkNodeIntent)
	log.Println("EvaluateOpenstackNetworkNode Called ", resourceObj)
}

// CreateOpenstackNetworkNode handles create request
func (service *PluginService) CreateOpenstackNetworkNode(ctx context.Context, request *services.CreateOpenstackNetworkNodeRequest) (*services.CreateOpenstackNetworkNodeResponse, error) {
	log.Println(" CreateOpenstackNetworkNode Entered")

	obj := request.GetOpenstackNetworkNode()

	intentObj := OpenstackNetworkNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackNetworkNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackNetworkNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("OpenstackNetworkNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateOpenstackNetworkNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "OpenstackNetworkNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateOpenstackNetworkNode(ctx, request)
}

// UpdateOpenstackNetworkNode handles update request
func (service *PluginService) UpdateOpenstackNetworkNode(ctx context.Context, request *services.UpdateOpenstackNetworkNodeRequest) (*services.UpdateOpenstackNetworkNodeResponse, error) {
	log.Println(" UpdateOpenstackNetworkNode ENTERED")

	obj := request.GetOpenstackNetworkNode()

	intentObj := OpenstackNetworkNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackNetworkNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackNetworkNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "OpenstackNetworkNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateOpenstackNetworkNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateOpenstackNetworkNode(ctx, request)
}

// DeleteOpenstackNetworkNode handles delete request
func (service *PluginService) DeleteOpenstackNetworkNode(ctx context.Context, request *services.DeleteOpenstackNetworkNodeRequest) (*services.DeleteOpenstackNetworkNodeResponse, error) {
	log.Println(" DeleteOpenstackNetworkNode ENTERED")

	objUUID := request.GetID()

	//intentObj := OpenstackNetworkNodeIntent {
	//OpenstackNetworkNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "OpenstackNetworkNode")

	objMap, ok := compilationif.ObjsCache.Load("OpenstackNetworkNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteOpenstackNetworkNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteOpenstackNetworkNode(ctx, request)
}

// GetOpenstackNetworkNode handles get request
func (service *PluginService) GetOpenstackNetworkNode(ctx context.Context, request *services.GetOpenstackNetworkNodeRequest) (*services.GetOpenstackNetworkNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("OpenstackNetworkNode")
	if !ok {
		return nil, errors.New("OpenstackNetworkNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("OpenstackNetworkNode get failed ")
	}

	response := &services.GetOpenstackNetworkNodeResponse{
		OpenstackNetworkNode: obj.(*models.OpenstackNetworkNode),
	}
	return response, nil
}
