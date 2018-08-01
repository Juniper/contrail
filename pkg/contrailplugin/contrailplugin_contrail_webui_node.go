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

// ContrailWebuiNodeIntent
//   A struct to store attributes related to ContrailWebuiNode
//   needed by Intent Compiler
type ContrailWebuiNodeIntent struct {
	Uuid string
}

// EvaluateContrailWebuiNode - evaluates the ContrailWebuiNode
func EvaluateContrailWebuiNode(obj interface{}) {
	resourceObj := obj.(ContrailWebuiNodeIntent)
	log.Println("EvaluateContrailWebuiNode Called ", resourceObj)
}

// CreateContrailWebuiNode handles create request
func (service *PluginService) CreateContrailWebuiNode(ctx context.Context, request *services.CreateContrailWebuiNodeRequest) (*services.CreateContrailWebuiNodeResponse, error) {
	log.Println(" CreateContrailWebuiNode Entered")

	obj := request.GetContrailWebuiNode()

	intentObj := ContrailWebuiNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailWebuiNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailWebuiNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailWebuiNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailWebuiNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailWebuiNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailWebuiNode(ctx, request)
}

// UpdateContrailWebuiNode handles update request
func (service *PluginService) UpdateContrailWebuiNode(ctx context.Context, request *services.UpdateContrailWebuiNodeRequest) (*services.UpdateContrailWebuiNodeResponse, error) {
	log.Println(" UpdateContrailWebuiNode ENTERED")

	obj := request.GetContrailWebuiNode()

	intentObj := ContrailWebuiNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailWebuiNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailWebuiNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailWebuiNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailWebuiNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailWebuiNode(ctx, request)
}

// DeleteContrailWebuiNode handles delete request
func (service *PluginService) DeleteContrailWebuiNode(ctx context.Context, request *services.DeleteContrailWebuiNodeRequest) (*services.DeleteContrailWebuiNodeResponse, error) {
	log.Println(" DeleteContrailWebuiNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailWebuiNodeIntent {
	//ContrailWebuiNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailWebuiNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailWebuiNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailWebuiNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailWebuiNode(ctx, request)
}

// GetContrailWebuiNode handles get request
func (service *PluginService) GetContrailWebuiNode(ctx context.Context, request *services.GetContrailWebuiNodeRequest) (*services.GetContrailWebuiNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailWebuiNode")
	if !ok {
		return nil, errors.New("ContrailWebuiNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailWebuiNode get failed ")
	}

	response := &services.GetContrailWebuiNodeResponse{
		ContrailWebuiNode: obj.(*models.ContrailWebuiNode),
	}
	return response, nil
}
