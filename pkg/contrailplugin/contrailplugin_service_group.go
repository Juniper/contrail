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

// ServiceGroupIntent
//   A struct to store attributes related to ServiceGroup
//   needed by Intent Compiler
type ServiceGroupIntent struct {
	Uuid string
}

// EvaluateServiceGroup - evaluates the ServiceGroup
func EvaluateServiceGroup(obj interface{}) {
	resourceObj := obj.(ServiceGroupIntent)
	log.Println("EvaluateServiceGroup Called ", resourceObj)
}

// CreateServiceGroup handles create request
func (service *PluginService) CreateServiceGroup(ctx context.Context, request *services.CreateServiceGroupRequest) (*services.CreateServiceGroupResponse, error) {
	log.Println(" CreateServiceGroup Entered")

	obj := request.GetServiceGroup()

	intentObj := ServiceGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceGroupIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceGroupIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceGroup", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceGroup")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceGroup(ctx, request)
}

// UpdateServiceGroup handles update request
func (service *PluginService) UpdateServiceGroup(ctx context.Context, request *services.UpdateServiceGroupRequest) (*services.UpdateServiceGroupResponse, error) {
	log.Println(" UpdateServiceGroup ENTERED")

	obj := request.GetServiceGroup()

	intentObj := ServiceGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceGroupIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceGroupIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceGroup")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceGroup(ctx, request)
}

// DeleteServiceGroup handles delete request
func (service *PluginService) DeleteServiceGroup(ctx context.Context, request *services.DeleteServiceGroupRequest) (*services.DeleteServiceGroupResponse, error) {
	log.Println(" DeleteServiceGroup ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceGroupIntent {
	//ServiceGroup: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceGroup")

	objMap, ok := compilationif.ObjsCache.Load("ServiceGroupIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceGroup(ctx, request)
}

// GetServiceGroup handles get request
func (service *PluginService) GetServiceGroup(ctx context.Context, request *services.GetServiceGroupRequest) (*services.GetServiceGroupResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceGroup")
	if !ok {
		return nil, errors.New("ServiceGroup get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceGroup get failed ")
	}

	response := &services.GetServiceGroupResponse{
		ServiceGroup: obj.(*models.ServiceGroup),
	}
	return response, nil
}
