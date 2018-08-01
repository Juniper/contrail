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

// LoadbalancerHealthmonitorIntent
//   A struct to store attributes related to LoadbalancerHealthmonitor
//   needed by Intent Compiler
type LoadbalancerHealthmonitorIntent struct {
	Uuid string
}

// EvaluateLoadbalancerHealthmonitor - evaluates the LoadbalancerHealthmonitor
func EvaluateLoadbalancerHealthmonitor(obj interface{}) {
	resourceObj := obj.(LoadbalancerHealthmonitorIntent)
	log.Println("EvaluateLoadbalancerHealthmonitor Called ", resourceObj)
}

// CreateLoadbalancerHealthmonitor handles create request
func (service *PluginService) CreateLoadbalancerHealthmonitor(ctx context.Context, request *services.CreateLoadbalancerHealthmonitorRequest) (*services.CreateLoadbalancerHealthmonitorResponse, error) {
	log.Println(" CreateLoadbalancerHealthmonitor Entered")

	obj := request.GetLoadbalancerHealthmonitor()

	intentObj := LoadbalancerHealthmonitorIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerHealthmonitorIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerHealthmonitorIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerHealthmonitorIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLoadbalancerHealthmonitor", objMap.(*sync.Map))

	EvaluateDependencies(obj, "LoadbalancerHealthmonitor")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerHealthmonitor(ctx, request)
}

// UpdateLoadbalancerHealthmonitor handles update request
func (service *PluginService) UpdateLoadbalancerHealthmonitor(ctx context.Context, request *services.UpdateLoadbalancerHealthmonitorRequest) (*services.UpdateLoadbalancerHealthmonitorResponse, error) {
	log.Println(" UpdateLoadbalancerHealthmonitor ENTERED")

	obj := request.GetLoadbalancerHealthmonitor()

	intentObj := LoadbalancerHealthmonitorIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerHealthmonitorIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerHealthmonitorIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "LoadbalancerHealthmonitor")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLoadbalancerHealthmonitor", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerHealthmonitor(ctx, request)
}

// DeleteLoadbalancerHealthmonitor handles delete request
func (service *PluginService) DeleteLoadbalancerHealthmonitor(ctx context.Context, request *services.DeleteLoadbalancerHealthmonitorRequest) (*services.DeleteLoadbalancerHealthmonitorResponse, error) {
	log.Println(" DeleteLoadbalancerHealthmonitor ENTERED")

	objUUID := request.GetID()

	//intentObj := LoadbalancerHealthmonitorIntent {
	//LoadbalancerHealthmonitor: *obj,
	//}

	//EvaluateDependencies(intentObj, "LoadbalancerHealthmonitor")

	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerHealthmonitorIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLoadbalancerHealthmonitor", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerHealthmonitor(ctx, request)
}

// GetLoadbalancerHealthmonitor handles get request
func (service *PluginService) GetLoadbalancerHealthmonitor(ctx context.Context, request *services.GetLoadbalancerHealthmonitorRequest) (*services.GetLoadbalancerHealthmonitorResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerHealthmonitor")
	if !ok {
		return nil, errors.New("LoadbalancerHealthmonitor get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("LoadbalancerHealthmonitor get failed ")
	}

	response := &services.GetLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitor: obj.(*models.LoadbalancerHealthmonitor),
	}
	return response, nil
}
