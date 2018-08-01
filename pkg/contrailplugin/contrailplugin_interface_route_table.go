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

// InterfaceRouteTableIntent
//   A struct to store attributes related to InterfaceRouteTable
//   needed by Intent Compiler
type InterfaceRouteTableIntent struct {
	Uuid string
}

// EvaluateInterfaceRouteTable - evaluates the InterfaceRouteTable
func EvaluateInterfaceRouteTable(obj interface{}) {
	resourceObj := obj.(InterfaceRouteTableIntent)
	log.Println("EvaluateInterfaceRouteTable Called ", resourceObj)
}

// CreateInterfaceRouteTable handles create request
func (service *PluginService) CreateInterfaceRouteTable(ctx context.Context, request *services.CreateInterfaceRouteTableRequest) (*services.CreateInterfaceRouteTableResponse, error) {
	log.Println(" CreateInterfaceRouteTable Entered")

	obj := request.GetInterfaceRouteTable()

	intentObj := InterfaceRouteTableIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("InterfaceRouteTableIntent"); !ok {
		compilationif.ObjsCache.Store("InterfaceRouteTableIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("InterfaceRouteTableIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateInterfaceRouteTable", objMap.(*sync.Map))

	EvaluateDependencies(obj, "InterfaceRouteTable")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateInterfaceRouteTable(ctx, request)
}

// UpdateInterfaceRouteTable handles update request
func (service *PluginService) UpdateInterfaceRouteTable(ctx context.Context, request *services.UpdateInterfaceRouteTableRequest) (*services.UpdateInterfaceRouteTableResponse, error) {
	log.Println(" UpdateInterfaceRouteTable ENTERED")

	obj := request.GetInterfaceRouteTable()

	intentObj := InterfaceRouteTableIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("InterfaceRouteTableIntent"); !ok {
		compilationif.ObjsCache.Store("InterfaceRouteTableIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "InterfaceRouteTable")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateInterfaceRouteTable", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateInterfaceRouteTable(ctx, request)
}

// DeleteInterfaceRouteTable handles delete request
func (service *PluginService) DeleteInterfaceRouteTable(ctx context.Context, request *services.DeleteInterfaceRouteTableRequest) (*services.DeleteInterfaceRouteTableResponse, error) {
	log.Println(" DeleteInterfaceRouteTable ENTERED")

	objUUID := request.GetID()

	//intentObj := InterfaceRouteTableIntent {
	//InterfaceRouteTable: *obj,
	//}

	//EvaluateDependencies(intentObj, "InterfaceRouteTable")

	objMap, ok := compilationif.ObjsCache.Load("InterfaceRouteTableIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteInterfaceRouteTable", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteInterfaceRouteTable(ctx, request)
}

// GetInterfaceRouteTable handles get request
func (service *PluginService) GetInterfaceRouteTable(ctx context.Context, request *services.GetInterfaceRouteTableRequest) (*services.GetInterfaceRouteTableResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("InterfaceRouteTable")
	if !ok {
		return nil, errors.New("InterfaceRouteTable get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("InterfaceRouteTable get failed ")
	}

	response := &services.GetInterfaceRouteTableResponse{
		InterfaceRouteTable: obj.(*models.InterfaceRouteTable),
	}
	return response, nil
}
