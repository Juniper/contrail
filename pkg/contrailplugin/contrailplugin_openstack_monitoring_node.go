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

// OpenstackMonitoringNodeIntent
//   A struct to store attributes related to OpenstackMonitoringNode
//   needed by Intent Compiler
type OpenstackMonitoringNodeIntent struct {
	Uuid string
}

// EvaluateOpenstackMonitoringNode - evaluates the OpenstackMonitoringNode
func EvaluateOpenstackMonitoringNode(obj interface{}) {
	resourceObj := obj.(OpenstackMonitoringNodeIntent)
	log.Println("EvaluateOpenstackMonitoringNode Called ", resourceObj)
}

// CreateOpenstackMonitoringNode handles create request
func (service *PluginService) CreateOpenstackMonitoringNode(ctx context.Context, request *services.CreateOpenstackMonitoringNodeRequest) (*services.CreateOpenstackMonitoringNodeResponse, error) {
	log.Println(" CreateOpenstackMonitoringNode Entered")

	obj := request.GetOpenstackMonitoringNode()

	intentObj := OpenstackMonitoringNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackMonitoringNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackMonitoringNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("OpenstackMonitoringNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateOpenstackMonitoringNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "OpenstackMonitoringNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateOpenstackMonitoringNode(ctx, request)
}

// UpdateOpenstackMonitoringNode handles update request
func (service *PluginService) UpdateOpenstackMonitoringNode(ctx context.Context, request *services.UpdateOpenstackMonitoringNodeRequest) (*services.UpdateOpenstackMonitoringNodeResponse, error) {
	log.Println(" UpdateOpenstackMonitoringNode ENTERED")

	obj := request.GetOpenstackMonitoringNode()

	intentObj := OpenstackMonitoringNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("OpenstackMonitoringNodeIntent"); !ok {
		compilationif.ObjsCache.Store("OpenstackMonitoringNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "OpenstackMonitoringNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateOpenstackMonitoringNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateOpenstackMonitoringNode(ctx, request)
}

// DeleteOpenstackMonitoringNode handles delete request
func (service *PluginService) DeleteOpenstackMonitoringNode(ctx context.Context, request *services.DeleteOpenstackMonitoringNodeRequest) (*services.DeleteOpenstackMonitoringNodeResponse, error) {
	log.Println(" DeleteOpenstackMonitoringNode ENTERED")

	objUUID := request.GetID()

	//intentObj := OpenstackMonitoringNodeIntent {
	//OpenstackMonitoringNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "OpenstackMonitoringNode")

	objMap, ok := compilationif.ObjsCache.Load("OpenstackMonitoringNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteOpenstackMonitoringNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteOpenstackMonitoringNode(ctx, request)
}

// GetOpenstackMonitoringNode handles get request
func (service *PluginService) GetOpenstackMonitoringNode(ctx context.Context, request *services.GetOpenstackMonitoringNodeRequest) (*services.GetOpenstackMonitoringNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("OpenstackMonitoringNode")
	if !ok {
		return nil, errors.New("OpenstackMonitoringNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("OpenstackMonitoringNode get failed ")
	}

	response := &services.GetOpenstackMonitoringNodeResponse{
		OpenstackMonitoringNode: obj.(*models.OpenstackMonitoringNode),
	}
	return response, nil
}
