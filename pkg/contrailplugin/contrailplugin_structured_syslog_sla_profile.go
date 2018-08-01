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

// StructuredSyslogSLAProfileIntent
//   A struct to store attributes related to StructuredSyslogSLAProfile
//   needed by Intent Compiler
type StructuredSyslogSLAProfileIntent struct {
	Uuid string
}

// EvaluateStructuredSyslogSLAProfile - evaluates the StructuredSyslogSLAProfile
func EvaluateStructuredSyslogSLAProfile(obj interface{}) {
	resourceObj := obj.(StructuredSyslogSLAProfileIntent)
	log.Println("EvaluateStructuredSyslogSLAProfile Called ", resourceObj)
}

// CreateStructuredSyslogSLAProfile handles create request
func (service *PluginService) CreateStructuredSyslogSLAProfile(ctx context.Context, request *services.CreateStructuredSyslogSLAProfileRequest) (*services.CreateStructuredSyslogSLAProfileResponse, error) {
	log.Println(" CreateStructuredSyslogSLAProfile Entered")

	obj := request.GetStructuredSyslogSLAProfile()

	intentObj := StructuredSyslogSLAProfileIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogSLAProfileIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogSLAProfileIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogSLAProfileIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateStructuredSyslogSLAProfile", objMap.(*sync.Map))

	EvaluateDependencies(obj, "StructuredSyslogSLAProfile")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateStructuredSyslogSLAProfile(ctx, request)
}

// UpdateStructuredSyslogSLAProfile handles update request
func (service *PluginService) UpdateStructuredSyslogSLAProfile(ctx context.Context, request *services.UpdateStructuredSyslogSLAProfileRequest) (*services.UpdateStructuredSyslogSLAProfileResponse, error) {
	log.Println(" UpdateStructuredSyslogSLAProfile ENTERED")

	obj := request.GetStructuredSyslogSLAProfile()

	intentObj := StructuredSyslogSLAProfileIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogSLAProfileIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogSLAProfileIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "StructuredSyslogSLAProfile")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateStructuredSyslogSLAProfile", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateStructuredSyslogSLAProfile(ctx, request)
}

// DeleteStructuredSyslogSLAProfile handles delete request
func (service *PluginService) DeleteStructuredSyslogSLAProfile(ctx context.Context, request *services.DeleteStructuredSyslogSLAProfileRequest) (*services.DeleteStructuredSyslogSLAProfileResponse, error) {
	log.Println(" DeleteStructuredSyslogSLAProfile ENTERED")

	objUUID := request.GetID()

	//intentObj := StructuredSyslogSLAProfileIntent {
	//StructuredSyslogSLAProfile: *obj,
	//}

	//EvaluateDependencies(intentObj, "StructuredSyslogSLAProfile")

	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogSLAProfileIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteStructuredSyslogSLAProfile", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteStructuredSyslogSLAProfile(ctx, request)
}

// GetStructuredSyslogSLAProfile handles get request
func (service *PluginService) GetStructuredSyslogSLAProfile(ctx context.Context, request *services.GetStructuredSyslogSLAProfileRequest) (*services.GetStructuredSyslogSLAProfileResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogSLAProfile")
	if !ok {
		return nil, errors.New("StructuredSyslogSLAProfile get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("StructuredSyslogSLAProfile get failed ")
	}

	response := &services.GetStructuredSyslogSLAProfileResponse{
		StructuredSyslogSLAProfile: obj.(*models.StructuredSyslogSLAProfile),
	}
	return response, nil
}
