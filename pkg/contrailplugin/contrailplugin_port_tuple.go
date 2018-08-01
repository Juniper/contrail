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

// PortTupleIntent
//   A struct to store attributes related to PortTuple
//   needed by Intent Compiler
type PortTupleIntent struct {
	Uuid string
}

// EvaluatePortTuple - evaluates the PortTuple
func EvaluatePortTuple(obj interface{}) {
	resourceObj := obj.(PortTupleIntent)
	log.Println("EvaluatePortTuple Called ", resourceObj)
}

// CreatePortTuple handles create request
func (service *PluginService) CreatePortTuple(ctx context.Context, request *services.CreatePortTupleRequest) (*services.CreatePortTupleResponse, error) {
	log.Println(" CreatePortTuple Entered")

	obj := request.GetPortTuple()

	intentObj := PortTupleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PortTupleIntent"); !ok {
		compilationif.ObjsCache.Store("PortTupleIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("PortTupleIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreatePortTuple", objMap.(*sync.Map))

	EvaluateDependencies(obj, "PortTuple")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreatePortTuple(ctx, request)
}

// UpdatePortTuple handles update request
func (service *PluginService) UpdatePortTuple(ctx context.Context, request *services.UpdatePortTupleRequest) (*services.UpdatePortTupleResponse, error) {
	log.Println(" UpdatePortTuple ENTERED")

	obj := request.GetPortTuple()

	intentObj := PortTupleIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PortTupleIntent"); !ok {
		compilationif.ObjsCache.Store("PortTupleIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "PortTuple")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdatePortTuple", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdatePortTuple(ctx, request)
}

// DeletePortTuple handles delete request
func (service *PluginService) DeletePortTuple(ctx context.Context, request *services.DeletePortTupleRequest) (*services.DeletePortTupleResponse, error) {
	log.Println(" DeletePortTuple ENTERED")

	objUUID := request.GetID()

	//intentObj := PortTupleIntent {
	//PortTuple: *obj,
	//}

	//EvaluateDependencies(intentObj, "PortTuple")

	objMap, ok := compilationif.ObjsCache.Load("PortTupleIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeletePortTuple", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeletePortTuple(ctx, request)
}

// GetPortTuple handles get request
func (service *PluginService) GetPortTuple(ctx context.Context, request *services.GetPortTupleRequest) (*services.GetPortTupleResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("PortTuple")
	if !ok {
		return nil, errors.New("PortTuple get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("PortTuple get failed ")
	}

	response := &services.GetPortTupleResponse{
		PortTuple: obj.(*models.PortTuple),
	}
	return response, nil
}
