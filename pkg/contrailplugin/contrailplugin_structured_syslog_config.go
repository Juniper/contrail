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

// StructuredSyslogConfigIntent
//   A struct to store attributes related to StructuredSyslogConfig
//   needed by Intent Compiler
type StructuredSyslogConfigIntent struct {
	Uuid string
}

// EvaluateStructuredSyslogConfig - evaluates the StructuredSyslogConfig
func EvaluateStructuredSyslogConfig(obj interface{}) {
	resourceObj := obj.(StructuredSyslogConfigIntent)
	log.Println("EvaluateStructuredSyslogConfig Called ", resourceObj)
}

// CreateStructuredSyslogConfig handles create request
func (service *PluginService) CreateStructuredSyslogConfig(ctx context.Context, request *services.CreateStructuredSyslogConfigRequest) (*services.CreateStructuredSyslogConfigResponse, error) {
	log.Println(" CreateStructuredSyslogConfig Entered")

	obj := request.GetStructuredSyslogConfig()

	intentObj := StructuredSyslogConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogConfigIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogConfigIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogConfigIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateStructuredSyslogConfig", objMap.(*sync.Map))

	EvaluateDependencies(obj, "StructuredSyslogConfig")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateStructuredSyslogConfig(ctx, request)
}

// UpdateStructuredSyslogConfig handles update request
func (service *PluginService) UpdateStructuredSyslogConfig(ctx context.Context, request *services.UpdateStructuredSyslogConfigRequest) (*services.UpdateStructuredSyslogConfigResponse, error) {
	log.Println(" UpdateStructuredSyslogConfig ENTERED")

	obj := request.GetStructuredSyslogConfig()

	intentObj := StructuredSyslogConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogConfigIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogConfigIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "StructuredSyslogConfig")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateStructuredSyslogConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateStructuredSyslogConfig(ctx, request)
}

// DeleteStructuredSyslogConfig handles delete request
func (service *PluginService) DeleteStructuredSyslogConfig(ctx context.Context, request *services.DeleteStructuredSyslogConfigRequest) (*services.DeleteStructuredSyslogConfigResponse, error) {
	log.Println(" DeleteStructuredSyslogConfig ENTERED")

	objUUID := request.GetID()

	//intentObj := StructuredSyslogConfigIntent {
	//StructuredSyslogConfig: *obj,
	//}

	//EvaluateDependencies(intentObj, "StructuredSyslogConfig")

	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogConfigIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteStructuredSyslogConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteStructuredSyslogConfig(ctx, request)
}

// GetStructuredSyslogConfig handles get request
func (service *PluginService) GetStructuredSyslogConfig(ctx context.Context, request *services.GetStructuredSyslogConfigRequest) (*services.GetStructuredSyslogConfigResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogConfig")
	if !ok {
		return nil, errors.New("StructuredSyslogConfig get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("StructuredSyslogConfig get failed ")
	}

	response := &services.GetStructuredSyslogConfigResponse{
		StructuredSyslogConfig: obj.(*models.StructuredSyslogConfig),
	}
	return response, nil
}
