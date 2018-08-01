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

// LocationIntent
//   A struct to store attributes related to Location
//   needed by Intent Compiler
type LocationIntent struct {
	Uuid string
}

// EvaluateLocation - evaluates the Location
func EvaluateLocation(obj interface{}) {
	resourceObj := obj.(LocationIntent)
	log.Println("EvaluateLocation Called ", resourceObj)
}

// CreateLocation handles create request
func (service *PluginService) CreateLocation(ctx context.Context, request *services.CreateLocationRequest) (*services.CreateLocationResponse, error) {
	log.Println(" CreateLocation Entered")

	obj := request.GetLocation()

	intentObj := LocationIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LocationIntent"); !ok {
		compilationif.ObjsCache.Store("LocationIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LocationIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLocation", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Location")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLocation(ctx, request)
}

// UpdateLocation handles update request
func (service *PluginService) UpdateLocation(ctx context.Context, request *services.UpdateLocationRequest) (*services.UpdateLocationResponse, error) {
	log.Println(" UpdateLocation ENTERED")

	obj := request.GetLocation()

	intentObj := LocationIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LocationIntent"); !ok {
		compilationif.ObjsCache.Store("LocationIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Location")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLocation", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLocation(ctx, request)
}

// DeleteLocation handles delete request
func (service *PluginService) DeleteLocation(ctx context.Context, request *services.DeleteLocationRequest) (*services.DeleteLocationResponse, error) {
	log.Println(" DeleteLocation ENTERED")

	objUUID := request.GetID()

	//intentObj := LocationIntent {
	//Location: *obj,
	//}

	//EvaluateDependencies(intentObj, "Location")

	objMap, ok := compilationif.ObjsCache.Load("LocationIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLocation", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLocation(ctx, request)
}

// GetLocation handles get request
func (service *PluginService) GetLocation(ctx context.Context, request *services.GetLocationRequest) (*services.GetLocationResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Location")
	if !ok {
		return nil, errors.New("Location get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Location get failed ")
	}

	response := &services.GetLocationResponse{
		Location: obj.(*models.Location),
	}
	return response, nil
}
