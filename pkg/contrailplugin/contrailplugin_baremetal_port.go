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

// BaremetalPortIntent
//   A struct to store attributes related to BaremetalPort
//   needed by Intent Compiler
type BaremetalPortIntent struct {
	Uuid string
}

// EvaluateBaremetalPort - evaluates the BaremetalPort
func EvaluateBaremetalPort(obj interface{}) {
	resourceObj := obj.(BaremetalPortIntent)
	log.Println("EvaluateBaremetalPort Called ", resourceObj)
}

// CreateBaremetalPort handles create request
func (service *PluginService) CreateBaremetalPort(ctx context.Context, request *services.CreateBaremetalPortRequest) (*services.CreateBaremetalPortResponse, error) {
	log.Println(" CreateBaremetalPort Entered")

	obj := request.GetBaremetalPort()

	intentObj := BaremetalPortIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BaremetalPortIntent"); !ok {
		compilationif.ObjsCache.Store("BaremetalPortIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("BaremetalPortIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateBaremetalPort", objMap.(*sync.Map))

	EvaluateDependencies(obj, "BaremetalPort")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateBaremetalPort(ctx, request)
}

// UpdateBaremetalPort handles update request
func (service *PluginService) UpdateBaremetalPort(ctx context.Context, request *services.UpdateBaremetalPortRequest) (*services.UpdateBaremetalPortResponse, error) {
	log.Println(" UpdateBaremetalPort ENTERED")

	obj := request.GetBaremetalPort()

	intentObj := BaremetalPortIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("BaremetalPortIntent"); !ok {
		compilationif.ObjsCache.Store("BaremetalPortIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "BaremetalPort")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateBaremetalPort", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateBaremetalPort(ctx, request)
}

// DeleteBaremetalPort handles delete request
func (service *PluginService) DeleteBaremetalPort(ctx context.Context, request *services.DeleteBaremetalPortRequest) (*services.DeleteBaremetalPortResponse, error) {
	log.Println(" DeleteBaremetalPort ENTERED")

	objUUID := request.GetID()

	//intentObj := BaremetalPortIntent {
	//BaremetalPort: *obj,
	//}

	//EvaluateDependencies(intentObj, "BaremetalPort")

	objMap, ok := compilationif.ObjsCache.Load("BaremetalPortIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteBaremetalPort", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteBaremetalPort(ctx, request)
}

// GetBaremetalPort handles get request
func (service *PluginService) GetBaremetalPort(ctx context.Context, request *services.GetBaremetalPortRequest) (*services.GetBaremetalPortResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("BaremetalPort")
	if !ok {
		return nil, errors.New("BaremetalPort get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("BaremetalPort get failed ")
	}

	response := &services.GetBaremetalPortResponse{
		BaremetalPort: obj.(*models.BaremetalPort),
	}
	return response, nil
}
