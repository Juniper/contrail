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

// ServiceObjectIntent
//   A struct to store attributes related to ServiceObject
//   needed by Intent Compiler
type ServiceObjectIntent struct {
	Uuid string
}

// EvaluateServiceObject - evaluates the ServiceObject
func EvaluateServiceObject(obj interface{}) {
	resourceObj := obj.(ServiceObjectIntent)
	log.Println("EvaluateServiceObject Called ", resourceObj)
}

// CreateServiceObject handles create request
func (service *PluginService) CreateServiceObject(ctx context.Context, request *services.CreateServiceObjectRequest) (*services.CreateServiceObjectResponse, error) {
	log.Println(" CreateServiceObject Entered")

	obj := request.GetServiceObject()

	intentObj := ServiceObjectIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceObjectIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceObjectIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceObjectIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceObject", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceObject")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceObject(ctx, request)
}

// UpdateServiceObject handles update request
func (service *PluginService) UpdateServiceObject(ctx context.Context, request *services.UpdateServiceObjectRequest) (*services.UpdateServiceObjectResponse, error) {
	log.Println(" UpdateServiceObject ENTERED")

	obj := request.GetServiceObject()

	intentObj := ServiceObjectIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceObjectIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceObjectIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceObject")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceObject", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceObject(ctx, request)
}

// DeleteServiceObject handles delete request
func (service *PluginService) DeleteServiceObject(ctx context.Context, request *services.DeleteServiceObjectRequest) (*services.DeleteServiceObjectResponse, error) {
	log.Println(" DeleteServiceObject ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceObjectIntent {
	//ServiceObject: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceObject")

	objMap, ok := compilationif.ObjsCache.Load("ServiceObjectIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceObject", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceObject(ctx, request)
}

// GetServiceObject handles get request
func (service *PluginService) GetServiceObject(ctx context.Context, request *services.GetServiceObjectRequest) (*services.GetServiceObjectResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceObject")
	if !ok {
		return nil, errors.New("ServiceObject get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceObject get failed ")
	}

	response := &services.GetServiceObjectResponse{
		ServiceObject: obj.(*models.ServiceObject),
	}
	return response, nil
}
