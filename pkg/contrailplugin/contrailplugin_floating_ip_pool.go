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

// FloatingIPPoolIntent
//   A struct to store attributes related to FloatingIPPool
//   needed by Intent Compiler
type FloatingIPPoolIntent struct {
	Uuid string
}

// EvaluateFloatingIPPool - evaluates the FloatingIPPool
func EvaluateFloatingIPPool(obj interface{}) {
	resourceObj := obj.(FloatingIPPoolIntent)
	log.Println("EvaluateFloatingIPPool Called ", resourceObj)
}

// CreateFloatingIPPool handles create request
func (service *PluginService) CreateFloatingIPPool(ctx context.Context, request *services.CreateFloatingIPPoolRequest) (*services.CreateFloatingIPPoolResponse, error) {
	log.Println(" CreateFloatingIPPool Entered")

	obj := request.GetFloatingIPPool()

	intentObj := FloatingIPPoolIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FloatingIPPoolIntent"); !ok {
		compilationif.ObjsCache.Store("FloatingIPPoolIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("FloatingIPPoolIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateFloatingIPPool", objMap.(*sync.Map))

	EvaluateDependencies(obj, "FloatingIPPool")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateFloatingIPPool(ctx, request)
}

// UpdateFloatingIPPool handles update request
func (service *PluginService) UpdateFloatingIPPool(ctx context.Context, request *services.UpdateFloatingIPPoolRequest) (*services.UpdateFloatingIPPoolResponse, error) {
	log.Println(" UpdateFloatingIPPool ENTERED")

	obj := request.GetFloatingIPPool()

	intentObj := FloatingIPPoolIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("FloatingIPPoolIntent"); !ok {
		compilationif.ObjsCache.Store("FloatingIPPoolIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "FloatingIPPool")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateFloatingIPPool", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateFloatingIPPool(ctx, request)
}

// DeleteFloatingIPPool handles delete request
func (service *PluginService) DeleteFloatingIPPool(ctx context.Context, request *services.DeleteFloatingIPPoolRequest) (*services.DeleteFloatingIPPoolResponse, error) {
	log.Println(" DeleteFloatingIPPool ENTERED")

	objUUID := request.GetID()

	//intentObj := FloatingIPPoolIntent {
	//FloatingIPPool: *obj,
	//}

	//EvaluateDependencies(intentObj, "FloatingIPPool")

	objMap, ok := compilationif.ObjsCache.Load("FloatingIPPoolIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteFloatingIPPool", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteFloatingIPPool(ctx, request)
}

// GetFloatingIPPool handles get request
func (service *PluginService) GetFloatingIPPool(ctx context.Context, request *services.GetFloatingIPPoolRequest) (*services.GetFloatingIPPoolResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("FloatingIPPool")
	if !ok {
		return nil, errors.New("FloatingIPPool get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("FloatingIPPool get failed ")
	}

	response := &services.GetFloatingIPPoolResponse{
		FloatingIPPool: obj.(*models.FloatingIPPool),
	}
	return response, nil
}
