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

// ContrailVrouterNodeIntent
//   A struct to store attributes related to ContrailVrouterNode
//   needed by Intent Compiler
type ContrailVrouterNodeIntent struct {
	Uuid string
}

// EvaluateContrailVrouterNode - evaluates the ContrailVrouterNode
func EvaluateContrailVrouterNode(obj interface{}) {
	resourceObj := obj.(ContrailVrouterNodeIntent)
	log.Println("EvaluateContrailVrouterNode Called ", resourceObj)
}

// CreateContrailVrouterNode handles create request
func (service *PluginService) CreateContrailVrouterNode(ctx context.Context, request *services.CreateContrailVrouterNodeRequest) (*services.CreateContrailVrouterNodeResponse, error) {
	log.Println(" CreateContrailVrouterNode Entered")

	obj := request.GetContrailVrouterNode()

	intentObj := ContrailVrouterNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailVrouterNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailVrouterNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailVrouterNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailVrouterNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailVrouterNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailVrouterNode(ctx, request)
}

// UpdateContrailVrouterNode handles update request
func (service *PluginService) UpdateContrailVrouterNode(ctx context.Context, request *services.UpdateContrailVrouterNodeRequest) (*services.UpdateContrailVrouterNodeResponse, error) {
	log.Println(" UpdateContrailVrouterNode ENTERED")

	obj := request.GetContrailVrouterNode()

	intentObj := ContrailVrouterNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailVrouterNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailVrouterNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailVrouterNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailVrouterNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailVrouterNode(ctx, request)
}

// DeleteContrailVrouterNode handles delete request
func (service *PluginService) DeleteContrailVrouterNode(ctx context.Context, request *services.DeleteContrailVrouterNodeRequest) (*services.DeleteContrailVrouterNodeResponse, error) {
	log.Println(" DeleteContrailVrouterNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailVrouterNodeIntent {
	//ContrailVrouterNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailVrouterNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailVrouterNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailVrouterNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailVrouterNode(ctx, request)
}

// GetContrailVrouterNode handles get request
func (service *PluginService) GetContrailVrouterNode(ctx context.Context, request *services.GetContrailVrouterNodeRequest) (*services.GetContrailVrouterNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailVrouterNode")
	if !ok {
		return nil, errors.New("ContrailVrouterNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailVrouterNode get failed ")
	}

	response := &services.GetContrailVrouterNodeResponse{
		ContrailVrouterNode: obj.(*models.ContrailVrouterNode),
	}
	return response, nil
}
