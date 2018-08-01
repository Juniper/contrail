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

// LoadbalancerListenerIntent
//   A struct to store attributes related to LoadbalancerListener
//   needed by Intent Compiler
type LoadbalancerListenerIntent struct {
	Uuid string
}

// EvaluateLoadbalancerListener - evaluates the LoadbalancerListener
func EvaluateLoadbalancerListener(obj interface{}) {
	resourceObj := obj.(LoadbalancerListenerIntent)
	log.Println("EvaluateLoadbalancerListener Called ", resourceObj)
}

// CreateLoadbalancerListener handles create request
func (service *PluginService) CreateLoadbalancerListener(ctx context.Context, request *services.CreateLoadbalancerListenerRequest) (*services.CreateLoadbalancerListenerResponse, error) {
	log.Println(" CreateLoadbalancerListener Entered")

	obj := request.GetLoadbalancerListener()

	intentObj := LoadbalancerListenerIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerListenerIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerListenerIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerListenerIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLoadbalancerListener", objMap.(*sync.Map))

	EvaluateDependencies(obj, "LoadbalancerListener")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLoadbalancerListener(ctx, request)
}

// UpdateLoadbalancerListener handles update request
func (service *PluginService) UpdateLoadbalancerListener(ctx context.Context, request *services.UpdateLoadbalancerListenerRequest) (*services.UpdateLoadbalancerListenerResponse, error) {
	log.Println(" UpdateLoadbalancerListener ENTERED")

	obj := request.GetLoadbalancerListener()

	intentObj := LoadbalancerListenerIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LoadbalancerListenerIntent"); !ok {
		compilationif.ObjsCache.Store("LoadbalancerListenerIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "LoadbalancerListener")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLoadbalancerListener", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLoadbalancerListener(ctx, request)
}

// DeleteLoadbalancerListener handles delete request
func (service *PluginService) DeleteLoadbalancerListener(ctx context.Context, request *services.DeleteLoadbalancerListenerRequest) (*services.DeleteLoadbalancerListenerResponse, error) {
	log.Println(" DeleteLoadbalancerListener ENTERED")

	objUUID := request.GetID()

	//intentObj := LoadbalancerListenerIntent {
	//LoadbalancerListener: *obj,
	//}

	//EvaluateDependencies(intentObj, "LoadbalancerListener")

	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerListenerIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLoadbalancerListener", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLoadbalancerListener(ctx, request)
}

// GetLoadbalancerListener handles get request
func (service *PluginService) GetLoadbalancerListener(ctx context.Context, request *services.GetLoadbalancerListenerRequest) (*services.GetLoadbalancerListenerResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("LoadbalancerListener")
	if !ok {
		return nil, errors.New("LoadbalancerListener get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("LoadbalancerListener get failed ")
	}

	response := &services.GetLoadbalancerListenerResponse{
		LoadbalancerListener: obj.(*models.LoadbalancerListener),
	}
	return response, nil
}
