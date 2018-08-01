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

// QosQueueIntent
//   A struct to store attributes related to QosQueue
//   needed by Intent Compiler
type QosQueueIntent struct {
	Uuid string
}

// EvaluateQosQueue - evaluates the QosQueue
func EvaluateQosQueue(obj interface{}) {
	resourceObj := obj.(QosQueueIntent)
	log.Println("EvaluateQosQueue Called ", resourceObj)
}

// CreateQosQueue handles create request
func (service *PluginService) CreateQosQueue(ctx context.Context, request *services.CreateQosQueueRequest) (*services.CreateQosQueueResponse, error) {
	log.Println(" CreateQosQueue Entered")

	obj := request.GetQosQueue()

	intentObj := QosQueueIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("QosQueueIntent"); !ok {
		compilationif.ObjsCache.Store("QosQueueIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("QosQueueIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateQosQueue", objMap.(*sync.Map))

	EvaluateDependencies(obj, "QosQueue")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateQosQueue(ctx, request)
}

// UpdateQosQueue handles update request
func (service *PluginService) UpdateQosQueue(ctx context.Context, request *services.UpdateQosQueueRequest) (*services.UpdateQosQueueResponse, error) {
	log.Println(" UpdateQosQueue ENTERED")

	obj := request.GetQosQueue()

	intentObj := QosQueueIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("QosQueueIntent"); !ok {
		compilationif.ObjsCache.Store("QosQueueIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "QosQueue")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateQosQueue", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateQosQueue(ctx, request)
}

// DeleteQosQueue handles delete request
func (service *PluginService) DeleteQosQueue(ctx context.Context, request *services.DeleteQosQueueRequest) (*services.DeleteQosQueueResponse, error) {
	log.Println(" DeleteQosQueue ENTERED")

	objUUID := request.GetID()

	//intentObj := QosQueueIntent {
	//QosQueue: *obj,
	//}

	//EvaluateDependencies(intentObj, "QosQueue")

	objMap, ok := compilationif.ObjsCache.Load("QosQueueIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteQosQueue", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteQosQueue(ctx, request)
}

// GetQosQueue handles get request
func (service *PluginService) GetQosQueue(ctx context.Context, request *services.GetQosQueueRequest) (*services.GetQosQueueResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("QosQueue")
	if !ok {
		return nil, errors.New("QosQueue get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("QosQueue get failed ")
	}

	response := &services.GetQosQueueResponse{
		QosQueue: obj.(*models.QosQueue),
	}
	return response, nil
}
