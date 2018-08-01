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

// RoutingPolicyIntent
//   A struct to store attributes related to RoutingPolicy
//   needed by Intent Compiler
type RoutingPolicyIntent struct {
	Uuid string
}

// EvaluateRoutingPolicy - evaluates the RoutingPolicy
func EvaluateRoutingPolicy(obj interface{}) {
	resourceObj := obj.(RoutingPolicyIntent)
	log.Println("EvaluateRoutingPolicy Called ", resourceObj)
}

// CreateRoutingPolicy handles create request
func (service *PluginService) CreateRoutingPolicy(ctx context.Context, request *services.CreateRoutingPolicyRequest) (*services.CreateRoutingPolicyResponse, error) {
	log.Println(" CreateRoutingPolicy Entered")

	obj := request.GetRoutingPolicy()

	intentObj := RoutingPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RoutingPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("RoutingPolicyIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("RoutingPolicyIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateRoutingPolicy", objMap.(*sync.Map))

	EvaluateDependencies(obj, "RoutingPolicy")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateRoutingPolicy(ctx, request)
}

// UpdateRoutingPolicy handles update request
func (service *PluginService) UpdateRoutingPolicy(ctx context.Context, request *services.UpdateRoutingPolicyRequest) (*services.UpdateRoutingPolicyResponse, error) {
	log.Println(" UpdateRoutingPolicy ENTERED")

	obj := request.GetRoutingPolicy()

	intentObj := RoutingPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("RoutingPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("RoutingPolicyIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "RoutingPolicy")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateRoutingPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateRoutingPolicy(ctx, request)
}

// DeleteRoutingPolicy handles delete request
func (service *PluginService) DeleteRoutingPolicy(ctx context.Context, request *services.DeleteRoutingPolicyRequest) (*services.DeleteRoutingPolicyResponse, error) {
	log.Println(" DeleteRoutingPolicy ENTERED")

	objUUID := request.GetID()

	//intentObj := RoutingPolicyIntent {
	//RoutingPolicy: *obj,
	//}

	//EvaluateDependencies(intentObj, "RoutingPolicy")

	objMap, ok := compilationif.ObjsCache.Load("RoutingPolicyIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteRoutingPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteRoutingPolicy(ctx, request)
}

// GetRoutingPolicy handles get request
func (service *PluginService) GetRoutingPolicy(ctx context.Context, request *services.GetRoutingPolicyRequest) (*services.GetRoutingPolicyResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("RoutingPolicy")
	if !ok {
		return nil, errors.New("RoutingPolicy get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("RoutingPolicy get failed ")
	}

	response := &services.GetRoutingPolicyResponse{
		RoutingPolicy: obj.(*models.RoutingPolicy),
	}
	return response, nil
}
