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

// ContrailControlNodeIntent
//   A struct to store attributes related to ContrailControlNode
//   needed by Intent Compiler
type ContrailControlNodeIntent struct {
	Uuid string
}

// EvaluateContrailControlNode - evaluates the ContrailControlNode
func EvaluateContrailControlNode(obj interface{}) {
	resourceObj := obj.(ContrailControlNodeIntent)
	log.Println("EvaluateContrailControlNode Called ", resourceObj)
}

// CreateContrailControlNode handles create request
func (service *PluginService) CreateContrailControlNode(ctx context.Context, request *services.CreateContrailControlNodeRequest) (*services.CreateContrailControlNodeResponse, error) {
	log.Println(" CreateContrailControlNode Entered")

	obj := request.GetContrailControlNode()

	intentObj := ContrailControlNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailControlNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailControlNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailControlNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailControlNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailControlNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailControlNode(ctx, request)
}

// UpdateContrailControlNode handles update request
func (service *PluginService) UpdateContrailControlNode(ctx context.Context, request *services.UpdateContrailControlNodeRequest) (*services.UpdateContrailControlNodeResponse, error) {
	log.Println(" UpdateContrailControlNode ENTERED")

	obj := request.GetContrailControlNode()

	intentObj := ContrailControlNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailControlNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailControlNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailControlNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailControlNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailControlNode(ctx, request)
}

// DeleteContrailControlNode handles delete request
func (service *PluginService) DeleteContrailControlNode(ctx context.Context, request *services.DeleteContrailControlNodeRequest) (*services.DeleteContrailControlNodeResponse, error) {
	log.Println(" DeleteContrailControlNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailControlNodeIntent {
	//ContrailControlNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailControlNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailControlNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailControlNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailControlNode(ctx, request)
}

// GetContrailControlNode handles get request
func (service *PluginService) GetContrailControlNode(ctx context.Context, request *services.GetContrailControlNodeRequest) (*services.GetContrailControlNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailControlNode")
	if !ok {
		return nil, errors.New("ContrailControlNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailControlNode get failed ")
	}

	response := &services.GetContrailControlNodeResponse{
		ContrailControlNode: obj.(*models.ContrailControlNode),
	}
	return response, nil
}
