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

// PortIntent
//   A struct to store attributes related to Port
//   needed by Intent Compiler
type PortIntent struct {
	Uuid string
}

// EvaluatePort - evaluates the Port
func EvaluatePort(obj interface{}) {
	resourceObj := obj.(PortIntent)
	log.Println("EvaluatePort Called ", resourceObj)
}

// CreatePort handles create request
func (service *PluginService) CreatePort(ctx context.Context, request *services.CreatePortRequest) (*services.CreatePortResponse, error) {
	log.Println(" CreatePort Entered")

	obj := request.GetPort()

	intentObj := PortIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PortIntent"); !ok {
		compilationif.ObjsCache.Store("PortIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("PortIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreatePort", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Port")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreatePort(ctx, request)
}

// UpdatePort handles update request
func (service *PluginService) UpdatePort(ctx context.Context, request *services.UpdatePortRequest) (*services.UpdatePortResponse, error) {
	log.Println(" UpdatePort ENTERED")

	obj := request.GetPort()

	intentObj := PortIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PortIntent"); !ok {
		compilationif.ObjsCache.Store("PortIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Port")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdatePort", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdatePort(ctx, request)
}

// DeletePort handles delete request
func (service *PluginService) DeletePort(ctx context.Context, request *services.DeletePortRequest) (*services.DeletePortResponse, error) {
	log.Println(" DeletePort ENTERED")

	objUUID := request.GetID()

	//intentObj := PortIntent {
	//Port: *obj,
	//}

	//EvaluateDependencies(intentObj, "Port")

	objMap, ok := compilationif.ObjsCache.Load("PortIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeletePort", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeletePort(ctx, request)
}

// GetPort handles get request
func (service *PluginService) GetPort(ctx context.Context, request *services.GetPortRequest) (*services.GetPortResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Port")
	if !ok {
		return nil, errors.New("Port get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Port get failed ")
	}

	response := &services.GetPortResponse{
		Port: obj.(*models.Port),
	}
	return response, nil
}
