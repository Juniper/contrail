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

// NamespaceIntent
//   A struct to store attributes related to Namespace
//   needed by Intent Compiler
type NamespaceIntent struct {
	Uuid string
}

// EvaluateNamespace - evaluates the Namespace
func EvaluateNamespace(obj interface{}) {
	resourceObj := obj.(NamespaceIntent)
	log.Println("EvaluateNamespace Called ", resourceObj)
}

// CreateNamespace handles create request
func (service *PluginService) CreateNamespace(ctx context.Context, request *services.CreateNamespaceRequest) (*services.CreateNamespaceResponse, error) {
	log.Println(" CreateNamespace Entered")

	obj := request.GetNamespace()

	intentObj := NamespaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NamespaceIntent"); !ok {
		compilationif.ObjsCache.Store("NamespaceIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("NamespaceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateNamespace", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Namespace")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateNamespace(ctx, request)
}

// UpdateNamespace handles update request
func (service *PluginService) UpdateNamespace(ctx context.Context, request *services.UpdateNamespaceRequest) (*services.UpdateNamespaceResponse, error) {
	log.Println(" UpdateNamespace ENTERED")

	obj := request.GetNamespace()

	intentObj := NamespaceIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("NamespaceIntent"); !ok {
		compilationif.ObjsCache.Store("NamespaceIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Namespace")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateNamespace", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateNamespace(ctx, request)
}

// DeleteNamespace handles delete request
func (service *PluginService) DeleteNamespace(ctx context.Context, request *services.DeleteNamespaceRequest) (*services.DeleteNamespaceResponse, error) {
	log.Println(" DeleteNamespace ENTERED")

	objUUID := request.GetID()

	//intentObj := NamespaceIntent {
	//Namespace: *obj,
	//}

	//EvaluateDependencies(intentObj, "Namespace")

	objMap, ok := compilationif.ObjsCache.Load("NamespaceIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteNamespace", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteNamespace(ctx, request)
}

// GetNamespace handles get request
func (service *PluginService) GetNamespace(ctx context.Context, request *services.GetNamespaceRequest) (*services.GetNamespaceResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Namespace")
	if !ok {
		return nil, errors.New("Namespace get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Namespace get failed ")
	}

	response := &services.GetNamespaceResponse{
		Namespace: obj.(*models.Namespace),
	}
	return response, nil
}
