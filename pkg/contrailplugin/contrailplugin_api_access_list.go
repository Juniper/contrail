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

// APIAccessListIntent
//   A struct to store attributes related to APIAccessList
//   needed by Intent Compiler
type APIAccessListIntent struct {
	Uuid string
}

// EvaluateAPIAccessList - evaluates the APIAccessList
func EvaluateAPIAccessList(obj interface{}) {
	resourceObj := obj.(APIAccessListIntent)
	log.Println("EvaluateAPIAccessList Called ", resourceObj)
}

// CreateAPIAccessList handles create request
func (service *PluginService) CreateAPIAccessList(ctx context.Context, request *services.CreateAPIAccessListRequest) (*services.CreateAPIAccessListResponse, error) {
	log.Println(" CreateAPIAccessList Entered")

	obj := request.GetAPIAccessList()

	intentObj := APIAccessListIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("APIAccessListIntent"); !ok {
		compilationif.ObjsCache.Store("APIAccessListIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("APIAccessListIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAPIAccessList", objMap.(*sync.Map))

	EvaluateDependencies(obj, "APIAccessList")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAPIAccessList(ctx, request)
}

// UpdateAPIAccessList handles update request
func (service *PluginService) UpdateAPIAccessList(ctx context.Context, request *services.UpdateAPIAccessListRequest) (*services.UpdateAPIAccessListResponse, error) {
	log.Println(" UpdateAPIAccessList ENTERED")

	obj := request.GetAPIAccessList()

	intentObj := APIAccessListIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("APIAccessListIntent"); !ok {
		compilationif.ObjsCache.Store("APIAccessListIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "APIAccessList")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAPIAccessList", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAPIAccessList(ctx, request)
}

// DeleteAPIAccessList handles delete request
func (service *PluginService) DeleteAPIAccessList(ctx context.Context, request *services.DeleteAPIAccessListRequest) (*services.DeleteAPIAccessListResponse, error) {
	log.Println(" DeleteAPIAccessList ENTERED")

	objUUID := request.GetID()

	//intentObj := APIAccessListIntent {
	//APIAccessList: *obj,
	//}

	//EvaluateDependencies(intentObj, "APIAccessList")

	objMap, ok := compilationif.ObjsCache.Load("APIAccessListIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAPIAccessList", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAPIAccessList(ctx, request)
}

// GetAPIAccessList handles get request
func (service *PluginService) GetAPIAccessList(ctx context.Context, request *services.GetAPIAccessListRequest) (*services.GetAPIAccessListResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("APIAccessList")
	if !ok {
		return nil, errors.New("APIAccessList get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("APIAccessList get failed ")
	}

	response := &services.GetAPIAccessListResponse{
		APIAccessList: obj.(*models.APIAccessList),
	}
	return response, nil
}
