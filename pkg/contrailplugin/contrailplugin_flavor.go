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

// FlavorIntent
//   A struct to store attributes related to Flavor
//   needed by Intent Compiler
type FlavorIntent struct {
	Uuid string
}

// EvaluateFlavor - evaluates the Flavor
func EvaluateFlavor(obj interface{}) {
	resourceObj := obj.(FlavorIntent)
	log.Println("EvaluateFlavor Called ", resourceObj)
}

// CreateFlavor handles create request
func (service *PluginService) CreateFlavor(ctx context.Context, request *services.CreateFlavorRequest) (*services.CreateFlavorResponse, error) {
	log.Println(" CreateFlavor Entered")

	obj := request.GetFlavor()

	intentObj := FlavorIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FlavorIntent"); !ok {
		compilationif.ObjsCache.Store("FlavorIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("FlavorIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateFlavor", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Flavor")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateFlavor(ctx, request)
}

// UpdateFlavor handles update request
func (service *PluginService) UpdateFlavor(ctx context.Context, request *services.UpdateFlavorRequest) (*services.UpdateFlavorResponse, error) {
	log.Println(" UpdateFlavor ENTERED")

	obj := request.GetFlavor()

	intentObj := FlavorIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FlavorIntent"); !ok {
		compilationif.ObjsCache.Store("FlavorIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Flavor")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateFlavor", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateFlavor(ctx, request)
}

// DeleteFlavor handles delete request
func (service *PluginService) DeleteFlavor(ctx context.Context, request *services.DeleteFlavorRequest) (*services.DeleteFlavorResponse, error) {
	log.Println(" DeleteFlavor ENTERED")

	objUUID := request.GetID()

	//intentObj := FlavorIntent {
	//Flavor: *obj,
	//}

	//EvaluateDependencies(intentObj, "Flavor")

	objMap, ok := compilationif.ObjsCache.Load("FlavorIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteFlavor", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteFlavor(ctx, request)
}

// GetFlavor handles get request
func (service *PluginService) GetFlavor(ctx context.Context, request *services.GetFlavorRequest) (*services.GetFlavorResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Flavor")
	if !ok {
		return nil, errors.New("Flavor get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Flavor get failed ")
	}

	response := &services.GetFlavorResponse{
		Flavor: obj.(*models.Flavor),
	}
	return response, nil
}
