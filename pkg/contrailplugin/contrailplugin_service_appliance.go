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

// ServiceApplianceIntent
//   A struct to store attributes related to ServiceAppliance
//   needed by Intent Compiler
type ServiceApplianceIntent struct {
	Uuid string
}

// EvaluateServiceAppliance - evaluates the ServiceAppliance
func EvaluateServiceAppliance(obj interface{}) {
	resourceObj := obj.(ServiceApplianceIntent)
	log.Println("EvaluateServiceAppliance Called ", resourceObj)
}

// CreateServiceAppliance handles create request
func (service *PluginService) CreateServiceAppliance(ctx context.Context, request *services.CreateServiceApplianceRequest) (*services.CreateServiceApplianceResponse, error) {
	log.Println(" CreateServiceAppliance Entered")

	obj := request.GetServiceAppliance()

	intentObj := ServiceApplianceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceApplianceIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceApplianceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ServiceApplianceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateServiceAppliance", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ServiceAppliance")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateServiceAppliance(ctx, request)
}

// UpdateServiceAppliance handles update request
func (service *PluginService) UpdateServiceAppliance(ctx context.Context, request *services.UpdateServiceApplianceRequest) (*services.UpdateServiceApplianceResponse, error) {
	log.Println(" UpdateServiceAppliance ENTERED")

	obj := request.GetServiceAppliance()

	intentObj := ServiceApplianceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ServiceApplianceIntent"); !ok {
		compilationif.ObjsCache.Store("ServiceApplianceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ServiceAppliance")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateServiceAppliance", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateServiceAppliance(ctx, request)
}

// DeleteServiceAppliance handles delete request
func (service *PluginService) DeleteServiceAppliance(ctx context.Context, request *services.DeleteServiceApplianceRequest) (*services.DeleteServiceApplianceResponse, error) {
	log.Println(" DeleteServiceAppliance ENTERED")

	objUUID := request.GetID()

	//intentObj := ServiceApplianceIntent {
	//ServiceAppliance: *obj,
	//}

	//EvaluateDependencies(intentObj, "ServiceAppliance")

	objMap, ok := compilationif.ObjsCache.Load("ServiceApplianceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteServiceAppliance", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteServiceAppliance(ctx, request)
}

// GetServiceAppliance handles get request
func (service *PluginService) GetServiceAppliance(ctx context.Context, request *services.GetServiceApplianceRequest) (*services.GetServiceApplianceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ServiceAppliance")
	if !ok {
		return nil, errors.New("ServiceAppliance get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ServiceAppliance get failed ")
	}

	response := &services.GetServiceApplianceResponse{
		ServiceAppliance: obj.(*models.ServiceAppliance),
	}
	return response, nil
}
