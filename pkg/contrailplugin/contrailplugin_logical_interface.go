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

// LogicalInterfaceIntent
//   A struct to store attributes related to LogicalInterface
//   needed by Intent Compiler
type LogicalInterfaceIntent struct {
	Uuid string
}

// EvaluateLogicalInterface - evaluates the LogicalInterface
func EvaluateLogicalInterface(obj interface{}) {
	resourceObj := obj.(LogicalInterfaceIntent)
	log.Println("EvaluateLogicalInterface Called ", resourceObj)
}

// CreateLogicalInterface handles create request
func (service *PluginService) CreateLogicalInterface(ctx context.Context, request *services.CreateLogicalInterfaceRequest) (*services.CreateLogicalInterfaceResponse, error) {
	log.Println(" CreateLogicalInterface Entered")

	obj := request.GetLogicalInterface()

	intentObj := LogicalInterfaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LogicalInterfaceIntent"); !ok {
		compilationif.ObjsCache.Store("LogicalInterfaceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("LogicalInterfaceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateLogicalInterface", objMap.(*sync.Map))

	EvaluateDependencies(obj, "LogicalInterface")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateLogicalInterface(ctx, request)
}

// UpdateLogicalInterface handles update request
func (service *PluginService) UpdateLogicalInterface(ctx context.Context, request *services.UpdateLogicalInterfaceRequest) (*services.UpdateLogicalInterfaceResponse, error) {
	log.Println(" UpdateLogicalInterface ENTERED")

	obj := request.GetLogicalInterface()

	intentObj := LogicalInterfaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("LogicalInterfaceIntent"); !ok {
		compilationif.ObjsCache.Store("LogicalInterfaceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "LogicalInterface")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateLogicalInterface", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateLogicalInterface(ctx, request)
}

// DeleteLogicalInterface handles delete request
func (service *PluginService) DeleteLogicalInterface(ctx context.Context, request *services.DeleteLogicalInterfaceRequest) (*services.DeleteLogicalInterfaceResponse, error) {
	log.Println(" DeleteLogicalInterface ENTERED")

	objUUID := request.GetID()

	//intentObj := LogicalInterfaceIntent {
	//LogicalInterface: *obj,
	//}

	//EvaluateDependencies(intentObj, "LogicalInterface")

	objMap, ok := compilationif.ObjsCache.Load("LogicalInterfaceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteLogicalInterface", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteLogicalInterface(ctx, request)
}

// GetLogicalInterface handles get request
func (service *PluginService) GetLogicalInterface(ctx context.Context, request *services.GetLogicalInterfaceRequest) (*services.GetLogicalInterfaceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("LogicalInterface")
	if !ok {
		return nil, errors.New("LogicalInterface get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("LogicalInterface get failed ")
	}

	response := &services.GetLogicalInterfaceResponse{
		LogicalInterface: obj.(*models.LogicalInterface),
	}
	return response, nil
}
