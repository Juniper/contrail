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

// DiscoveryServiceAssignmentIntent
//   A struct to store attributes related to DiscoveryServiceAssignment
//   needed by Intent Compiler
type DiscoveryServiceAssignmentIntent struct {
	Uuid string
}

// EvaluateDiscoveryServiceAssignment - evaluates the DiscoveryServiceAssignment
func EvaluateDiscoveryServiceAssignment(obj interface{}) {
	resourceObj := obj.(DiscoveryServiceAssignmentIntent)
	log.Println("EvaluateDiscoveryServiceAssignment Called ", resourceObj)
}

// CreateDiscoveryServiceAssignment handles create request
func (service *PluginService) CreateDiscoveryServiceAssignment(ctx context.Context, request *services.CreateDiscoveryServiceAssignmentRequest) (*services.CreateDiscoveryServiceAssignmentResponse, error) {
	log.Println(" CreateDiscoveryServiceAssignment Entered")

	obj := request.GetDiscoveryServiceAssignment()

	intentObj := DiscoveryServiceAssignmentIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DiscoveryServiceAssignmentIntent"); !ok {
		compilationif.ObjsCache.Store("DiscoveryServiceAssignmentIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("DiscoveryServiceAssignmentIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateDiscoveryServiceAssignment", objMap.(*sync.Map))

	EvaluateDependencies(obj, "DiscoveryServiceAssignment")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateDiscoveryServiceAssignment(ctx, request)
}

// UpdateDiscoveryServiceAssignment handles update request
func (service *PluginService) UpdateDiscoveryServiceAssignment(ctx context.Context, request *services.UpdateDiscoveryServiceAssignmentRequest) (*services.UpdateDiscoveryServiceAssignmentResponse, error) {
	log.Println(" UpdateDiscoveryServiceAssignment ENTERED")

	obj := request.GetDiscoveryServiceAssignment()

	intentObj := DiscoveryServiceAssignmentIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DiscoveryServiceAssignmentIntent"); !ok {
		compilationif.ObjsCache.Store("DiscoveryServiceAssignmentIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "DiscoveryServiceAssignment")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateDiscoveryServiceAssignment", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateDiscoveryServiceAssignment(ctx, request)
}

// DeleteDiscoveryServiceAssignment handles delete request
func (service *PluginService) DeleteDiscoveryServiceAssignment(ctx context.Context, request *services.DeleteDiscoveryServiceAssignmentRequest) (*services.DeleteDiscoveryServiceAssignmentResponse, error) {
	log.Println(" DeleteDiscoveryServiceAssignment ENTERED")

	objUUID := request.GetID()

	//intentObj := DiscoveryServiceAssignmentIntent {
	//DiscoveryServiceAssignment: *obj,
	//}

	//EvaluateDependencies(intentObj, "DiscoveryServiceAssignment")

	objMap, ok := compilationif.ObjsCache.Load("DiscoveryServiceAssignmentIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteDiscoveryServiceAssignment", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteDiscoveryServiceAssignment(ctx, request)
}

// GetDiscoveryServiceAssignment handles get request
func (service *PluginService) GetDiscoveryServiceAssignment(ctx context.Context, request *services.GetDiscoveryServiceAssignmentRequest) (*services.GetDiscoveryServiceAssignmentResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("DiscoveryServiceAssignment")
	if !ok {
		return nil, errors.New("DiscoveryServiceAssignment get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("DiscoveryServiceAssignment get failed ")
	}

	response := &services.GetDiscoveryServiceAssignmentResponse{
		DiscoveryServiceAssignment: obj.(*models.DiscoveryServiceAssignment),
	}
	return response, nil
}
