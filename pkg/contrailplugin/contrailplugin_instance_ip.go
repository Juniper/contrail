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

// InstanceIPIntent
//   A struct to store attributes related to InstanceIP
//   needed by Intent Compiler
type InstanceIPIntent struct {
	Uuid string
}

// EvaluateInstanceIP - evaluates the InstanceIP
func EvaluateInstanceIP(obj interface{}) {
	resourceObj := obj.(InstanceIPIntent)
	log.Println("EvaluateInstanceIP Called ", resourceObj)
}

// CreateInstanceIP handles create request
func (service *PluginService) CreateInstanceIP(ctx context.Context, request *services.CreateInstanceIPRequest) (*services.CreateInstanceIPResponse, error) {
	log.Println(" CreateInstanceIP Entered")

	obj := request.GetInstanceIP()

	intentObj := InstanceIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("InstanceIPIntent"); !ok {
		compilationif.ObjsCache.Store("InstanceIPIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("InstanceIPIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateInstanceIP", objMap.(*sync.Map))

	EvaluateDependencies(obj, "InstanceIP")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateInstanceIP(ctx, request)
}

// UpdateInstanceIP handles update request
func (service *PluginService) UpdateInstanceIP(ctx context.Context, request *services.UpdateInstanceIPRequest) (*services.UpdateInstanceIPResponse, error) {
	log.Println(" UpdateInstanceIP ENTERED")

	obj := request.GetInstanceIP()

	intentObj := InstanceIPIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("InstanceIPIntent"); !ok {
		compilationif.ObjsCache.Store("InstanceIPIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "InstanceIP")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateInstanceIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateInstanceIP(ctx, request)
}

// DeleteInstanceIP handles delete request
func (service *PluginService) DeleteInstanceIP(ctx context.Context, request *services.DeleteInstanceIPRequest) (*services.DeleteInstanceIPResponse, error) {
	log.Println(" DeleteInstanceIP ENTERED")

	objUUID := request.GetID()

	//intentObj := InstanceIPIntent {
	//InstanceIP: *obj,
	//}

	//EvaluateDependencies(intentObj, "InstanceIP")

	objMap, ok := compilationif.ObjsCache.Load("InstanceIPIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteInstanceIP", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteInstanceIP(ctx, request)
}

// GetInstanceIP handles get request
func (service *PluginService) GetInstanceIP(ctx context.Context, request *services.GetInstanceIPRequest) (*services.GetInstanceIPResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("InstanceIP")
	if !ok {
		return nil, errors.New("InstanceIP get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("InstanceIP get failed ")
	}

	response := &services.GetInstanceIPResponse{
		InstanceIP: obj.(*models.InstanceIP),
	}
	return response, nil
}
