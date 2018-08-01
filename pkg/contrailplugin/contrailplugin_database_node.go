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

// DatabaseNodeIntent
//   A struct to store attributes related to DatabaseNode
//   needed by Intent Compiler
type DatabaseNodeIntent struct {
	Uuid string
}

// EvaluateDatabaseNode - evaluates the DatabaseNode
func EvaluateDatabaseNode(obj interface{}) {
	resourceObj := obj.(DatabaseNodeIntent)
	log.Println("EvaluateDatabaseNode Called ", resourceObj)
}

// CreateDatabaseNode handles create request
func (service *PluginService) CreateDatabaseNode(ctx context.Context, request *services.CreateDatabaseNodeRequest) (*services.CreateDatabaseNodeResponse, error) {
	log.Println(" CreateDatabaseNode Entered")

	obj := request.GetDatabaseNode()

	intentObj := DatabaseNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DatabaseNodeIntent"); !ok {
		compilationif.ObjsCache.Store("DatabaseNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("DatabaseNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateDatabaseNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "DatabaseNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateDatabaseNode(ctx, request)
}

// UpdateDatabaseNode handles update request
func (service *PluginService) UpdateDatabaseNode(ctx context.Context, request *services.UpdateDatabaseNodeRequest) (*services.UpdateDatabaseNodeResponse, error) {
	log.Println(" UpdateDatabaseNode ENTERED")

	obj := request.GetDatabaseNode()

	intentObj := DatabaseNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DatabaseNodeIntent"); !ok {
		compilationif.ObjsCache.Store("DatabaseNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "DatabaseNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateDatabaseNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateDatabaseNode(ctx, request)
}

// DeleteDatabaseNode handles delete request
func (service *PluginService) DeleteDatabaseNode(ctx context.Context, request *services.DeleteDatabaseNodeRequest) (*services.DeleteDatabaseNodeResponse, error) {
	log.Println(" DeleteDatabaseNode ENTERED")

	objUUID := request.GetID()

	//intentObj := DatabaseNodeIntent {
	//DatabaseNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "DatabaseNode")

	objMap, ok := compilationif.ObjsCache.Load("DatabaseNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteDatabaseNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteDatabaseNode(ctx, request)
}

// GetDatabaseNode handles get request
func (service *PluginService) GetDatabaseNode(ctx context.Context, request *services.GetDatabaseNodeRequest) (*services.GetDatabaseNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("DatabaseNode")
	if !ok {
		return nil, errors.New("DatabaseNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("DatabaseNode get failed ")
	}

	response := &services.GetDatabaseNodeResponse{
		DatabaseNode: obj.(*models.DatabaseNode),
	}
	return response, nil
}
