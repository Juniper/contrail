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

// KubernetesNodeIntent
//   A struct to store attributes related to KubernetesNode
//   needed by Intent Compiler
type KubernetesNodeIntent struct {
	Uuid string
}

// EvaluateKubernetesNode - evaluates the KubernetesNode
func EvaluateKubernetesNode(obj interface{}) {
	resourceObj := obj.(KubernetesNodeIntent)
	log.Println("EvaluateKubernetesNode Called ", resourceObj)
}

// CreateKubernetesNode handles create request
func (service *PluginService) CreateKubernetesNode(ctx context.Context, request *services.CreateKubernetesNodeRequest) (*services.CreateKubernetesNodeResponse, error) {
	log.Println(" CreateKubernetesNode Entered")

	obj := request.GetKubernetesNode()

	intentObj := KubernetesNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KubernetesNodeIntent"); !ok {
		compilationif.ObjsCache.Store("KubernetesNodeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("KubernetesNodeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateKubernetesNode", objMap.(*sync.Map))

	EvaluateDependencies(obj, "KubernetesNode")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateKubernetesNode(ctx, request)
}

// UpdateKubernetesNode handles update request
func (service *PluginService) UpdateKubernetesNode(ctx context.Context, request *services.UpdateKubernetesNodeRequest) (*services.UpdateKubernetesNodeResponse, error) {
	log.Println(" UpdateKubernetesNode ENTERED")

	obj := request.GetKubernetesNode()

	intentObj := KubernetesNodeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("KubernetesNodeIntent"); !ok {
		compilationif.ObjsCache.Store("KubernetesNodeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "KubernetesNode")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateKubernetesNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateKubernetesNode(ctx, request)
}

// DeleteKubernetesNode handles delete request
func (service *PluginService) DeleteKubernetesNode(ctx context.Context, request *services.DeleteKubernetesNodeRequest) (*services.DeleteKubernetesNodeResponse, error) {
	log.Println(" DeleteKubernetesNode ENTERED")

	objUUID := request.GetID()

	//intentObj := KubernetesNodeIntent {
	//KubernetesNode: *obj,
	//}

	//EvaluateDependencies(intentObj, "KubernetesNode")

	objMap, ok := compilationif.ObjsCache.Load("KubernetesNodeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteKubernetesNode", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteKubernetesNode(ctx, request)
}

// GetKubernetesNode handles get request
func (service *PluginService) GetKubernetesNode(ctx context.Context, request *services.GetKubernetesNodeRequest) (*services.GetKubernetesNodeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("KubernetesNode")
	if !ok {
		return nil, errors.New("KubernetesNode get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("KubernetesNode get failed ")
	}

	response := &services.GetKubernetesNodeResponse{
		KubernetesNode: obj.(*models.KubernetesNode),
	}
	return response, nil
}
