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

// FabricIntent
//   A struct to store attributes related to Fabric
//   needed by Intent Compiler
type FabricIntent struct {
	Uuid string
}

// EvaluateFabric - evaluates the Fabric
func EvaluateFabric(obj interface{}) {
	resourceObj := obj.(FabricIntent)
	log.Println("EvaluateFabric Called ", resourceObj)
}

// CreateFabric handles create request
func (service *PluginService) CreateFabric(ctx context.Context, request *services.CreateFabricRequest) (*services.CreateFabricResponse, error) {
	log.Println(" CreateFabric Entered")

	obj := request.GetFabric()

	intentObj := FabricIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FabricIntent"); !ok {
		compilationif.ObjsCache.Store("FabricIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("FabricIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateFabric", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Fabric")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateFabric(ctx, request)
}

// UpdateFabric handles update request
func (service *PluginService) UpdateFabric(ctx context.Context, request *services.UpdateFabricRequest) (*services.UpdateFabricResponse, error) {
	log.Println(" UpdateFabric ENTERED")

	obj := request.GetFabric()

	intentObj := FabricIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FabricIntent"); !ok {
		compilationif.ObjsCache.Store("FabricIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Fabric")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateFabric", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateFabric(ctx, request)
}

// DeleteFabric handles delete request
func (service *PluginService) DeleteFabric(ctx context.Context, request *services.DeleteFabricRequest) (*services.DeleteFabricResponse, error) {
	log.Println(" DeleteFabric ENTERED")

	objUUID := request.GetID()

	//intentObj := FabricIntent {
	//Fabric: *obj,
	//}

	//EvaluateDependencies(intentObj, "Fabric")

	objMap, ok := compilationif.ObjsCache.Load("FabricIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteFabric", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteFabric(ctx, request)
}

// GetFabric handles get request
func (service *PluginService) GetFabric(ctx context.Context, request *services.GetFabricRequest) (*services.GetFabricResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Fabric")
	if !ok {
		return nil, errors.New("Fabric get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Fabric get failed ")
	}

	response := &services.GetFabricResponse{
		Fabric: obj.(*models.Fabric),
	}
	return response, nil
}
