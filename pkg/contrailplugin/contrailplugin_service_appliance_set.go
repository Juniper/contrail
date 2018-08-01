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

// ServiceApplianceSetIntent
//   A struct to store attributes related to ServiceApplianceSet
//   needed by Intent Compiler
type ServiceApplianceSetIntent struct {
	Uuid string
}

// EvaluateServiceApplianceSet - evaluates the ServiceApplianceSet
func EvaluateServiceApplianceSet(obj interface{}) {
	resourceObj := obj.(ServiceApplianceSetIntent)
	log.Println("EvaluateServiceApplianceSet Called ", resourceObj)
}

// CreateServiceApplianceSet handles create request
func (service *PluginService) CreateServiceApplianceSet(ctx context.Context, request *services.CreateServiceApplianceSetRequest) (*services.CreateServiceApplianceSetResponse, error) {
	log.Println(" CreateServiceApplianceSet Entered")

	obj := request.GetServiceApplianceSet()

	intentObj := ServiceApplianceSetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceApplianceSetIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceApplianceSetIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceApplianceSetIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceApplianceSet", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceApplianceSet")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceApplianceSet(ctx, request)
}

// UpdateServiceApplianceSet handles update request
func (service *PluginService) UpdateServiceApplianceSet(ctx context.Context, request *services.UpdateServiceApplianceSetRequest) (*services.UpdateServiceApplianceSetResponse, error) {
	log.Println(" UpdateServiceApplianceSet ENTERED")

	obj := request.GetServiceApplianceSet()

	intentObj := ServiceApplianceSetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceApplianceSetIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceApplianceSetIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceApplianceSet")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceApplianceSet", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceApplianceSet(ctx, request)
}

// DeleteServiceApplianceSet handles delete request
func (service *PluginService) DeleteServiceApplianceSet(ctx context.Context, request *services.DeleteServiceApplianceSetRequest) (*services.DeleteServiceApplianceSetResponse, error) {
	log.Println(" DeleteServiceApplianceSet ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceApplianceSetIntent {
	//ServiceApplianceSet: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceApplianceSet")

	objMap, ok := compilationif.ObjsCache.Load("ServiceApplianceSetIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceApplianceSet", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceApplianceSet(ctx, request)
}

// GetServiceApplianceSet handles get request
func (service *PluginService) GetServiceApplianceSet(ctx context.Context, request *services.GetServiceApplianceSetRequest) (*services.GetServiceApplianceSetResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceApplianceSet")
	if !ok {
		return nil, errors.New("ServiceApplianceSet get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceApplianceSet get failed ")
	}

	response := &services.GetServiceApplianceSetResponse{
		ServiceApplianceSet: obj.(*models.ServiceApplianceSet),
	}
	return response, nil
}
