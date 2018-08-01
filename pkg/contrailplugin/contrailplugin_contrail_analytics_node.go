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

// ContrailAnalyticsNodeIntent
//   A struct to store attributes related to ContrailAnalyticsNode
//   needed by Intent Compiler
type ContrailAnalyticsNodeIntent struct {
	Uuid string
}

// EvaluateContrailAnalyticsNode - evaluates the ContrailAnalyticsNode
func EvaluateContrailAnalyticsNode(obj interface{}) {
	resourceObj := obj.(ContrailAnalyticsNodeIntent)
	log.Println("EvaluateContrailAnalyticsNode Called ", resourceObj)
}

// CreateContrailAnalyticsNode handles create request
func (service *PluginService) CreateContrailAnalyticsNode(ctx context.Context, request *services.CreateContrailAnalyticsNodeRequest) (*services.CreateContrailAnalyticsNodeResponse, error) {
	log.Println(" CreateContrailAnalyticsNode Entered")

	obj := request.GetContrailAnalyticsNode()

	intentObj := ContrailAnalyticsNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailAnalyticsNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailAnalyticsNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ContrailAnalyticsNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateContrailAnalyticsNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ContrailAnalyticsNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateContrailAnalyticsNode(ctx, request)
}

// UpdateContrailAnalyticsNode handles update request
func (service *PluginService) UpdateContrailAnalyticsNode(ctx context.Context, request *services.UpdateContrailAnalyticsNodeRequest) (*services.UpdateContrailAnalyticsNodeResponse, error) {
	log.Println(" UpdateContrailAnalyticsNode ENTERED")

	obj := request.GetContrailAnalyticsNode()

	intentObj := ContrailAnalyticsNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ContrailAnalyticsNodeIntent"); !ok {
		compilationif.ObjsCache.Store("ContrailAnalyticsNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ContrailAnalyticsNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateContrailAnalyticsNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateContrailAnalyticsNode(ctx, request)
}

// DeleteContrailAnalyticsNode handles delete request
func (service *PluginService) DeleteContrailAnalyticsNode(ctx context.Context, request *services.DeleteContrailAnalyticsNodeRequest) (*services.DeleteContrailAnalyticsNodeResponse, error) {
	log.Println(" DeleteContrailAnalyticsNode ENTERED")

	objUUID := request.GetID()

	//intentObj := ContrailAnalyticsNodeIntent {
	//ContrailAnalyticsNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "ContrailAnalyticsNode")

	objMap, ok := compilationif.ObjsCache.Load("ContrailAnalyticsNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteContrailAnalyticsNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteContrailAnalyticsNode(ctx, request)
}

// GetContrailAnalyticsNode handles get request
func (service *PluginService) GetContrailAnalyticsNode(ctx context.Context, request *services.GetContrailAnalyticsNodeRequest) (*services.GetContrailAnalyticsNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ContrailAnalyticsNode")
	if !ok {
		return nil, errors.New("ContrailAnalyticsNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ContrailAnalyticsNode get failed ")
	}

	response := &services.GetContrailAnalyticsNodeResponse{
		ContrailAnalyticsNode: obj.(*models.ContrailAnalyticsNode),
	}
	return response, nil
}
