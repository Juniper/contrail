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

// NetworkPolicyIntent
//   A struct to store attributes related to NetworkPolicy
//   needed by Intent Compiler
type NetworkPolicyIntent struct {
	Uuid string
}

// EvaluateNetworkPolicy - evaluates the NetworkPolicy
func EvaluateNetworkPolicy(obj interface{}) {
	resourceObj := obj.(NetworkPolicyIntent)
	log.Println("EvaluateNetworkPolicy Called ", resourceObj)
}

// CreateNetworkPolicy handles create request
func (service *PluginService) CreateNetworkPolicy(ctx context.Context, request *services.CreateNetworkPolicyRequest) (*services.CreateNetworkPolicyResponse, error) {
	log.Println(" CreateNetworkPolicy Entered")

	obj := request.GetNetworkPolicy()

	intentObj := NetworkPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NetworkPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("NetworkPolicyIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("NetworkPolicyIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateNetworkPolicy", objMap.(*sync.Map))

	EvaluateDependencies(obj, "NetworkPolicy")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateNetworkPolicy(ctx, request)
}

// UpdateNetworkPolicy handles update request
func (service *PluginService) UpdateNetworkPolicy(ctx context.Context, request *services.UpdateNetworkPolicyRequest) (*services.UpdateNetworkPolicyResponse, error) {
	log.Println(" UpdateNetworkPolicy ENTERED")

	obj := request.GetNetworkPolicy()

	intentObj := NetworkPolicyIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NetworkPolicyIntent"); !ok {
		compilationif.ObjsCache.Store("NetworkPolicyIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "NetworkPolicy")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateNetworkPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateNetworkPolicy(ctx, request)
}

// DeleteNetworkPolicy handles delete request
func (service *PluginService) DeleteNetworkPolicy(ctx context.Context, request *services.DeleteNetworkPolicyRequest) (*services.DeleteNetworkPolicyResponse, error) {
	log.Println(" DeleteNetworkPolicy ENTERED")

	objUUID := request.GetID()

	//intentObj := NetworkPolicyIntent {
	//NetworkPolicy: *obj,
	//}

	//EvaluateDependencies(intentObj, "NetworkPolicy")

	objMap, ok := compilationif.ObjsCache.Load("NetworkPolicyIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteNetworkPolicy", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteNetworkPolicy(ctx, request)
}

// GetNetworkPolicy handles get request
func (service *PluginService) GetNetworkPolicy(ctx context.Context, request *services.GetNetworkPolicyRequest) (*services.GetNetworkPolicyResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("NetworkPolicy")
	if !ok {
		return nil, errors.New("NetworkPolicy get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("NetworkPolicy get failed ")
	}

	response := &services.GetNetworkPolicyResponse{
		NetworkPolicy: obj.(*models.NetworkPolicy),
	}
	return response, nil
}
