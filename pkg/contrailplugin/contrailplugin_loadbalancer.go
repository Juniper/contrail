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

// LoadbalancerIntent
//   A struct to store attributes related to Loadbalancer
//   needed by Intent Compiler
type LoadbalancerIntent struct {
	Uuid string
}

// EvaluateLoadbalancer - evaluates the Loadbalancer
func EvaluateLoadbalancer(obj interface{}) {
	resourceObj := obj.(LoadbalancerIntent)
	log.Println("EvaluateLoadbalancer Called ", resourceObj)
}

// CreateLoadbalancer handles create request
func (service *PluginService) CreateLoadbalancer(ctx context.Context, request *services.CreateLoadbalancerRequest) (*services.CreateLoadbalancerResponse, error) {
	log.Println(" CreateLoadbalancer Entered")

	obj := request.GetLoadbalancer()

	intentObj := LoadbalancerIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLoadbalancer", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Loadbalancer")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancer(ctx, request)
}

// UpdateLoadbalancer handles update request
func (service *PluginService) UpdateLoadbalancer(ctx context.Context, request *services.UpdateLoadbalancerRequest) (*services.UpdateLoadbalancerResponse, error) {
	log.Println(" UpdateLoadbalancer ENTERED")

	obj := request.GetLoadbalancer()

	intentObj := LoadbalancerIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Loadbalancer")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLoadbalancer", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancer(ctx, request)
}

// DeleteLoadbalancer handles delete request
func (service *PluginService) DeleteLoadbalancer(ctx context.Context, request *services.DeleteLoadbalancerRequest) (*services.DeleteLoadbalancerResponse, error) {
	log.Println(" DeleteLoadbalancer ENTERED")

	objUUID := request.GetID()

	//intentObj := LoadbalancerIntent {
	//Loadbalancer: *obj,
	//}

	//EvaluateDependencies(intentObj, "Loadbalancer")

	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLoadbalancer", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancer(ctx, request)
}

// GetLoadbalancer handles get request
func (service *PluginService) GetLoadbalancer(ctx context.Context, request *services.GetLoadbalancerRequest) (*services.GetLoadbalancerResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Loadbalancer")
	if !ok {
		return nil, errors.New("Loadbalancer get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Loadbalancer get failed ")
	}

	response := &services.GetLoadbalancerResponse{
		Loadbalancer: obj.(*models.Loadbalancer),
	}
	return response, nil
}
