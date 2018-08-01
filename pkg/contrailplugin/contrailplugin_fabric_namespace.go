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

// FabricNamespaceIntent
//   A struct to store attributes related to FabricNamespace
//   needed by Intent Compiler
type FabricNamespaceIntent struct {
	Uuid string
}

// EvaluateFabricNamespace - evaluates the FabricNamespace
func EvaluateFabricNamespace(obj interface{}) {
	resourceObj := obj.(FabricNamespaceIntent)
	log.Println("EvaluateFabricNamespace Called ", resourceObj)
}

// CreateFabricNamespace handles create request
func (service *PluginService) CreateFabricNamespace(ctx context.Context, request *services.CreateFabricNamespaceRequest) (*services.CreateFabricNamespaceResponse, error) {
	log.Println(" CreateFabricNamespace Entered")

	obj := request.GetFabricNamespace()

	intentObj := FabricNamespaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FabricNamespaceIntent"); !ok {
		compilationif.ObjsCache.Store("FabricNamespaceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("FabricNamespaceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateFabricNamespace", objMap.(*sync.Map))

	EvaluateDependencies(obj, "FabricNamespace")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateFabricNamespace(ctx, request)
}

// UpdateFabricNamespace handles update request
func (service *PluginService) UpdateFabricNamespace(ctx context.Context, request *services.UpdateFabricNamespaceRequest) (*services.UpdateFabricNamespaceResponse, error) {
	log.Println(" UpdateFabricNamespace ENTERED")

	obj := request.GetFabricNamespace()

	intentObj := FabricNamespaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FabricNamespaceIntent"); !ok {
		compilationif.ObjsCache.Store("FabricNamespaceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "FabricNamespace")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateFabricNamespace", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateFabricNamespace(ctx, request)
}

// DeleteFabricNamespace handles delete request
func (service *PluginService) DeleteFabricNamespace(ctx context.Context, request *services.DeleteFabricNamespaceRequest) (*services.DeleteFabricNamespaceResponse, error) {
	log.Println(" DeleteFabricNamespace ENTERED")

	objUUID := request.GetID()

	//intentObj := FabricNamespaceIntent {
	//FabricNamespace: *obj,
	//}

	//EvaluateDependencies(intentObj, "FabricNamespace")

	objMap, ok := compilationif.ObjsCache.Load("FabricNamespaceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteFabricNamespace", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteFabricNamespace(ctx, request)
}

// GetFabricNamespace handles get request
func (service *PluginService) GetFabricNamespace(ctx context.Context, request *services.GetFabricNamespaceRequest) (*services.GetFabricNamespaceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("FabricNamespace")
	if !ok {
		return nil, errors.New("FabricNamespace get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("FabricNamespace get failed ")
	}

	response := &services.GetFabricNamespaceResponse{
		FabricNamespace: obj.(*models.FabricNamespace),
	}
	return response, nil
}
