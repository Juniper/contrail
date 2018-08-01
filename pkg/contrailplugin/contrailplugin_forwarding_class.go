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

// ForwardingClassIntent
//   A struct to store attributes related to ForwardingClass
//   needed by Intent Compiler
type ForwardingClassIntent struct {
	Uuid string
}

// EvaluateForwardingClass - evaluates the ForwardingClass
func EvaluateForwardingClass(obj interface{}) {
	resourceObj := obj.(ForwardingClassIntent)
	log.Println("EvaluateForwardingClass Called ", resourceObj)
}

// CreateForwardingClass handles create request
func (service *PluginService) CreateForwardingClass(ctx context.Context, request *services.CreateForwardingClassRequest) (*services.CreateForwardingClassResponse, error) {
	log.Println(" CreateForwardingClass Entered")

	obj := request.GetForwardingClass()

	intentObj := ForwardingClassIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ForwardingClassIntent"); !ok {
		compilationif.ObjsCache.Store("ForwardingClassIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("ForwardingClassIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateForwardingClass", objMap.(*sync.Map))

	EvaluateDependencies(obj, "ForwardingClass")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateForwardingClass(ctx, request)
}

// UpdateForwardingClass handles update request
func (service *PluginService) UpdateForwardingClass(ctx context.Context, request *services.UpdateForwardingClassRequest) (*services.UpdateForwardingClassResponse, error) {
	log.Println(" UpdateForwardingClass ENTERED")

	obj := request.GetForwardingClass()

	intentObj := ForwardingClassIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("ForwardingClassIntent"); !ok {
		compilationif.ObjsCache.Store("ForwardingClassIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "ForwardingClass")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateForwardingClass", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateForwardingClass(ctx, request)
}

// DeleteForwardingClass handles delete request
func (service *PluginService) DeleteForwardingClass(ctx context.Context, request *services.DeleteForwardingClassRequest) (*services.DeleteForwardingClassResponse, error) {
	log.Println(" DeleteForwardingClass ENTERED")

	objUUID := request.GetID()

	//intentObj := ForwardingClassIntent {
	//ForwardingClass: *obj,
	//}

	//EvaluateDependencies(intentObj, "ForwardingClass")

	objMap, ok := compilationif.ObjsCache.Load("ForwardingClassIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteForwardingClass", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteForwardingClass(ctx, request)
}

// GetForwardingClass handles get request
func (service *PluginService) GetForwardingClass(ctx context.Context, request *services.GetForwardingClassRequest) (*services.GetForwardingClassResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("ForwardingClass")
	if !ok {
		return nil, errors.New("ForwardingClass get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("ForwardingClass get failed ")
	}

	response := &services.GetForwardingClassResponse{
		ForwardingClass: obj.(*models.ForwardingClass),
	}
	return response, nil
}
