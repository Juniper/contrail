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

// ContrailConfigDatabaseNodeIntent
//   A struct to store attributes related to ContrailConfigDatabaseNode
//   needed by Intent Compiler
type ContrailConfigDatabaseNodeIntent struct {
	Uuid string
}

// EvaluateContrailConfigDatabaseNode - evaluates the ContrailConfigDatabaseNode
func EvaluateContrailConfigDatabaseNode(obj interface{}) {
	resourceObj := obj.(ContrailConfigDatabaseNodeIntent)
	log.Println("EvaluateContrailConfigDatabaseNode Called ", resourceObj)
}

// CreateContrailConfigDatabaseNode handles create request
func (service *PluginService) CreateContrailConfigDatabaseNode(ctx context.Context, request *services.CreateContrailConfigDatabaseNodeRequest) (*services.CreateContrailConfigDatabaseNodeResponse, error) {
	log.Println(" CreateContrailConfigDatabaseNode Entered")

	obj := request.GetContrailConfigDatabaseNode()

	intentObj := ContrailConfigDatabaseNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailConfigDatabaseNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailConfigDatabaseNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailConfigDatabaseNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailConfigDatabaseNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailConfigDatabaseNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailConfigDatabaseNode(ctx, request)
}

// UpdateContrailConfigDatabaseNode handles update request
func (service *PluginService) UpdateContrailConfigDatabaseNode(ctx context.Context, request *services.UpdateContrailConfigDatabaseNodeRequest) (*services.UpdateContrailConfigDatabaseNodeResponse, error) {
	log.Println(" UpdateContrailConfigDatabaseNode ENTERED")

	obj := request.GetContrailConfigDatabaseNode()

	intentObj := ContrailConfigDatabaseNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailConfigDatabaseNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailConfigDatabaseNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailConfigDatabaseNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailConfigDatabaseNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailConfigDatabaseNode(ctx, request)
}

// DeleteContrailConfigDatabaseNode handles delete request
func (service *PluginService) DeleteContrailConfigDatabaseNode(ctx context.Context, request *services.DeleteContrailConfigDatabaseNodeRequest) (*services.DeleteContrailConfigDatabaseNodeResponse, error) {
	log.Println(" DeleteContrailConfigDatabaseNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailConfigDatabaseNodeIntent {
	//ContrailConfigDatabaseNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailConfigDatabaseNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailConfigDatabaseNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailConfigDatabaseNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailConfigDatabaseNode(ctx, request)
}

// GetContrailConfigDatabaseNode handles get request
func (service *PluginService) GetContrailConfigDatabaseNode(ctx context.Context, request *services.GetContrailConfigDatabaseNodeRequest) (*services.GetContrailConfigDatabaseNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailConfigDatabaseNode")
	if !ok {
		return nil, errors.New("ContrailConfigDatabaseNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailConfigDatabaseNode get failed ")
	}

	response := &services.GetContrailConfigDatabaseNodeResponse{
		ContrailConfigDatabaseNode: obj.(*models.ContrailConfigDatabaseNode),
	}
	return response, nil
}
