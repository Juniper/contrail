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

// BGPRouterIntent
//   A struct to store attributes related to BGPRouter
//   needed by Intent Compiler
type BGPRouterIntent struct {
	Uuid string
}

// EvaluateBGPRouter - evaluates the BGPRouter
func EvaluateBGPRouter(obj interface{}) {
	resourceObj := obj.(BGPRouterIntent)
	log.Println("EvaluateBGPRouter Called ", resourceObj)
}

// CreateBGPRouter handles create request
func (service *PluginService) CreateBGPRouter(ctx context.Context, request *services.CreateBGPRouterRequest) (*services.CreateBGPRouterResponse, error) {
	log.Println(" CreateBGPRouter Entered")

	obj := request.GetBGPRouter()

	intentObj := BGPRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BGPRouterIntent"); !ok {
		compilationif.ObjsCache.Store("BGPRouterIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("BGPRouterIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateBGPRouter", objMap.(*sync.Map))

	EvaluateDependencies(obj, "BGPRouter")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateBGPRouter(ctx, request)
}

// UpdateBGPRouter handles update request
func (service *PluginService) UpdateBGPRouter(ctx context.Context, request *services.UpdateBGPRouterRequest) (*services.UpdateBGPRouterResponse, error) {
	log.Println(" UpdateBGPRouter ENTERED")

	obj := request.GetBGPRouter()

	intentObj := BGPRouterIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BGPRouterIntent"); !ok {
		compilationif.ObjsCache.Store("BGPRouterIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "BGPRouter")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateBGPRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPRouter(ctx, request)
}

// DeleteBGPRouter handles delete request
func (service *PluginService) DeleteBGPRouter(ctx context.Context, request *services.DeleteBGPRouterRequest) (*services.DeleteBGPRouterResponse, error) {
	log.Println(" DeleteBGPRouter ENTERED")

	objUUID := request.GetID()

	//intentObj := BGPRouterIntent {
	//BGPRouter: *obj,
	//}

	//EvaluateDependencies(intentObj, "BGPRouter")

	objMap, ok := compilationif.ObjsCache.Load("BGPRouterIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteBGPRouter", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPRouter(ctx, request)
}

// GetBGPRouter handles get request
func (service *PluginService) GetBGPRouter(ctx context.Context, request *services.GetBGPRouterRequest) (*services.GetBGPRouterResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("BGPRouter")
	if !ok {
		return nil, errors.New("BGPRouter get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("BGPRouter get failed ")
	}

	response := &services.GetBGPRouterResponse{
		BGPRouter: obj.(*models.BGPRouter),
	}
	return response, nil
}
