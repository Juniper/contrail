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

// QosConfigIntent
//   A struct to store attributes related to QosConfig
//   needed by Intent Compiler
type QosConfigIntent struct {
	Uuid string
}

// EvaluateQosConfig - evaluates the QosConfig
func EvaluateQosConfig(obj interface{}) {
	resourceObj := obj.(QosConfigIntent)
	log.Println("EvaluateQosConfig Called ", resourceObj)
}

// CreateQosConfig handles create request
func (service *PluginService) CreateQosConfig(ctx context.Context, request *services.CreateQosConfigRequest) (*services.CreateQosConfigResponse, error) {
	log.Println(" CreateQosConfig Entered")

	obj := request.GetQosConfig()

	intentObj := QosConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("QosConfigIntent"); !ok {
		compilationif.ObjsCache.Store("QosConfigIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("QosConfigIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateQosConfig", objMap.(*sync.Map))

	EvaluateDependencies(obj, "QosConfig")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateQosConfig(ctx, request)
}

// UpdateQosConfig handles update request
func (service *PluginService) UpdateQosConfig(ctx context.Context, request *services.UpdateQosConfigRequest) (*services.UpdateQosConfigResponse, error) {
	log.Println(" UpdateQosConfig ENTERED")

	obj := request.GetQosConfig()

	intentObj := QosConfigIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("QosConfigIntent"); !ok {
		compilationif.ObjsCache.Store("QosConfigIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "QosConfig")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateQosConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateQosConfig(ctx, request)
}

// DeleteQosConfig handles delete request
func (service *PluginService) DeleteQosConfig(ctx context.Context, request *services.DeleteQosConfigRequest) (*services.DeleteQosConfigResponse, error) {
	log.Println(" DeleteQosConfig ENTERED")

	objUUID := request.GetID()

	//intentObj := QosConfigIntent {
	//QosConfig: *obj,
	//}

	//EvaluateDependencies(intentObj, "QosConfig")

	objMap, ok := compilationif.ObjsCache.Load("QosConfigIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteQosConfig", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteQosConfig(ctx, request)
}

// GetQosConfig handles get request
func (service *PluginService) GetQosConfig(ctx context.Context, request *services.GetQosConfigRequest) (*services.GetQosConfigResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("QosConfig")
	if !ok {
		return nil, errors.New("QosConfig get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("QosConfig get failed ")
	}

	response := &services.GetQosConfigResponse{
		QosConfig: obj.(*models.QosConfig),
	}
	return response, nil
}
