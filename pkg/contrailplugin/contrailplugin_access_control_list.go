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

// AccessControlListIntent
//   A struct to store attributes related to AccessControlList
//   needed by Intent Compiler
type AccessControlListIntent struct {
	Uuid string
}

// EvaluateAccessControlList - evaluates the AccessControlList
func EvaluateAccessControlList(obj interface{}) {
	resourceObj := obj.(AccessControlListIntent)
	log.Println("EvaluateAccessControlList Called ", resourceObj)
}

// CreateAccessControlList handles create request
func (service *PluginService) CreateAccessControlList(ctx context.Context, request *services.CreateAccessControlListRequest) (*services.CreateAccessControlListResponse, error) {
	log.Println(" CreateAccessControlList Entered")

	obj := request.GetAccessControlList()

	intentObj := AccessControlListIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AccessControlListIntent"); !ok {
		compilationif.ObjsCache.Store("AccessControlListIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("AccessControlListIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateAccessControlList", objMap.(*sync.Map))

	EvaluateDependencies(obj, "AccessControlList")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateAccessControlList(ctx, request)
}

// UpdateAccessControlList handles update request
func (service *PluginService) UpdateAccessControlList(ctx context.Context, request *services.UpdateAccessControlListRequest) (*services.UpdateAccessControlListResponse, error) {
	log.Println(" UpdateAccessControlList ENTERED")

	obj := request.GetAccessControlList()

	intentObj := AccessControlListIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("AccessControlListIntent"); !ok {
		compilationif.ObjsCache.Store("AccessControlListIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "AccessControlList")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateAccessControlList", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateAccessControlList(ctx, request)
}

// DeleteAccessControlList handles delete request
func (service *PluginService) DeleteAccessControlList(ctx context.Context, request *services.DeleteAccessControlListRequest) (*services.DeleteAccessControlListResponse, error) {
	log.Println(" DeleteAccessControlList ENTERED")

	objUUID := request.GetID()

	//intentObj := AccessControlListIntent {
	//AccessControlList: *obj,
	//}

	//EvaluateDependencies(intentObj, "AccessControlList")

	objMap, ok := compilationif.ObjsCache.Load("AccessControlListIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteAccessControlList", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteAccessControlList(ctx, request)
}

// GetAccessControlList handles get request
func (service *PluginService) GetAccessControlList(ctx context.Context, request *services.GetAccessControlListRequest) (*services.GetAccessControlListResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("AccessControlList")
	if !ok {
		return nil, errors.New("AccessControlList get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("AccessControlList get failed ")
	}

	response := &services.GetAccessControlListResponse{
		AccessControlList: obj.(*models.AccessControlList),
	}
	return response, nil
}
