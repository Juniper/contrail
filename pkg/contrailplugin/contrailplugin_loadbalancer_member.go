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

// LoadbalancerMemberIntent
//   A struct to store attributes related to LoadbalancerMember
//   needed by Intent Compiler
type LoadbalancerMemberIntent struct {
	Uuid string
}

// EvaluateLoadbalancerMember - evaluates the LoadbalancerMember
func EvaluateLoadbalancerMember(obj interface{}) {
	resourceObj := obj.(LoadbalancerMemberIntent)
	log.Println("EvaluateLoadbalancerMember Called ", resourceObj)
}

// CreateLoadbalancerMember handles create request
func (service *PluginService) CreateLoadbalancerMember(ctx context.Context, request *services.CreateLoadbalancerMemberRequest) (*services.CreateLoadbalancerMemberResponse, error) {
	log.Println(" CreateLoadbalancerMember Entered")

	obj := request.GetLoadbalancerMember()

	intentObj := LoadbalancerMemberIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerMemberIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerMemberIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerMemberIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLoadbalancerMember", objMap.(*sync.Map))

	EvaluateDependencies(obj, "LoadbalancerMember")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerMember(ctx, request)
}

// UpdateLoadbalancerMember handles update request
func (service *PluginService) UpdateLoadbalancerMember(ctx context.Context, request *services.UpdateLoadbalancerMemberRequest) (*services.UpdateLoadbalancerMemberResponse, error) {
	log.Println(" UpdateLoadbalancerMember ENTERED")

	obj := request.GetLoadbalancerMember()

	intentObj := LoadbalancerMemberIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerMemberIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerMemberIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "LoadbalancerMember")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLoadbalancerMember", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerMember(ctx, request)
}

// DeleteLoadbalancerMember handles delete request
func (service *PluginService) DeleteLoadbalancerMember(ctx context.Context, request *services.DeleteLoadbalancerMemberRequest) (*services.DeleteLoadbalancerMemberResponse, error) {
	log.Println(" DeleteLoadbalancerMember ENTERED")

	objUUID := request.GetID()

	//intentObj := LoadbalancerMemberIntent {
	//LoadbalancerMember: *obj,
	//}

	//EvaluateDependencies(intentObj, "LoadbalancerMember")

	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerMemberIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLoadbalancerMember", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerMember(ctx, request)
}

// GetLoadbalancerMember handles get request
func (service *PluginService) GetLoadbalancerMember(ctx context.Context, request *services.GetLoadbalancerMemberRequest) (*services.GetLoadbalancerMemberResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerMember")
	if !ok {
		return nil, errors.New("LoadbalancerMember get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("LoadbalancerMember get failed ")
	}

	response := &services.GetLoadbalancerMemberResponse{
		LoadbalancerMember: obj.(*models.LoadbalancerMember),
	}
	return response, nil
}
