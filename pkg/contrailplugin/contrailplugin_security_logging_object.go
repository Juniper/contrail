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

// SecurityLoggingObjectIntent
//   A struct to store attributes related to SecurityLoggingObject
//   needed by Intent Compiler
type SecurityLoggingObjectIntent struct {
	Uuid string
}

// EvaluateSecurityLoggingObject - evaluates the SecurityLoggingObject
func EvaluateSecurityLoggingObject(obj interface{}) {
	resourceObj := obj.(SecurityLoggingObjectIntent)
	log.Println("EvaluateSecurityLoggingObject Called ", resourceObj)
}

// CreateSecurityLoggingObject handles create request
func (service *PluginService) CreateSecurityLoggingObject(ctx context.Context, request *services.CreateSecurityLoggingObjectRequest) (*services.CreateSecurityLoggingObjectResponse, error) {
	log.Println(" CreateSecurityLoggingObject Entered")

	obj := request.GetSecurityLoggingObject()

	intentObj := SecurityLoggingObjectIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SecurityLoggingObjectIntent"); !ok {
		compilationif.ObjsCache.Store("SecurityLoggingObjectIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("SecurityLoggingObjectIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateSecurityLoggingObject", objMap.(*sync.Map))

	EvaluateDependencies(obj, "SecurityLoggingObject")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateSecurityLoggingObject(ctx, request)
}

// UpdateSecurityLoggingObject handles update request
func (service *PluginService) UpdateSecurityLoggingObject(ctx context.Context, request *services.UpdateSecurityLoggingObjectRequest) (*services.UpdateSecurityLoggingObjectResponse, error) {
	log.Println(" UpdateSecurityLoggingObject ENTERED")

	obj := request.GetSecurityLoggingObject()

	intentObj := SecurityLoggingObjectIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("SecurityLoggingObjectIntent"); !ok {
		compilationif.ObjsCache.Store("SecurityLoggingObjectIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "SecurityLoggingObject")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateSecurityLoggingObject", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateSecurityLoggingObject(ctx, request)
}

// DeleteSecurityLoggingObject handles delete request
func (service *PluginService) DeleteSecurityLoggingObject(ctx context.Context, request *services.DeleteSecurityLoggingObjectRequest) (*services.DeleteSecurityLoggingObjectResponse, error) {
	log.Println(" DeleteSecurityLoggingObject ENTERED")

	objUUID := request.GetID()

	//intentObj := SecurityLoggingObjectIntent {
	//SecurityLoggingObject: *obj,
	//}

	//EvaluateDependencies(intentObj, "SecurityLoggingObject")

	objMap, ok := compilationif.ObjsCache.Load("SecurityLoggingObjectIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteSecurityLoggingObject", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteSecurityLoggingObject(ctx, request)
}

// GetSecurityLoggingObject handles get request
func (service *PluginService) GetSecurityLoggingObject(ctx context.Context, request *services.GetSecurityLoggingObjectRequest) (*services.GetSecurityLoggingObjectResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("SecurityLoggingObject")
	if !ok {
		return nil, errors.New("SecurityLoggingObject get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("SecurityLoggingObject get failed ")
	}

	response := &services.GetSecurityLoggingObjectResponse{
		SecurityLoggingObject: obj.(*models.SecurityLoggingObject),
	}
	return response, nil
}
