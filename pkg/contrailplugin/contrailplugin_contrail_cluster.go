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

// ContrailClusterIntent
//   A struct to store attributes related to ContrailCluster
//   needed by Intent Compiler
type ContrailClusterIntent struct {
	Uuid string
}

// EvaluateContrailCluster - evaluates the ContrailCluster
func EvaluateContrailCluster(obj interface{}) {
	resourceObj := obj.(ContrailClusterIntent)
	log.Println("EvaluateContrailCluster Called ", resourceObj)
}

// CreateContrailCluster handles create request
func (service *PluginService) CreateContrailCluster(ctx context.Context, request *services.CreateContrailClusterRequest) (*services.CreateContrailClusterResponse, error) {
	log.Println(" CreateContrailCluster Entered")

	obj := request.GetContrailCluster()

	intentObj := ContrailClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailClusterIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailClusterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailClusterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailCluster", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailCluster")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailCluster(ctx, request)
}

// UpdateContrailCluster handles update request
func (service *PluginService) UpdateContrailCluster(ctx context.Context, request *services.UpdateContrailClusterRequest) (*services.UpdateContrailClusterResponse, error) {
	log.Println(" UpdateContrailCluster ENTERED")

	obj := request.GetContrailCluster()

	intentObj := ContrailClusterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailClusterIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailClusterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailCluster")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailCluster(ctx, request)
}

// DeleteContrailCluster handles delete request
func (service *PluginService) DeleteContrailCluster(ctx context.Context, request *services.DeleteContrailClusterRequest) (*services.DeleteContrailClusterResponse, error) {
	log.Println(" DeleteContrailCluster ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailClusterIntent {
	//ContrailCluster: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailCluster")

	objMap, ok := compilationif.ObjsCache.Load("ContrailClusterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailCluster", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailCluster(ctx, request)
}

// GetContrailCluster handles get request
func (service *PluginService) GetContrailCluster(ctx context.Context, request *services.GetContrailClusterRequest) (*services.GetContrailClusterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailCluster")
	if !ok {
		return nil, errors.New("ContrailCluster get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailCluster get failed ")
	}

	response := &services.GetContrailClusterResponse{
		ContrailCluster: obj.(*models.ContrailCluster),
	}
	return response, nil
}
