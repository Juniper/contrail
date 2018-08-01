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

// BGPAsAServiceIntent
//   A struct to store attributes related to BGPAsAService
//   needed by Intent Compiler
type BGPAsAServiceIntent struct {
	Uuid string
}

// EvaluateBGPAsAService - evaluates the BGPAsAService
func EvaluateBGPAsAService(obj interface{}) {
	resourceObj := obj.(BGPAsAServiceIntent)
	log.Println("EvaluateBGPAsAService Called ", resourceObj)
}

// CreateBGPAsAService handles create request
func (service *PluginService) CreateBGPAsAService(ctx context.Context, request *services.CreateBGPAsAServiceRequest) (*services.CreateBGPAsAServiceResponse, error) {
	log.Println(" CreateBGPAsAService Entered")

	obj := request.GetBGPAsAService()

	intentObj := BGPAsAServiceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BGPAsAServiceIntent"); !ok {
		compilationif.ObjsCache.Store("BGPAsAServiceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("BGPAsAServiceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateBGPAsAService", objMap.(*sync.Map))

	EvaluateDependencies(obj, "BGPAsAService")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateBGPAsAService(ctx, request)
}

// UpdateBGPAsAService handles update request
func (service *PluginService) UpdateBGPAsAService(ctx context.Context, request *services.UpdateBGPAsAServiceRequest) (*services.UpdateBGPAsAServiceResponse, error) {
	log.Println(" UpdateBGPAsAService ENTERED")

	obj := request.GetBGPAsAService()

	intentObj := BGPAsAServiceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BGPAsAServiceIntent"); !ok {
		compilationif.ObjsCache.Store("BGPAsAServiceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "BGPAsAService")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateBGPAsAService", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateBGPAsAService(ctx, request)
}

// DeleteBGPAsAService handles delete request
func (service *PluginService) DeleteBGPAsAService(ctx context.Context, request *services.DeleteBGPAsAServiceRequest) (*services.DeleteBGPAsAServiceResponse, error) {
	log.Println(" DeleteBGPAsAService ENTERED")

	objUUID := request.GetID()

	//intentObj := BGPAsAServiceIntent {
	//BGPAsAService: *obj,
	//}

	//EvaluateDependencies(intentObj, "BGPAsAService")

	objMap, ok := compilationif.ObjsCache.Load("BGPAsAServiceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteBGPAsAService", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteBGPAsAService(ctx, request)
}

// GetBGPAsAService handles get request
func (service *PluginService) GetBGPAsAService(ctx context.Context, request *services.GetBGPAsAServiceRequest) (*services.GetBGPAsAServiceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("BGPAsAService")
	if !ok {
		return nil, errors.New("BGPAsAService get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("BGPAsAService get failed ")
	}

	response := &services.GetBGPAsAServiceResponse{
		BGPAsAService: obj.(*models.BGPAsAService),
	}
	return response, nil
}
