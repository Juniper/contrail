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

// PeeringPolicyIntent
//   A struct to store attributes related to PeeringPolicy
//   needed by Intent Compiler
type PeeringPolicyIntent struct {
	Uuid string
}

// EvaluatePeeringPolicy - evaluates the PeeringPolicy
func EvaluatePeeringPolicy(obj interface{}) {
	resourceObj := obj.(PeeringPolicyIntent)
	log.Println("EvaluatePeeringPolicy Called ", resourceObj)
}

// CreatePeeringPolicy handles create request
func (service *PluginService) CreatePeeringPolicy(ctx context.Context, request *services.CreatePeeringPolicyRequest) (*services.CreatePeeringPolicyResponse, error) {
	log.Println(" CreatePeeringPolicy Entered")

	obj := request.GetPeeringPolicy()

	intentObj := PeeringPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PeeringPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("PeeringPolicyIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("PeeringPolicyIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreatePeeringPolicy", objMap.(*sync.Map))

	EvaluateDependencies(obj, "PeeringPolicy")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreatePeeringPolicy(ctx, request)
}

// UpdatePeeringPolicy handles update request
func (service *PluginService) UpdatePeeringPolicy(ctx context.Context, request *services.UpdatePeeringPolicyRequest) (*services.UpdatePeeringPolicyResponse, error) {
	log.Println(" UpdatePeeringPolicy ENTERED")

	obj := request.GetPeeringPolicy()

	intentObj := PeeringPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PeeringPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("PeeringPolicyIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "PeeringPolicy")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdatePeeringPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdatePeeringPolicy(ctx, request)
}

// DeletePeeringPolicy handles delete request
func (service *PluginService) DeletePeeringPolicy(ctx context.Context, request *services.DeletePeeringPolicyRequest) (*services.DeletePeeringPolicyResponse, error) {
	log.Println(" DeletePeeringPolicy ENTERED")

	objUUID := request.GetID()

	//intentObj := PeeringPolicyIntent {
	//PeeringPolicy: *obj,
	//}

	//EvaluateDependencies(intentObj, "PeeringPolicy")

	objMap, ok := compilationif.ObjsCache.Load("PeeringPolicyIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeletePeeringPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeletePeeringPolicy(ctx, request)
}

// GetPeeringPolicy handles get request
func (service *PluginService) GetPeeringPolicy(ctx context.Context, request *services.GetPeeringPolicyRequest) (*services.GetPeeringPolicyResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("PeeringPolicy")
	if !ok {
		return nil, errors.New("PeeringPolicy get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("PeeringPolicy get failed ")
	}

	response := &services.GetPeeringPolicyResponse{
		PeeringPolicy: obj.(*models.PeeringPolicy),
	}
	return response, nil
}
