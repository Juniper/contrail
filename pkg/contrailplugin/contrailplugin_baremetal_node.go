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

// BaremetalNodeIntent
//   A struct to store attributes related to BaremetalNode
//   needed by Intent Compiler
type BaremetalNodeIntent struct {
	Uuid string
}

// EvaluateBaremetalNode - evaluates the BaremetalNode
func EvaluateBaremetalNode(obj interface{}) {
	resourceObj := obj.(BaremetalNodeIntent)
	log.Println("EvaluateBaremetalNode Called ", resourceObj)
}

// CreateBaremetalNode handles create request
func (service *PluginService) CreateBaremetalNode(ctx context.Context, request *services.CreateBaremetalNodeRequest) (*services.CreateBaremetalNodeResponse, error) {
	log.Println(" CreateBaremetalNode Entered")

	obj := request.GetBaremetalNode()

	intentObj := BaremetalNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BaremetalNodeIntent"); !ok {
		compilationif.ObjsCache.Store("BaremetalNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("BaremetalNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateBaremetalNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "BaremetalNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateBaremetalNode(ctx, request)
}

// UpdateBaremetalNode handles update request
func (service *PluginService) UpdateBaremetalNode(ctx context.Context, request *services.UpdateBaremetalNodeRequest) (*services.UpdateBaremetalNodeResponse, error) {
	log.Println(" UpdateBaremetalNode ENTERED")

	obj := request.GetBaremetalNode()

	intentObj := BaremetalNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BaremetalNodeIntent"); !ok {
		compilationif.ObjsCache.Store("BaremetalNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "BaremetalNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateBaremetalNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateBaremetalNode(ctx, request)
}

// DeleteBaremetalNode handles delete request
func (service *PluginService) DeleteBaremetalNode(ctx context.Context, request *services.DeleteBaremetalNodeRequest) (*services.DeleteBaremetalNodeResponse, error) {
	log.Println(" DeleteBaremetalNode ENTERED")

	objUUID := request.GetID()

	//intentObj := BaremetalNodeIntent {
	//BaremetalNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "BaremetalNode")

	objMap, ok := compilationif.ObjsCache.Load("BaremetalNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteBaremetalNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteBaremetalNode(ctx, request)
}

// GetBaremetalNode handles get request
func (service *PluginService) GetBaremetalNode(ctx context.Context, request *services.GetBaremetalNodeRequest) (*services.GetBaremetalNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("BaremetalNode")
	if !ok {
		return nil, errors.New("BaremetalNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("BaremetalNode get failed ")
	}

	response := &services.GetBaremetalNodeResponse{
		BaremetalNode: obj.(*models.BaremetalNode),
	}
	return response, nil
}
