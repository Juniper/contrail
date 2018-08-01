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

// AliasIPPoolIntent
//   A struct to store attributes related to AliasIPPool
//   needed by Intent Compiler
type AliasIPPoolIntent struct {
	Uuid string
}

// EvaluateAliasIPPool - evaluates the AliasIPPool
func EvaluateAliasIPPool(obj interface{}) {
	resourceObj := obj.(AliasIPPoolIntent)
	log.Println("EvaluateAliasIPPool Called ", resourceObj)
}

// CreateAliasIPPool handles create request
func (service *PluginService) CreateAliasIPPool(ctx context.Context, request *services.CreateAliasIPPoolRequest) (*services.CreateAliasIPPoolResponse, error) {
	log.Println(" CreateAliasIPPool Entered")

	obj := request.GetAliasIPPool()

	intentObj := AliasIPPoolIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AliasIPPoolIntent"); !ok {
		compilationif.ObjsCache.Store("AliasIPPoolIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("AliasIPPoolIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAliasIPPool", objMap.(*sync.Map))

	EvaluateDependencies(obj, "AliasIPPool")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAliasIPPool(ctx, request)
}

// UpdateAliasIPPool handles update request
func (service *PluginService) UpdateAliasIPPool(ctx context.Context, request *services.UpdateAliasIPPoolRequest) (*services.UpdateAliasIPPoolResponse, error) {
	log.Println(" UpdateAliasIPPool ENTERED")

	obj := request.GetAliasIPPool()

	intentObj := AliasIPPoolIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AliasIPPoolIntent"); !ok {
		compilationif.ObjsCache.Store("AliasIPPoolIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "AliasIPPool")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAliasIPPool", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAliasIPPool(ctx, request)
}

// DeleteAliasIPPool handles delete request
func (service *PluginService) DeleteAliasIPPool(ctx context.Context, request *services.DeleteAliasIPPoolRequest) (*services.DeleteAliasIPPoolResponse, error) {
	log.Println(" DeleteAliasIPPool ENTERED")

	objUUID := request.GetID()

	//intentObj := AliasIPPoolIntent {
	//AliasIPPool: *obj,
	//}

	//EvaluateDependencies(intentObj, "AliasIPPool")

	objMap, ok := compilationif.ObjsCache.Load("AliasIPPoolIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAliasIPPool", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAliasIPPool(ctx, request)
}

// GetAliasIPPool handles get request
func (service *PluginService) GetAliasIPPool(ctx context.Context, request *services.GetAliasIPPoolRequest) (*services.GetAliasIPPoolResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("AliasIPPool")
	if !ok {
		return nil, errors.New("AliasIPPool get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("AliasIPPool get failed ")
	}

	response := &services.GetAliasIPPoolResponse{
		AliasIPPool: obj.(*models.AliasIPPool),
	}
	return response, nil
}
