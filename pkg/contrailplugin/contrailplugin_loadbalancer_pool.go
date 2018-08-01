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

// LoadbalancerPoolIntent
//   A struct to store attributes related to LoadbalancerPool
//   needed by Intent Compiler
type LoadbalancerPoolIntent struct {
	Uuid string
}

// EvaluateLoadbalancerPool - evaluates the LoadbalancerPool
func EvaluateLoadbalancerPool(obj interface{}) {
	resourceObj := obj.(LoadbalancerPoolIntent)
	log.Println("EvaluateLoadbalancerPool Called ", resourceObj)
}

// CreateLoadbalancerPool handles create request
func (service *PluginService) CreateLoadbalancerPool(ctx context.Context, request *services.CreateLoadbalancerPoolRequest) (*services.CreateLoadbalancerPoolResponse, error) {
	log.Println(" CreateLoadbalancerPool Entered")

	obj := request.GetLoadbalancerPool()

	intentObj := LoadbalancerPoolIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerPoolIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerPoolIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerPoolIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLoadbalancerPool", objMap.(*sync.Map))

	EvaluateDependencies(obj, "LoadbalancerPool")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerPool(ctx, request)
}

// UpdateLoadbalancerPool handles update request
func (service *PluginService) UpdateLoadbalancerPool(ctx context.Context, request *services.UpdateLoadbalancerPoolRequest) (*services.UpdateLoadbalancerPoolResponse, error) {
	log.Println(" UpdateLoadbalancerPool ENTERED")

	obj := request.GetLoadbalancerPool()

	intentObj := LoadbalancerPoolIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerPoolIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerPoolIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "LoadbalancerPool")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLoadbalancerPool", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerPool(ctx, request)
}

// DeleteLoadbalancerPool handles delete request
func (service *PluginService) DeleteLoadbalancerPool(ctx context.Context, request *services.DeleteLoadbalancerPoolRequest) (*services.DeleteLoadbalancerPoolResponse, error) {
	log.Println(" DeleteLoadbalancerPool ENTERED")

	objUUID := request.GetID()

	//intentObj := LoadbalancerPoolIntent {
	//LoadbalancerPool: *obj,
	//}

	//EvaluateDependencies(intentObj, "LoadbalancerPool")

	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerPoolIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLoadbalancerPool", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerPool(ctx, request)
}

// GetLoadbalancerPool handles get request
func (service *PluginService) GetLoadbalancerPool(ctx context.Context, request *services.GetLoadbalancerPoolRequest) (*services.GetLoadbalancerPoolResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerPool")
	if !ok {
		return nil, errors.New("LoadbalancerPool get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("LoadbalancerPool get failed ")
	}

	response := &services.GetLoadbalancerPoolResponse{
		LoadbalancerPool: obj.(*models.LoadbalancerPool),
	}
	return response, nil
}
