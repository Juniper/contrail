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

// ApplicationPolicySetIntent
//   A struct to store attributes related to ApplicationPolicySet
//   needed by Intent Compiler
type ApplicationPolicySetIntent struct {
	Uuid string
}

// EvaluateApplicationPolicySet - evaluates the ApplicationPolicySet
func EvaluateApplicationPolicySet(obj interface{}) {
	resourceObj := obj.(ApplicationPolicySetIntent)
	log.Println("EvaluateApplicationPolicySet Called ", resourceObj)
}

// CreateApplicationPolicySet handles create request
func (service *PluginService) CreateApplicationPolicySet(ctx context.Context, request *services.CreateApplicationPolicySetRequest) (*services.CreateApplicationPolicySetResponse, error) {
	log.Println(" CreateApplicationPolicySet Entered")

	obj := request.GetApplicationPolicySet()

	intentObj := ApplicationPolicySetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ApplicationPolicySetIntent"); !ok {
		compilationif.ObjsCache.Store("ApplicationPolicySetIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ApplicationPolicySetIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateApplicationPolicySet", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ApplicationPolicySet")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateApplicationPolicySet(ctx, request)
}

// UpdateApplicationPolicySet handles update request
func (service *PluginService) UpdateApplicationPolicySet(ctx context.Context, request *services.UpdateApplicationPolicySetRequest) (*services.UpdateApplicationPolicySetResponse, error) {
	log.Println(" UpdateApplicationPolicySet ENTERED")

	obj := request.GetApplicationPolicySet()

	intentObj := ApplicationPolicySetIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ApplicationPolicySetIntent"); !ok {
		compilationif.ObjsCache.Store("ApplicationPolicySetIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ApplicationPolicySet")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateApplicationPolicySet", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateApplicationPolicySet(ctx, request)
}

// DeleteApplicationPolicySet handles delete request
func (service *PluginService) DeleteApplicationPolicySet(ctx context.Context, request *services.DeleteApplicationPolicySetRequest) (*services.DeleteApplicationPolicySetResponse, error) {
	log.Println(" DeleteApplicationPolicySet ENTERED")

	objUUID := request.GetID()

	//intentObj := ApplicationPolicySetIntent {
	//ApplicationPolicySet: *obj,
	//}

	//EvaluateDependencies(intentObj, "ApplicationPolicySet")

	objMap, ok := compilationif.ObjsCache.Load("ApplicationPolicySetIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteApplicationPolicySet", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteApplicationPolicySet(ctx, request)
}

// GetApplicationPolicySet handles get request
func (service *PluginService) GetApplicationPolicySet(ctx context.Context, request *services.GetApplicationPolicySetRequest) (*services.GetApplicationPolicySetResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ApplicationPolicySet")
	if !ok {
		return nil, errors.New("ApplicationPolicySet get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ApplicationPolicySet get failed ")
	}

	response := &services.GetApplicationPolicySetResponse{
		ApplicationPolicySet: obj.(*models.ApplicationPolicySet),
	}
	return response, nil
}
