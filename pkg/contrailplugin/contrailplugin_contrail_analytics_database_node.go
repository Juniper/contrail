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

// ContrailAnalyticsDatabaseNodeIntent
//   A struct to store attributes related to ContrailAnalyticsDatabaseNode
//   needed by Intent Compiler
type ContrailAnalyticsDatabaseNodeIntent struct {
	Uuid string
}

// EvaluateContrailAnalyticsDatabaseNode - evaluates the ContrailAnalyticsDatabaseNode
func EvaluateContrailAnalyticsDatabaseNode(obj interface{}) {
	resourceObj := obj.(ContrailAnalyticsDatabaseNodeIntent)
	log.Println("EvaluateContrailAnalyticsDatabaseNode Called ", resourceObj)
}

// CreateContrailAnalyticsDatabaseNode handles create request
func (service *PluginService) CreateContrailAnalyticsDatabaseNode(ctx context.Context, request *services.CreateContrailAnalyticsDatabaseNodeRequest) (*services.CreateContrailAnalyticsDatabaseNodeResponse, error) {
	log.Println(" CreateContrailAnalyticsDatabaseNode Entered")

	obj := request.GetContrailAnalyticsDatabaseNode()

	intentObj := ContrailAnalyticsDatabaseNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailAnalyticsDatabaseNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailAnalyticsDatabaseNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailAnalyticsDatabaseNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailAnalyticsDatabaseNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailAnalyticsDatabaseNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailAnalyticsDatabaseNode(ctx, request)
}

// UpdateContrailAnalyticsDatabaseNode handles update request
func (service *PluginService) UpdateContrailAnalyticsDatabaseNode(ctx context.Context, request *services.UpdateContrailAnalyticsDatabaseNodeRequest) (*services.UpdateContrailAnalyticsDatabaseNodeResponse, error) {
	log.Println(" UpdateContrailAnalyticsDatabaseNode ENTERED")

	obj := request.GetContrailAnalyticsDatabaseNode()

	intentObj := ContrailAnalyticsDatabaseNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailAnalyticsDatabaseNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailAnalyticsDatabaseNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailAnalyticsDatabaseNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailAnalyticsDatabaseNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailAnalyticsDatabaseNode(ctx, request)
}

// DeleteContrailAnalyticsDatabaseNode handles delete request
func (service *PluginService) DeleteContrailAnalyticsDatabaseNode(ctx context.Context, request *services.DeleteContrailAnalyticsDatabaseNodeRequest) (*services.DeleteContrailAnalyticsDatabaseNodeResponse, error) {
	log.Println(" DeleteContrailAnalyticsDatabaseNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailAnalyticsDatabaseNodeIntent {
	//ContrailAnalyticsDatabaseNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailAnalyticsDatabaseNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailAnalyticsDatabaseNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailAnalyticsDatabaseNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailAnalyticsDatabaseNode(ctx, request)
}

// GetContrailAnalyticsDatabaseNode handles get request
func (service *PluginService) GetContrailAnalyticsDatabaseNode(ctx context.Context, request *services.GetContrailAnalyticsDatabaseNodeRequest) (*services.GetContrailAnalyticsDatabaseNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailAnalyticsDatabaseNode")
	if !ok {
		return nil, errors.New("ContrailAnalyticsDatabaseNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailAnalyticsDatabaseNode get failed ")
	}

	response := &services.GetContrailAnalyticsDatabaseNodeResponse{
		ContrailAnalyticsDatabaseNode: obj.(*models.ContrailAnalyticsDatabaseNode),
	}
	return response, nil
}
