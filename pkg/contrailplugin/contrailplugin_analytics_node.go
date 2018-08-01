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

// AnalyticsNodeIntent
//   A struct to store attributes related to AnalyticsNode
//   needed by Intent Compiler
type AnalyticsNodeIntent struct {
	Uuid string
}

// EvaluateAnalyticsNode - evaluates the AnalyticsNode
func EvaluateAnalyticsNode(obj interface{}) {
	resourceObj := obj.(AnalyticsNodeIntent)
	log.Println("EvaluateAnalyticsNode Called ", resourceObj)
}

// CreateAnalyticsNode handles create request
func (service *PluginService) CreateAnalyticsNode(ctx context.Context, request *services.CreateAnalyticsNodeRequest) (*services.CreateAnalyticsNodeResponse, error) {
	log.Println(" CreateAnalyticsNode Entered")

	obj := request.GetAnalyticsNode()

	intentObj := AnalyticsNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AnalyticsNodeIntent"); !ok {
		compilationif.ObjsCache.Store("AnalyticsNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("AnalyticsNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAnalyticsNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "AnalyticsNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAnalyticsNode(ctx, request)
}

// UpdateAnalyticsNode handles update request
func (service *PluginService) UpdateAnalyticsNode(ctx context.Context, request *services.UpdateAnalyticsNodeRequest) (*services.UpdateAnalyticsNodeResponse, error) {
	log.Println(" UpdateAnalyticsNode ENTERED")

	obj := request.GetAnalyticsNode()

	intentObj := AnalyticsNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AnalyticsNodeIntent"); !ok {
		compilationif.ObjsCache.Store("AnalyticsNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "AnalyticsNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAnalyticsNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAnalyticsNode(ctx, request)
}

// DeleteAnalyticsNode handles delete request
func (service *PluginService) DeleteAnalyticsNode(ctx context.Context, request *services.DeleteAnalyticsNodeRequest) (*services.DeleteAnalyticsNodeResponse, error) {
	log.Println(" DeleteAnalyticsNode ENTERED")

	objUUID := request.GetID()

	//intentObj := AnalyticsNodeIntent {
	//AnalyticsNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "AnalyticsNode")

	objMap, ok := compilationif.ObjsCache.Load("AnalyticsNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAnalyticsNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAnalyticsNode(ctx, request)
}

// GetAnalyticsNode handles get request
func (service *PluginService) GetAnalyticsNode(ctx context.Context, request *services.GetAnalyticsNodeRequest) (*services.GetAnalyticsNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("AnalyticsNode")
	if !ok {
		return nil, errors.New("AnalyticsNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("AnalyticsNode get failed ")
	}

	response := &services.GetAnalyticsNodeResponse{
		AnalyticsNode: obj.(*models.AnalyticsNode),
	}
	return response, nil
}
