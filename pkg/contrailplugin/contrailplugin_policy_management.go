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

// PolicyManagementIntent
//   A struct to store attributes related to PolicyManagement
//   needed by Intent Compiler
type PolicyManagementIntent struct {
	Uuid string
}

// EvaluatePolicyManagement - evaluates the PolicyManagement
func EvaluatePolicyManagement(obj interface{}) {
	resourceObj := obj.(PolicyManagementIntent)
	log.Println("EvaluatePolicyManagement Called ", resourceObj)
}

// CreatePolicyManagement handles create request
func (service *PluginService) CreatePolicyManagement(ctx context.Context, request *services.CreatePolicyManagementRequest) (*services.CreatePolicyManagementResponse, error) {
	log.Println(" CreatePolicyManagement Entered")

	obj := request.GetPolicyManagement()

	intentObj := PolicyManagementIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PolicyManagementIntent"); !ok {
		compilationif.ObjsCache.Store("PolicyManagementIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("PolicyManagementIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreatePolicyManagement", objMap.(*sync.Map))

	EvaluateDependencies(obj, "PolicyManagement")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreatePolicyManagement(ctx, request)
}

// UpdatePolicyManagement handles update request
func (service *PluginService) UpdatePolicyManagement(ctx context.Context, request *services.UpdatePolicyManagementRequest) (*services.UpdatePolicyManagementResponse, error) {
	log.Println(" UpdatePolicyManagement ENTERED")

	obj := request.GetPolicyManagement()

	intentObj := PolicyManagementIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PolicyManagementIntent"); !ok {
		compilationif.ObjsCache.Store("PolicyManagementIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "PolicyManagement")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdatePolicyManagement", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdatePolicyManagement(ctx, request)
}

// DeletePolicyManagement handles delete request
func (service *PluginService) DeletePolicyManagement(ctx context.Context, request *services.DeletePolicyManagementRequest) (*services.DeletePolicyManagementResponse, error) {
	log.Println(" DeletePolicyManagement ENTERED")

	objUUID := request.GetID()

	//intentObj := PolicyManagementIntent {
	//PolicyManagement: *obj,
	//}

	//EvaluateDependencies(intentObj, "PolicyManagement")

	objMap, ok := compilationif.ObjsCache.Load("PolicyManagementIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeletePolicyManagement", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeletePolicyManagement(ctx, request)
}

// GetPolicyManagement handles get request
func (service *PluginService) GetPolicyManagement(ctx context.Context, request *services.GetPolicyManagementRequest) (*services.GetPolicyManagementResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("PolicyManagement")
	if !ok {
		return nil, errors.New("PolicyManagement get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("PolicyManagement get failed ")
	}

	response := &services.GetPolicyManagementResponse{
		PolicyManagement: obj.(*models.PolicyManagement),
	}
	return response, nil
}
