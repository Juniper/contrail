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

// LinkAggregationGroupIntent
//   A struct to store attributes related to LinkAggregationGroup
//   needed by Intent Compiler
type LinkAggregationGroupIntent struct {
	Uuid string
}

// EvaluateLinkAggregationGroup - evaluates the LinkAggregationGroup
func EvaluateLinkAggregationGroup(obj interface{}) {
	resourceObj := obj.(LinkAggregationGroupIntent)
	log.Println("EvaluateLinkAggregationGroup Called ", resourceObj)
}

// CreateLinkAggregationGroup handles create request
func (service *PluginService) CreateLinkAggregationGroup(ctx context.Context, request *services.CreateLinkAggregationGroupRequest) (*services.CreateLinkAggregationGroupResponse, error) {
	log.Println(" CreateLinkAggregationGroup Entered")

	obj := request.GetLinkAggregationGroup()

	intentObj := LinkAggregationGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LinkAggregationGroupIntent"); !ok {
		compilationif.ObjsCache.Store("LinkAggregationGroupIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LinkAggregationGroupIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLinkAggregationGroup", objMap.(*sync.Map))

	EvaluateDependencies(obj, "LinkAggregationGroup")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLinkAggregationGroup(ctx, request)
}

// UpdateLinkAggregationGroup handles update request
func (service *PluginService) UpdateLinkAggregationGroup(ctx context.Context, request *services.UpdateLinkAggregationGroupRequest) (*services.UpdateLinkAggregationGroupResponse, error) {
	log.Println(" UpdateLinkAggregationGroup ENTERED")

	obj := request.GetLinkAggregationGroup()

	intentObj := LinkAggregationGroupIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LinkAggregationGroupIntent"); !ok {
		compilationif.ObjsCache.Store("LinkAggregationGroupIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "LinkAggregationGroup")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLinkAggregationGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLinkAggregationGroup(ctx, request)
}

// DeleteLinkAggregationGroup handles delete request
func (service *PluginService) DeleteLinkAggregationGroup(ctx context.Context, request *services.DeleteLinkAggregationGroupRequest) (*services.DeleteLinkAggregationGroupResponse, error) {
	log.Println(" DeleteLinkAggregationGroup ENTERED")

	objUUID := request.GetID()

	//intentObj := LinkAggregationGroupIntent {
	//LinkAggregationGroup: *obj,
	//}

	//EvaluateDependencies(intentObj, "LinkAggregationGroup")

	objMap, ok := compilationif.ObjsCache.Load("LinkAggregationGroupIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLinkAggregationGroup", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLinkAggregationGroup(ctx, request)
}

// GetLinkAggregationGroup handles get request
func (service *PluginService) GetLinkAggregationGroup(ctx context.Context, request *services.GetLinkAggregationGroupRequest) (*services.GetLinkAggregationGroupResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("LinkAggregationGroup")
	if !ok {
		return nil, errors.New("LinkAggregationGroup get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("LinkAggregationGroup get failed ")
	}

	response := &services.GetLinkAggregationGroupResponse{
		LinkAggregationGroup: obj.(*models.LinkAggregationGroup),
	}
	return response, nil
}
