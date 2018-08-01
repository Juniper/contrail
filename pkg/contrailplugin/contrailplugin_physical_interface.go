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

// PhysicalInterfaceIntent
//   A struct to store attributes related to PhysicalInterface
//   needed by Intent Compiler
type PhysicalInterfaceIntent struct {
	Uuid string
}

// EvaluatePhysicalInterface - evaluates the PhysicalInterface
func EvaluatePhysicalInterface(obj interface{}) {
	resourceObj := obj.(PhysicalInterfaceIntent)
	log.Println("EvaluatePhysicalInterface Called ", resourceObj)
}

// CreatePhysicalInterface handles create request
func (service *PluginService) CreatePhysicalInterface(ctx context.Context, request *services.CreatePhysicalInterfaceRequest) (*services.CreatePhysicalInterfaceResponse, error) {
	log.Println(" CreatePhysicalInterface Entered")

	obj := request.GetPhysicalInterface()

	intentObj := PhysicalInterfaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PhysicalInterfaceIntent"); !ok {
		compilationif.ObjsCache.Store("PhysicalInterfaceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("PhysicalInterfaceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreatePhysicalInterface", objMap.(*sync.Map))

	EvaluateDependencies(obj, "PhysicalInterface")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreatePhysicalInterface(ctx, request)
}

// UpdatePhysicalInterface handles update request
func (service *PluginService) UpdatePhysicalInterface(ctx context.Context, request *services.UpdatePhysicalInterfaceRequest) (*services.UpdatePhysicalInterfaceResponse, error) {
	log.Println(" UpdatePhysicalInterface ENTERED")

	obj := request.GetPhysicalInterface()

	intentObj := PhysicalInterfaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("PhysicalInterfaceIntent"); !ok {
		compilationif.ObjsCache.Store("PhysicalInterfaceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "PhysicalInterface")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdatePhysicalInterface", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdatePhysicalInterface(ctx, request)
}

// DeletePhysicalInterface handles delete request
func (service *PluginService) DeletePhysicalInterface(ctx context.Context, request *services.DeletePhysicalInterfaceRequest) (*services.DeletePhysicalInterfaceResponse, error) {
	log.Println(" DeletePhysicalInterface ENTERED")

	objUUID := request.GetID()

	//intentObj := PhysicalInterfaceIntent {
	//PhysicalInterface: *obj,
	//}

	//EvaluateDependencies(intentObj, "PhysicalInterface")

	objMap, ok := compilationif.ObjsCache.Load("PhysicalInterfaceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeletePhysicalInterface", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeletePhysicalInterface(ctx, request)
}

// GetPhysicalInterface handles get request
func (service *PluginService) GetPhysicalInterface(ctx context.Context, request *services.GetPhysicalInterfaceRequest) (*services.GetPhysicalInterfaceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("PhysicalInterface")
	if !ok {
		return nil, errors.New("PhysicalInterface get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("PhysicalInterface get failed ")
	}

	response := &services.GetPhysicalInterfaceResponse{
		PhysicalInterface: obj.(*models.PhysicalInterface),
	}
	return response, nil
}
