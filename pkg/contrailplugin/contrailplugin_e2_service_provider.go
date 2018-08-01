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

// E2ServiceProviderIntent
//   A struct to store attributes related to E2ServiceProvider
//   needed by Intent Compiler
type E2ServiceProviderIntent struct {
	Uuid string
}

// EvaluateE2ServiceProvider - evaluates the E2ServiceProvider
func EvaluateE2ServiceProvider(obj interface{}) {
	resourceObj := obj.(E2ServiceProviderIntent)
	log.Println("EvaluateE2ServiceProvider Called ", resourceObj)
}

// CreateE2ServiceProvider handles create request
func (service *PluginService) CreateE2ServiceProvider(ctx context.Context, request *services.CreateE2ServiceProviderRequest) (*services.CreateE2ServiceProviderResponse, error) {
	log.Println(" CreateE2ServiceProvider Entered")

	obj := request.GetE2ServiceProvider()

	intentObj := E2ServiceProviderIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("E2ServiceProviderIntent"); !ok {
		compilationif.ObjsCache.Store("E2ServiceProviderIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("E2ServiceProviderIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateE2ServiceProvider", objMap.(*sync.Map))

	EvaluateDependencies(obj, "E2ServiceProvider")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateE2ServiceProvider(ctx, request)
}

// UpdateE2ServiceProvider handles update request
func (service *PluginService) UpdateE2ServiceProvider(ctx context.Context, request *services.UpdateE2ServiceProviderRequest) (*services.UpdateE2ServiceProviderResponse, error) {
	log.Println(" UpdateE2ServiceProvider ENTERED")

	obj := request.GetE2ServiceProvider()

	intentObj := E2ServiceProviderIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("E2ServiceProviderIntent"); !ok {
		compilationif.ObjsCache.Store("E2ServiceProviderIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "E2ServiceProvider")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateE2ServiceProvider", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateE2ServiceProvider(ctx, request)
}

// DeleteE2ServiceProvider handles delete request
func (service *PluginService) DeleteE2ServiceProvider(ctx context.Context, request *services.DeleteE2ServiceProviderRequest) (*services.DeleteE2ServiceProviderResponse, error) {
	log.Println(" DeleteE2ServiceProvider ENTERED")

	objUUID := request.GetID()

	//intentObj := E2ServiceProviderIntent {
	//E2ServiceProvider: *obj,
	//}

	//EvaluateDependencies(intentObj, "E2ServiceProvider")

	objMap, ok := compilationif.ObjsCache.Load("E2ServiceProviderIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteE2ServiceProvider", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteE2ServiceProvider(ctx, request)
}

// GetE2ServiceProvider handles get request
func (service *PluginService) GetE2ServiceProvider(ctx context.Context, request *services.GetE2ServiceProviderRequest) (*services.GetE2ServiceProviderResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("E2ServiceProvider")
	if !ok {
		return nil, errors.New("E2ServiceProvider get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("E2ServiceProvider get failed ")
	}

	response := &services.GetE2ServiceProviderResponse{
		E2ServiceProvider: obj.(*models.E2ServiceProvider),
	}
	return response, nil
}
